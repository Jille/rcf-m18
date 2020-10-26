package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func tokenizeAddress(ln string) []string {
	var ret []string
	s := 0
	quoted := false
	for i, c := range ln {
		switch c {
		case ' ':
			if quoted {
				continue
			}
			ret = append(ret, ln[s:i])
			s = i + 1
		case '"':
			quoted = !quoted
		}
	}
	ret = append(ret, ln[s:])
	return ret
}

type param struct {
	Static   interface{}
	APIType  string
	WireType string
}

type paramSlice []param

type enumValue struct {
	Name string
	Literal string
	Value interface{}
}

type enum struct {
	Name string
	Type string
	Values []enumValue
}

var enums = map[string]enum{}

type address struct {
	Address     string
	MinPage     int
	MaxPage     int
	MinChannel  int
	MaxChannel  int
	MinControl  int
	MaxControl  int
	AddressType string
	Signature   string
	Name        string
	Ins         paramSlice
	Outs        paramSlice
	Send        []string
	Receive     []string
	Encode      []string
	Decode      []string
}

func parseRange(s string) (int, int) {
	ex := strings.Split(s, "-")
	f, err := strconv.Atoi(ex[0])
	if err != nil {
		panic(err)
	}
	l, err := strconv.Atoi(ex[len(ex)-1])
	if err != nil {
		panic(err)
	}
	return int(f), int(l)
}

// Set {Min,Max}{Page,Channel,Control} and AddressType
func (a *address) ParseAddress() {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Sprintf("a.ParseAddress: %v", r))
		}
	}()
	re := regexp.MustCompile(`^/(\d{1,2}(?:-\d{1,2})?)/(\d{1,2}(?:-\d{1,2})?)/([a-z]{2,})_(\d{1,3}(?:-\d{1,3})?)$`)
	m := re.FindStringSubmatch(a.Address)
	a.MinPage, a.MaxPage = parseRange(m[1])
	a.MinChannel, a.MaxChannel = parseRange(m[2])
	a.AddressType = m[3]
	a.MinControl, a.MaxControl = parseRange(m[4])
}

func parseLiteral(a string) (interface{}, string) {
	if a[0] == '"' && a[len(a)-1] == '"' {
		return a[1 : len(a)-1], "string"
	} else {
		f, err := strconv.ParseFloat(a, 32)
		if err != nil {
			panic(err)
		}
		return float32(f), "float32"
	}
}

func parseArgs(args string) []param {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Sprintf("parseArgs: %v", r))
		}
	}()
	if args == "-" {
		return nil
	}
	var ret []param
	ex := strings.Split(args, ",")
	for _, a := range ex {
		var p param
		switch a {
		case "string", "float32":
			p.APIType = a
			p.WireType = a
		case "bool":
			p.APIType = a
			p.WireType = "float32"
		case "stringbool":
			p.APIType = "bool"
			p.WireType = "string"
		default:
			if e, found := enums[a]; found {
				p.APIType = e.Name
				p.WireType = e.Type
			} else {
				p.Static, p.APIType = parseLiteral(a)
				p.WireType = p.APIType
			}
		}
		ret = append(ret, p)
	}
	return ret
}

func parseAddress(ln string) *address {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Sprintf("parseAddress: %v", r))
		}
	}()
	ret := &address{}
	ex := tokenizeAddress(ln)
	ret.Address = ex[0]
	ret.ParseAddress()
	if ex[1] != "-" {
		panic("expected - as second argument")
	}
	ret.Signature = ex[2]
	ret.Name = ex[3]
	switch ret.Signature {
	case "request", "connset":
		ret.Ins = parseArgs(ex[4])
		ret.Outs = parseArgs(ex[5])
	case "set":
		ret.Ins = parseArgs(ex[4])
	}
	return ret
}

