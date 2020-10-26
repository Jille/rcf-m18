package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"regexp"
	"strconv"
	"sync"

	"github.com/Jille/go-osc/osc"
)

func main() {
	addr := "0.0.0.0:8000"

	fm18 := &FakeM18{}
	fm18.ListenAndServe(addr)
}

type FakeM18 struct {
	s                     *osc.Server
	mtx                   sync.Mutex
	OptimizedMeterSending bool
	Faders                [24][11]float32
}

func (f *FakeM18) ListenAndServe(addr string) error {
	f.s = &osc.Server{Addr: addr}

	f.s.Handle("*", f.Handle)

	return f.s.ListenAndServe()
}

func (f *FakeM18) Reply(m osc.Packet, src net.Addr) error {
	host, _, err := net.SplitHostPort(src.String())
	if err != nil {
		panic(err)
	}
	a, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, "9000"))
	if err != nil {
		return err
	}
	log.Printf("[%s] Reply: %s", src, m)
	return f.s.Reply(m, a)
}

func (f *FakeM18) Handle(msg *osc.Message, src net.Addr) {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	log.Printf("[%s] Packet: %s", src, msg)
	if err := f.handle(msg, src); err != nil {
		log.Printf("[%s] Handle of %s failed: %v", src, msg, err)
	}
}

var addrRe = regexp.MustCompile(`^/\d{2}/\d{2}/.._\d{3}$`)

func (f *FakeM18) handle(msg *osc.Message, src net.Addr) (retErr error) {
	/*
		defer func() {
			if r := recover(); r != nil {
				retErr = fmt.Errorf("panic: %v", r)
			}
		}()
	*/
	if !addrRe.MatchString(msg.Address) {
		return errors.New("address pattern mismatch")
	}
	page, err := strconv.Atoi(msg.Address[1:3])
	if err != nil {
		return err
	}
	channel, err := strconv.Atoi(msg.Address[4:6])
	if err != nil {
		return err
	}
	typ := msg.Address[7:9]
	control, err := strconv.Atoi(msg.Address[10:13])
	if err != nil {
		return err
	}
	return f.HandleMsg(msg, src, page, channel, typ, control)
}

func (f *FakeM18) HandleMsg(msg *osc.Message, src net.Addr, page, channel int, typ string, control int) error {
	rep := osc.NewMessage(msg.Address)
	switch typ {
	case "gb":
		switch page {
		case 22:
			switch control {
			case 100: // Firmware version
				return nil
			case 146: // Target id
				rep.Append(string("0")) // M18
				return f.Reply(rep, src)
			case 172: // Optimized meter sending
				s, err := f.StringArg(msg)
				if err != nil {
					return err
				}
				f.OptimizedMeterSending = (s == "1")
				return nil
			case 272: // Set date/time
				s, err := f.StringArg(msg)
				if err != nil {
					return err
				}
				// %Y%m%d%H%M.%s
				_ = s
			}
		case 27:
			switch control {
			case 78: // Get linked stereo channels
				return nil
			}
		}
	case "np": // Get preset names
		switch channel {
		case 11: // Snapshots
			return nil
		case 12: // Shows
			return nil
		case 13: // MultiFX3
			return nil
		case 14: // MultiFX5
			return nil
		case 15: // Graphic equalizer
			return nil
		}
	case "up":
		switch channel {
		case 99: // Request for channel update
			return nil
		}
	case "pl":
		if page == 0 && channel == 0 && control == 0 { // Poll
			if f.OptimizedMeterSending || true {
				clip := func(v float32) float32 {
					switch {
					case v < 0:
						return 0
					case v > 1:
						return 1
					default:
						return v
					}
				}
				rep = osc.NewMessage("/00/00/vmeter_000")
				for ch := 0; 20 > ch; ch++ {
					rep.Append(clip(f.Faders[ch][1] + (rand.Float32() / 10) - 0.05))
				}
				return f.Reply(rep, src)
			} else {
			}
			return nil
		}
	case "fd":
		v, err := f.FloatArg(msg)
		if err != nil {
			return err
		}
		f.Faders[channel][page] = v
		rep = osc.NewMessage(fmt.Sprintf("/%02d/%02d/fd_000", page, channel^1))
		rep.Append(float32(1 - v))
		return f.Reply(rep, src)
	}
	return errors.New("no handler")
}

func (f *FakeM18) StringArg(msg *osc.Message) (string, error) {
	if len(msg.Arguments) != 1 {
		return "", fmt.Errorf("message had %d arguments, expected 1", len(msg.Arguments))
	}
	s, ok := msg.Arguments[0].(string)
	if !ok {
		return "", fmt.Errorf("message should have had string argument, had %T", msg.Arguments[0])
	}
	return s, nil
}

func (f *FakeM18) FloatArg(msg *osc.Message) (float32, error) {
	if len(msg.Arguments) != 1 {
		return 0, fmt.Errorf("message had %d arguments, expected 1", len(msg.Arguments))
	}
	fv, ok := msg.Arguments[0].(float32)
	if !ok {
		return 0, fmt.Errorf("message should have had float32 argument, had %T", msg.Arguments[0])
	}
	return fv, nil
}

func (f *FakeM18) FloatsArg(msg *osc.Message) ([]float32, error) {
	ret := make([]float32, len(msg.Arguments))
	for i, v := range msg.Arguments {
		fv, ok := v.(float32)
		if !ok {
			return nil, fmt.Errorf("message should have had float32 argument, had %T", msg.Arguments[0])
		}
		ret[i] = fv
	}
	return ret, nil
}