func parse(in string) ([]address, []enum) {
	var as []address
	var es []enum
	i := 0
	var a *address
	var e *enum
	var codeBlock *[]string
	lines := strings.Split(in, "\n")
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Sprintf("Failed to parse line %d (%q): %v", i+1, lines[i], r))
		}
	}()
	for ; len(lines) > i; i++ {
		ln := lines[i]
		if strings.TrimSpace(ln) == "" {
			continue
		}
		switch ln[0] {
		case '#':
			continue
		case '/':
			if a != nil {
				as = append(as, *a)
			}
			if e != nil {
				es = append(es, *e)
				enums[e.Name] = *e
				e = nil
			}
			a = parseAddress(ln)
		case '!':
			if a != nil {
				as = append(as, *a)
				a = nil
			}
			if e != nil {
				es = append(es, *e)
				enums[e.Name] = *e
			}
			ts := tokenizeAddress(ln)
			e = &enum{
				Name: ts[1],
			}
		case '	':
			if e != nil {
				ts := tokenizeAddress(ln[1:])
				v, t := parseLiteral(ts[0])
				if e.Type == "" {
					e.Type = t
				} else if e.Type != t {
					panic("enum values must be the same type")
				}
				e.Values = append(e.Values, enumValue{
					Name: ts[1],
					Literal: ts[0],
					Value: v,
				})
			} else if strings.HasPrefix(ln, "\t>send") {
				codeBlock = &a.Send
			} else if strings.HasPrefix(ln, "\t>receive") {
				codeBlock = &a.Receive
			} else if strings.HasPrefix(ln, "\t>encode") {
				codeBlock = &a.Encode
			} else if strings.HasPrefix(ln, "\t>decode") {
				codeBlock = &a.Decode
			} else if ln[1] == '	' {
				*codeBlock = append(*codeBlock, ln[2:])
			}
		default:
			panic("unknown line prefix")
		}
	}
	if a != nil {
		as = append(as, *a)
	}
	if e != nil {
		es = append(es, *e)
		enums[e.Name] = *e
	}
	return as, es
}

func run() error {
	data, err := ioutil.ReadFile("source.txt")
	if err != nil {
		return err
	}
	as, es := parse(string(data))
	fmt.Println(gen(as, es))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type out struct {
	buf    bytes.Buffer
	indent int
}

func (o *out) wln(f string, args ...interface{}) {
	if (strings.HasSuffix(f, "}") || strings.HasSuffix(f, "},")) && !strings.HasSuffix(f, "{}") {
		o.indent--
	}
	fmt.Fprintf(&o.buf, "%s", strings.Repeat("\t", o.indent))
	fmt.Fprintf(&o.buf, f+"\n", args...)
	if strings.HasSuffix(f, "{") {
		o.indent++
	}
}

type FuncSig int

const (
	PassIns FuncSig = 1 << iota
	PassOuts
	PassContext
	ReturnOuts

	ClientRequest  = PassIns
	ServerHandler  = PassContext | PassIns | ReturnOuts
	ClientResponse = PassContext | PassOuts
	ClientSet      = PassContext | PassIns
)

func (a *address) AddressVariables() []string {
	var ret []string
	if a.MinPage != a.MaxPage {
		ret = append(ret, "page")
	}
	if a.MinChannel != a.MaxChannel {
		ret = append(ret, "channel")
	}
	if a.MinControl != a.MaxControl {
		ret = append(ret, "control")
	}
	return ret
}

func (a *address) FunctionSignature(f string, t FuncSig) string {
	var in []string
	if t&PassContext > 0 {
		in = append(in, "c *Context")
	}
	for _, v := range a.AddressVariables() {
		in = append(in, v+" int")
	}
	if t&PassIns > 0 {
		in = append(in, a.Ins.goParams(true, true)...)
	}
	if t&PassOuts > 0 {
		in = append(in, a.Outs.goParams(true, true)...)
	}
	var out []string
	if t&ReturnOuts > 0 {
		out = append(out, a.Outs.goParams(false, true)...)
	}
	out = append(out, "error")
	fn := fmt.Sprintf(f, a.Name)
	return fmt.Sprintf("%s(%s) (%s)", fn, strings.Join(in, ", "), strings.Join(out, ", "))
}

func (ps paramSlice) goParams(named bool, typed bool) []string {
	var ret []string
	for i, p := range ps {
		if p.Static != nil {
			continue
		}
		var o []string
		if named {
			o = append(o, ps.nameAt(i))
		}
		if typed {
			o = append(o, p.APIType)
		}
		ret = append(ret, strings.Join(o, " "))
	}
	return ret
}

func (ps paramSlice) nameAt(i int) string {
	p := ps[i]
	if p.Static != nil {
		panic("nameAt called for static parameter")
	}
	return fmt.Sprintf("%c%d", p.APIType[0], i+1)
}

func gen(as []address, es []enum) string {
	o := &out{}
	o.wln(`package rcfm18`)
	o.wln(`// This file is generated. Don't make modifications here.`)
	o.wln(``)
	for _, e := range es {
		o.wln(`type %s struct{`, e.Name)
		o.wln(`v %s`, e.Type)
		o.wln(`}`)
		o.wln(`func (e %s) ToWire() %s {`, e.Name, e.Type)
		o.wln(`return e.v`)
		o.wln(`}`)
		o.wln(`func (e %s) String() string {`, e.Name)
		o.wln(`switch e.v {`)
		for _, v := range e.Values {
			o.indent--
			o.wln(`case %s:`, v.Literal)
			o.indent++
			o.wln(`return %q`, v.Name)
		}
		o.indent--
		o.wln(`default:`)
		o.indent++
		o.wln(`panic("unknown enum value")`)
		o.wln(`}`)
		o.wln(`}`)
		o.wln(`var (`)
		o.indent++
		for _, v := range e.Values {
			o.wln(`%s = %s{%s} `, v.Name, e.Name, v.Literal)
		}
		o.indent--
		o.wln(`)`)
		o.wln(`func To%s(v %s) (%s, error) {`, e.Name, e.Type, e.Name)
		o.wln(`switch v {`)
		for _, v := range e.Values {
			o.indent--
			o.wln(`case %s:`, v.Literal)
			o.indent++
			o.wln(`return %s, nil`, v.Name)
		}
		o.indent--
		o.wln(`default:`)
		o.indent++
		o.wln(`return %s, errors.New("unknown enum value")`, e.Values[0].Name)
		o.wln(`}`)
		o.wln(`}`)
		o.wln(``)
	}
	o.wln(`type RCFM18Server struct {`)
	o.wln(`connections map[string]*Connection`)
	o.wln(`MangleReturnAddress func(net.Addr) net.Addr`)
	for _, a := range as {
		switch a.Signature {
		case "meters":
			continue
		case "request":
			o.wln(a.FunctionSignature("On%sRequest func", ServerHandler))
		case "connset", "set":
			o.wln(a.FunctionSignature("On%sSet func", ServerHandler))
		}
		o.wln(`%sMessageHandler func(c *Context) error`, a.Name)
	}
	o.wln(`}`)
	o.wln(`func NewServer() *RCFM18Server {`)
	o.wln(`s := &RCFM18Server{}`)
	o.wln(`s.connections = map[string]*Connection{}`)
	o.wln(`s.MangleReturnAddress = func(src net.Addr) net.Addr {`)
	o.wln(`host, _, err := net.SplitHostPort(src.String())`)
	o.wln(`if err != nil {`)
	o.wln(`panic(err)`)
	o.wln(`}`)
	o.wln(`a, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, "9000"))`)
	o.wln(`if err != nil {`)
	o.wln(`panic(err)`)
	o.wln(`}`)
	o.wln(`return a`)
	o.wln(`}`)
	for _, a := range as {
		var fn string
		switch a.Signature {
		case "meters":
			continue
		case "request":
			fn = fmt.Sprintf("On%sRequest", a.Name)
			if len(a.Receive) > 0 {
				o.wln(a.FunctionSignature("s.On%sRequest = func", ServerHandler) + " {")
				o.printArgs(fn, append(a.AddressVariables(), a.Ins.goParams(true, false)...))
				o.funcBody(a.Receive, true, a.Outs)
				o.wln(`}`)
			}
		case "connset", "set":
			fn = fmt.Sprintf("On%sSet", a.Name)
			o.wln(a.FunctionSignature("s.On%sSet = func", ServerHandler) + " {")
			o.printArgs(fn, append(a.AddressVariables(), a.Ins.goParams(true, false)...))
			o.funcBody(a.Receive, true, a.Outs)
			o.wln(`}`)
		}
		o.generateMessageHandler(a, fn, "s", ServerHandler)
	}
	o.wln(`return s`)
	o.wln(`}`)
	o.wln(``)
	o.wln(`func (s *RCFM18Server) Dispatch(cx *Context) error {`)
	o.generateDispatcher(as, "s", false)
	o.wln(`}`)
	o.wln(``)
	o.wln(`type RCFM18Client struct {`)
	o.wln(`conn *Connection`)
	o.wln(`Send func(p osc.Packet) error`)
	for _, a := range as {
		switch a.Signature {
		case "request":
			o.wln(a.FunctionSignature("On%sResponse func", ClientResponse))
		case "connset":
			o.wln(a.FunctionSignature("On%sSet func", ClientResponse))
		case "set":
			o.wln(a.FunctionSignature("On%sSet func", ClientSet))
		case "meters":
			o.wln(`OnMeters func(c *Context, inputs, player, fxReturns, main, fxSends, aux, phones, recording []float32) error`)
		}
		o.wln(`%sMessageHandler func(c *Context) error`, a.Name)
	}
	o.wln(`}`)
	o.wln(``)
	o.wln(`func (c *RCFM18Client) setupCallbacks() {`)
	for _, a := range as {
		switch a.Signature {
		case "request":
			o.generateMessageHandler(a, fmt.Sprintf("On%sResponse", a.Name), "c", ClientResponse)
		case "connset":
			o.generateMessageHandler(a, fmt.Sprintf("On%sSet", a.Name), "c", ClientResponse)
		case "set":
			o.generateMessageHandler(a, fmt.Sprintf("On%sSet", a.Name), "c", ClientSet)
		}
	}
	o.wln(`}`)
	o.wln(``)
	o.wln(`func (c *RCFM18Client) Dispatch(cx *Context) error {`)
	o.generateDispatcher(as, "c", true)
	o.wln(`}`)

	for _, a := range as {
		o.wln(``)
		switch a.Signature {
		case "request":
			o.wln(a.FunctionSignature("func (c *RCFM18Client) Request%s", ClientRequest) + " {")
		case "connset", "set":
			o.wln(a.FunctionSignature("func (c *RCFM18Client) Set%s", ClientRequest) + " {") // }
		case "meters":
			continue
		}
		o.funcBody(a.Encode, false, nil)
		f := "/"
		var args []string
		if a.MinPage == a.MaxPage {
			f += fmt.Sprintf("%02d/", a.MinPage)
		} else {
			o.wln(`if page < %d || page > %d {`, a.MinPage, a.MaxPage)
			o.wln(`return fmt.Errorf("page out of range")`)
			o.wln(`}`)
			f += "%02d/"
			args = append(args, "page")
		}
		if a.MinChannel == a.MaxChannel {
			f += fmt.Sprintf("%02d/", a.MinChannel)
		} else {
			o.wln(`if channel < %d || channel > %d {`, a.MinChannel, a.MaxChannel)
			o.wln(`return fmt.Errorf("channel out of range")`)
			o.wln(`}`)
			f += "%02d/"
			args = append(args, "channel")
		}
		f += a.AddressType + "_"
		if a.MinControl == a.MaxControl {
			f += fmt.Sprintf("%03d", a.MinControl)
		} else {
			o.wln(`if control < %d || control > %d {`, a.MinControl, a.MaxControl)
			o.wln(`return fmt.Errorf("control out of range")`)
			o.wln(`}`)
			f += "%03d"
			args = append(args, "control")
		}
		if len(args) == 0 {
			o.wln(`msg := osc.NewMessage(%q)`, a.Address)
		} else {
			o.wln(`msg := osc.NewMessage(fmt.Sprintf(%q, %s))`, f, strings.Join(args, ", "))
		}
		o.appendParamSlice("msg", a.Ins)
		o.wln(`return c.Send(msg)`)
		o.wln(`}`)
	}

	return o.buf.String()
}

func (o *out) generateDispatcher(as []address, vn string, client bool) {
	for _, a := range as {
		if a.Name == "Meters" && !client {
			continue
		}
		var conds []string
		conds = append(conds, fmt.Sprintf("cx.addressType == %q", a.AddressType))
		if a.MinPage == a.MaxPage {
			conds = append(conds, fmt.Sprintf("cx.page == %d", a.MinPage))
		} else {
			conds = append(conds, fmt.Sprintf("cx.page >= %d", a.MinPage))
			conds = append(conds, fmt.Sprintf("cx.page <= %d", a.MaxPage))
		}
		if a.MinChannel == a.MaxChannel {
			conds = append(conds, fmt.Sprintf("cx.channel == %d", a.MinChannel))
		} else {
			conds = append(conds, fmt.Sprintf("cx.channel >= %d", a.MinChannel))
			conds = append(conds, fmt.Sprintf("cx.channel <= %d", a.MaxChannel))
		}
		if a.MinControl == a.MaxControl {
			conds = append(conds, fmt.Sprintf("cx.control == %d", a.MinControl))
		} else {
			conds = append(conds, fmt.Sprintf("cx.control >= %d", a.MinControl))
			conds = append(conds, fmt.Sprintf("cx.control <= %d", a.MaxControl))
		}
		o.wln(`if %s {`, strings.Join(conds, " && "))
		o.wln(`return %s.%sMessageHandler(cx)`, vn, a.Name)
		o.wln(`}`)
	}
	o.wln(`return ErrNoHandler`)
}

func (o *out) generateMessageHandler(a address, fn, vn string, t FuncSig) {
	o.wln(`%s.%sMessageHandler = func(cx *Context) error {`, vn, a.Name)
	o.wln(`if len(cx.msg.Arguments) != %d {`, len(a.Ins))
	o.wln(`return fmt.Errorf("message had %%d argument(s), expected %d", len(cx.msg.Arguments))`, len(a.Ins))
	o.wln(`}`)
	ins := []string{"cx"}
	for _, v := range a.AddressVariables() {
		o.wln("%s := cx.%s", v, v)
		ins = append(ins, v)
	}
	o.funcBody(a.Decode, false, nil)
	inp := a.Ins
	if t & PassOuts > 0 {
		inp = a.Outs
	}
	for i, p := range inp {
		if p.Static == nil {
			n := inp.nameAt(i)
			if p.WireType != p.APIType {
				n = fmt.Sprintf("%c%s", p.WireType[0], n)
			}
			o.wln(`%s, ok := cx.msg.Arguments[%d].(%s)`, n, i, p.WireType)
			o.wln(`if !ok {`)
			o.wln(`return fmt.Errorf("message should have had %s argument, had %%T", cx.msg.Arguments[%d])`, p.WireType, i)
			o.wln(`}`)
			fallible := false
			if e, found := enums[p.APIType]; found {
				o.wln(`%s, err := To%s(%s)`, inp.nameAt(i), e.Name, n)
				fallible = true
			} else if p.WireType != p.APIType {
				switch p.WireType {
				case "string":
					o.wln(`%s, err := stringToBool(%s)`, inp.nameAt(i), n)
				case "float32":
					o.wln(`%s, err := floatToBool(%s)`, inp.nameAt(i), n)
				}
				fallible = true
			}
			if fallible {
				o.wln(`if err != nil {`)
				o.wln(`return fmt.Errorf("invalid value for %s: %%v", err)`, p.APIType)
				o.wln(`}`)
			}
			ins = append(ins, inp.nameAt(i))
		}
	}
	o.wln(`if %s.%s == nil {`, vn, fn)
	o.wln(`return ErrNoHandler`)
	o.wln(`}`)
	switch a.Signature {
	case "set":
		o.wln(`cx.broadcast = true`)
		o.wln(`cx.reply = cx.msg`)
	}
	outp := a.Outs
	if t & ReturnOuts == 0 {
		outp = nil
	}
	outs := outp.goParams(true, false)
	outs = append(outs, "err2")
	o.wln(`%s := %s.%s(%s)`, strings.Join(outs, ", "), vn, fn, strings.Join(ins, ", "))
	o.wln(`if err2 != nil {`)
	o.wln(`return err2`)
	o.wln(`}`)
	if len(outp) > 0 {
		o.wln(`if cx.reply != nil {`)
		o.appendParamSlice("cx.reply", outp)
		o.wln(`}`)
	}
	o.wln(`return nil`)
	o.wln(`}`)
}

func (o *out) appendParamSlice(vn string, ps paramSlice) {
	for i, p := range ps {
		if p.Static != nil {
			switch p.WireType {
			case "float32":
				o.wln(`%s.Append(float32(%v))`, vn, p.Static)
			case "string":
				o.wln(`%s.Append(%q)`, vn, p.Static)
			}
		} else if _, found := enums[p.APIType]; found {
			o.wln(`%s.Append(%s.ToWire())`, vn, ps.nameAt(i))
		} else {
			switch p.APIType +"-"+ p.WireType {
			case "bool-float32":
				o.wln(`%s.Append(boolToFloat(%s))`, vn, ps.nameAt(i))
			case "bool-string":
				o.wln(`%s.Append(boolToString(%s))`, vn, ps.nameAt(i))
			default:
				o.wln(`%s.Append(%s)`, vn, ps.nameAt(i))
			}
		}
	}
}

func (o *out) funcBody(lines []string, mustReturn bool, rets paramSlice) {
	var last string
	for _, ln := range lines {
		o.wln("%s", ln)
		last = ln
	}
	if !strings.HasPrefix(last, "return") && mustReturn {
		var s []string
		for _, p := range rets {
			if p.Static != nil {
				continue
			}
			switch p.APIType {
			case "string":
				s = append(s, `""`)
			case "float32":
				s = append(s, `0`)
			case "bool":
				s = append(s, `false`)
			}
		}
		s = append(s, "nil")
		o.wln("return %s", strings.Join(s, ", "))
	}
}

func (o *out) printArgs(fn string, args []string) {
	if len(args) == 0 {
		o.wln(`fmt.Println("%s()\n")`, fn)
		return
	}
	o.wln(`fmt.Printf("%s(%s)\n", %s)`, fn, strings.Repeat(", %v", len(args))[2:], strings.Join(args, ", "))
}
