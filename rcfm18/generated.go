package rcfm18

import (
	"errors"
	"fmt"
	"net"

	"github.com/Jille/go-osc/osc"
)

// This file is generated. Don't make modifications here.

type EQMode struct {
	v float32
}

func (e EQMode) ToWire() float32 {
	return e.v
}
func (e EQMode) String() string {
	switch e.v {
	case 0:
		return "StandardEQ"
	case 1:
		return "VintageEQ"
	case 2:
		return "SmoothEQ"
	default:
		panic("unknown enum value")
	}
}

var (
	StandardEQ = EQMode{0}
	VintageEQ  = EQMode{1}
	SmoothEQ   = EQMode{2}
)

func ToEQMode(v float32) (EQMode, error) {
	switch v {
	case 0:
		return StandardEQ, nil
	case 1:
		return VintageEQ, nil
	case 2:
		return SmoothEQ, nil
	default:
		return StandardEQ, errors.New("unknown enum value")
	}
}

type EQSmoothness struct {
	v float32
}

func (e EQSmoothness) ToWire() float32 {
	return e.v
}
func (e EQSmoothness) String() string {
	switch e.v {
	case 0:
		return "SmoothnessA"
	case 0.5:
		return "SmoothnessB"
	case 1:
		return "SmoothnessC"
	default:
		panic("unknown enum value")
	}
}

var (
	SmoothnessA = EQSmoothness{0}
	SmoothnessB = EQSmoothness{0.5}
	SmoothnessC = EQSmoothness{1}
)

func ToEQSmoothness(v float32) (EQSmoothness, error) {
	switch v {
	case 0:
		return SmoothnessA, nil
	case 0.5:
		return SmoothnessB, nil
	case 1:
		return SmoothnessC, nil
	default:
		return SmoothnessA, errors.New("unknown enum value")
	}
}

type StartupMode struct {
	v string
}

func (e StartupMode) ToWire() string {
	return e.v
}
func (e StartupMode) String() string {
	switch e.v {
	case "0":
		return "InitMixer"
	case "1":
		return "LastShow"
	default:
		panic("unknown enum value")
	}
}

var (
	InitMixer = StartupMode{"0"}
	LastShow  = StartupMode{"1"}
)

func ToStartupMode(v string) (StartupMode, error) {
	switch v {
	case "0":
		return InitMixer, nil
	case "1":
		return LastShow, nil
	default:
		return InitMixer, errors.New("unknown enum value")
	}
}

type PhonesMode struct {
	v string
}

func (e PhonesMode) ToWire() string {
	return e.v
}
func (e PhonesMode) String() string {
	switch e.v {
	case "0":
		return "PFLMode"
	case "1":
		return "PersonalMixMode"
	default:
		panic("unknown enum value")
	}
}

var (
	PFLMode         = PhonesMode{"0"}
	PersonalMixMode = PhonesMode{"1"}
)

func ToPhonesMode(v string) (PhonesMode, error) {
	switch v {
	case "0":
		return PFLMode, nil
	case "1":
		return PersonalMixMode, nil
	default:
		return PFLMode, errors.New("unknown enum value")
	}
}

type BridgeMode struct {
	v string
}

func (e BridgeMode) ToWire() string {
	return e.v
}
func (e BridgeMode) String() string {
	switch e.v {
	case "NONE":
		return "DontBridge"
	case "LAN_NAT":
		return "BridgeLANAndNAT"
	default:
		panic("unknown enum value")
	}
}

var (
	DontBridge      = BridgeMode{"NONE"}
	BridgeLANAndNAT = BridgeMode{"LAN_NAT"}
)

func ToBridgeMode(v string) (BridgeMode, error) {
	switch v {
	case "NONE":
		return DontBridge, nil
	case "LAN_NAT":
		return BridgeLANAndNAT, nil
	default:
		return DontBridge, errors.New("unknown enum value")
	}
}

type MFX1And2Channels struct {
	v string
}

func (e MFX1And2Channels) ToWire() string {
	return e.v
}
func (e MFX1And2Channels) String() string {
	switch e.v {
	case "0":
		return "Channel5And6"
	case "1":
		return "Channel7And8"
	default:
		panic("unknown enum value")
	}
}

var (
	Channel5And6 = MFX1And2Channels{"0"}
	Channel7And8 = MFX1And2Channels{"1"}
)

func ToMFX1And2Channels(v string) (MFX1And2Channels, error) {
	switch v {
	case "0":
		return Channel5And6, nil
	case "1":
		return Channel7And8, nil
	default:
		return Channel5And6, errors.New("unknown enum value")
	}
}

type Channel10Source struct {
	v string
}

func (e Channel10Source) ToWire() string {
	return e.v
}
func (e Channel10Source) String() string {
	switch e.v {
	case "0":
		return "AnalogInput10"
	case "1":
		return "AnalogInput9"
	default:
		panic("unknown enum value")
	}
}

var (
	AnalogInput10 = Channel10Source{"0"}
	AnalogInput9  = Channel10Source{"1"}
)

func ToChannel10Source(v string) (Channel10Source, error) {
	switch v {
	case "0":
		return AnalogInput10, nil
	case "1":
		return AnalogInput9, nil
	default:
		return AnalogInput10, errors.New("unknown enum value")
	}
}

type Channel17Source struct {
	v string
}

func (e Channel17Source) ToWire() string {
	return e.v
}
func (e Channel17Source) String() string {
	switch e.v {
	case "0":
		return "AnalogInput17"
	case "1":
		return "PlayerTrack3"
	default:
		panic("unknown enum value")
	}
}

var (
	AnalogInput17 = Channel17Source{"0"}
	PlayerTrack3  = Channel17Source{"1"}
)

func ToChannel17Source(v string) (Channel17Source, error) {
	switch v {
	case "0":
		return AnalogInput17, nil
	case "1":
		return PlayerTrack3, nil
	default:
		return AnalogInput17, errors.New("unknown enum value")
	}
}

type Channel18Source struct {
	v string
}

func (e Channel18Source) ToWire() string {
	return e.v
}
func (e Channel18Source) String() string {
	switch e.v {
	case "0":
		return "AnalogInput18"
	case "1":
		return "PlayerTrack4"
	case "2":
		return "Metronome"
	default:
		panic("unknown enum value")
	}
}

var (
	AnalogInput18 = Channel18Source{"0"}
	PlayerTrack4  = Channel18Source{"1"}
	Metronome     = Channel18Source{"2"}
)

func ToChannel18Source(v string) (Channel18Source, error) {
	switch v {
	case "0":
		return AnalogInput18, nil
	case "1":
		return PlayerTrack4, nil
	case "2":
		return Metronome, nil
	default:
		return AnalogInput18, errors.New("unknown enum value")
	}
}

type AssignDCAGroup struct {
	v float32
}

func (e AssignDCAGroup) ToWire() float32 {
	return e.v
}
func (e AssignDCAGroup) String() string {
	switch e.v {
	case 0:
		return "AssignDCANone"
	case 1:
		return "AssignDCA1"
	case 2:
		return "AssignDCA2"
	case 3:
		return "AssignDCA3"
	case 4:
		return "AssignDCA4"
	default:
		panic("unknown enum value")
	}
}

var (
	AssignDCANone = AssignDCAGroup{0}
	AssignDCA1    = AssignDCAGroup{1}
	AssignDCA2    = AssignDCAGroup{2}
	AssignDCA3    = AssignDCAGroup{3}
	AssignDCA4    = AssignDCAGroup{4}
)

func ToAssignDCAGroup(v float32) (AssignDCAGroup, error) {
	switch v {
	case 0:
		return AssignDCANone, nil
	case 1:
		return AssignDCA1, nil
	case 2:
		return AssignDCA2, nil
	case 3:
		return AssignDCA3, nil
	case 4:
		return AssignDCA4, nil
	default:
		return AssignDCANone, errors.New("unknown enum value")
	}
}

type RCFM18Server struct {
	connections                                   map[string]*Connection
	MangleReturnAddress                           func(net.Addr) net.Addr
	OnFirmwareVersionRequest                      func(c *Context) (string, error)
	FirmwareVersionMessageHandler                 func(c *Context) error
	OnTargetIDRequest                             func(c *Context) (string, error)
	TargetIDMessageHandler                        func(c *Context) error
	OnSerialNumberRequest                         func(c *Context) (string, error)
	SerialNumberMessageHandler                    func(c *Context) error
	OnOptimizedMeterSendingSet                    func(c *Context, b1 bool) error
	OptimizedMeterSendingMessageHandler           func(c *Context) error
	OnDateTimeSet                                 func(c *Context, s1 string) error
	DateTimeMessageHandler                        func(c *Context) error
	OnLinkedStereoChannelsRequest                 func(c *Context, s1 string) error
	LinkedStereoChannelsMessageHandler            func(c *Context) error
	OnLinkStereoChannelSet                        func(c *Context, control int, b1 bool) error
	LinkStereoChannelMessageHandler               func(c *Context) error
	OnGlobalMuteSet                               func(c *Context, b1 bool) error
	GlobalMuteMessageHandler                      func(c *Context) error
	OnMainMuteSet                                 func(c *Context, b1 bool) error
	MainMuteMessageHandler                        func(c *Context) error
	OnMainOutSet                                  func(c *Context, f1 float32) error
	MainOutMessageHandler                         func(c *Context) error
	OnFaderSet                                    func(c *Context, page int, channel int, f1 float32) error
	FaderMessageHandler                           func(c *Context) error
	OnMuteSet                                     func(c *Context, page int, channel int, b1 bool) error
	MuteMessageHandler                            func(c *Context) error
	OnInputPanSet                                 func(c *Context, channel int, f1 float32) error
	InputPanMessageHandler                        func(c *Context) error
	OnPhonesSet                                   func(c *Context, channel int, b1 bool) error
	PhonesMessageHandler                          func(c *Context) error
	OnInputPreGainSet                             func(c *Context, channel int, f1 float32) error
	InputPreGainMessageHandler                    func(c *Context) error
	OnInputGainLevelAndBoostSet                   func(c *Context, channel int, f1 float32) error
	InputGainLevelAndBoostMessageHandler          func(c *Context) error
	OnInputGainLevelSet                           func(c *Context, channel int, f1 float32) error
	InputGainLevelMessageHandler                  func(c *Context) error
	OnInputPreTrimSet                             func(c *Context, channel int, f1 float32) error
	InputPreTrimMessageHandler                    func(c *Context) error
	OnInputInvertPhaseSet                         func(c *Context, channel int, b1 bool) error
	InputInvertPhaseMessageHandler                func(c *Context) error
	OnInputPreFreqSet                             func(c *Context, channel int, f1 float32) error
	InputPreFreqMessageHandler                    func(c *Context) error
	OnInputHPFSet                                 func(c *Context, channel int, b1 bool) error
	InputHPFMessageHandler                        func(c *Context) error
	OnInputPrePostGainSet                         func(c *Context, page int, channel int, b1 bool) error
	InputPrePostGainMessageHandler                func(c *Context) error
	OnEQModeSet                                   func(c *Context, channel int, E1 EQMode) error
	EQModeMessageHandler                          func(c *Context) error
	OnAdvancedEQSet                               func(c *Context, channel int, b1 bool) error
	AdvancedEQMessageHandler                      func(c *Context) error
	OnEQSet                                       func(c *Context, channel int, b1 bool) error
	EQMessageHandler                              func(c *Context) error
	OnEQBand1GainSet                              func(c *Context, channel int, f1 float32) error
	EQBand1GainMessageHandler                     func(c *Context) error
	OnEQBand1FreqSet                              func(c *Context, channel int, f1 float32) error
	EQBand1FreqMessageHandler                     func(c *Context) error
	OnEQBand1SmoothnessSet                        func(c *Context, channel int, E1 EQSmoothness) error
	EQBand1SmoothnessMessageHandler               func(c *Context) error
	OnEQBand2GainSet                              func(c *Context, channel int, f1 float32) error
	EQBand2GainMessageHandler                     func(c *Context) error
	OnEQBand2FreqSet                              func(c *Context, channel int, f1 float32) error
	EQBand2FreqMessageHandler                     func(c *Context) error
	OnEQBand2QSet                                 func(c *Context, channel int, f1 float32) error
	EQBand2QMessageHandler                        func(c *Context) error
	OnEQBand3GainSet                              func(c *Context, channel int, f1 float32) error
	EQBand3GainMessageHandler                     func(c *Context) error
	OnEQBand3FreqSet                              func(c *Context, channel int, f1 float32) error
	EQBand3FreqMessageHandler                     func(c *Context) error
	OnEQBand3QSet                                 func(c *Context, channel int, f1 float32) error
	EQBand3QMessageHandler                        func(c *Context) error
	OnEQBand4GainSet                              func(c *Context, channel int, f1 float32) error
	EQBand4GainMessageHandler                     func(c *Context) error
	OnEQBand4FreqSet                              func(c *Context, channel int, f1 float32) error
	EQBand4FreqMessageHandler                     func(c *Context) error
	OnEQBand4SmoothnessSet                        func(c *Context, channel int, E1 EQSmoothness) error
	EQBand4SmoothnessMessageHandler               func(c *Context) error
	OnAuxOutSet                                   func(c *Context, channel int, f1 float32) error
	AuxOutMessageHandler                          func(c *Context) error
	OnAuxOutMuteSet                               func(c *Context, channel int, b1 bool) error
	AuxOutMuteMessageHandler                      func(c *Context) error
	OnAuxOutEQSet                                 func(c *Context, channel int, b1 bool) error
	AuxOutEQMessageHandler                        func(c *Context) error
	OnPhonesOutSet                                func(c *Context, f1 float32) error
	PhonesOutMessageHandler                       func(c *Context) error
	OnPhonesMuteSet                               func(c *Context, b1 bool) error
	PhonesMuteMessageHandler                      func(c *Context) error
	OnFXsMuteSet                                  func(c *Context, b1 bool) error
	FXsMuteMessageHandler                         func(c *Context) error
	OnFXSendSet                                   func(c *Context, page int, f1 float32) error
	FXSendMessageHandler                          func(c *Context) error
	OnFXSendMuteSet                               func(c *Context, page int, b1 bool) error
	FXSendMuteMessageHandler                      func(c *Context) error
	OnInputNameSet                                func(c *Context, control int, s1 string) error
	InputNameMessageHandler                       func(c *Context) error
	OnAuxOutNameSet                               func(c *Context, control int, s1 string) error
	AuxOutNameMessageHandler                      func(c *Context) error
	OnStartupModeSet                              func(c *Context, S1 StartupMode) error
	StartupModeMessageHandler                     func(c *Context) error
	OnRestoreDefaultSet                           func(c *Context) error
	RestoreDefaultMessageHandler                  func(c *Context) error
	OnPhonesModeSet                               func(c *Context, P1 PhonesMode) error
	PhonesModeMessageHandler                      func(c *Context) error
	OnBridgeModeSet                               func(c *Context, B1 BridgeMode) error
	BridgeModeMessageHandler                      func(c *Context) error
	OnPhantom1To4Set                              func(c *Context, b1 bool) error
	Phantom1To4MessageHandler                     func(c *Context) error
	OnPhantom5To8Set                              func(c *Context, b1 bool) error
	Phantom5To8MessageHandler                     func(c *Context) error
	OnSnapshotIncludesMFX1PatchRecallSet          func(c *Context, b1 bool) error
	SnapshotIncludesMFX1PatchRecallMessageHandler func(c *Context) error
	OnSnapshotIncludesMFX2PatchRecallSet          func(c *Context, b1 bool) error
	SnapshotIncludesMFX2PatchRecallMessageHandler func(c *Context) error
	OnSnapshotIncludesMFX3PatchRecallSet          func(c *Context, b1 bool) error
	SnapshotIncludesMFX3PatchRecallMessageHandler func(c *Context) error
	OnSnapshotIncludesMFX4PatchRecallSet          func(c *Context, b1 bool) error
	SnapshotIncludesMFX4PatchRecallMessageHandler func(c *Context) error
	OnMFX1And2ChannelsSet                         func(c *Context, M1 MFX1And2Channels) error
	MFX1And2ChannelsMessageHandler                func(c *Context) error
	OnMFX1BoostSet                                func(c *Context, b1 bool) error
	MFX1BoostMessageHandler                       func(c *Context) error
	OnMFX2BoostSet                                func(c *Context, b1 bool) error
	MFX2BoostMessageHandler                       func(c *Context) error
	OnMFX3BoostSet                                func(c *Context, b1 bool) error
	MFX3BoostMessageHandler                       func(c *Context) error
	OnMFX4BoostSet                                func(c *Context, b1 bool) error
	MFX4BoostMessageHandler                       func(c *Context) error
	OnCh10SourceSet                               func(c *Context, C1 Channel10Source) error
	Ch10SourceMessageHandler                      func(c *Context) error
	OnCh17SourceSet                               func(c *Context, C1 Channel17Source) error
	Ch17SourceMessageHandler                      func(c *Context) error
	OnCh18SourceSet                               func(c *Context, C1 Channel18Source) error
	Ch18SourceMessageHandler                      func(c *Context) error
	OnDCAMuteSet                                  func(c *Context, channel int, b1 bool) error
	DCAMuteMessageHandler                         func(c *Context) error
	OnAssignDCAGroupSet                           func(c *Context, channel int, A1 AssignDCAGroup) error
	AssignDCAGroupMessageHandler                  func(c *Context) error
	OnPollRequest                                 func(c *Context) error
	PollMessageHandler                            func(c *Context) error
}

func NewServer() *RCFM18Server {
	s := &RCFM18Server{}
	s.connections = map[string]*Connection{}
	s.MangleReturnAddress = func(src net.Addr) net.Addr {
		host, _, err := net.SplitHostPort(src.String())
		if err != nil {
			panic(err)
		}
		a, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, "9000"))
		if err != nil {
			panic(err)
		}
		return a
	}
	s.OnFirmwareVersionRequest = func(c *Context) (string, error) {
		fmt.Println("OnFirmwareVersionRequest()\n")
		return "02.10.0225", nil
	}
	s.FirmwareVersionMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		if s.OnFirmwareVersionRequest == nil {
			return ErrNoHandler
		}
		s1, err2 := s.OnFirmwareVersionRequest(cx)
		if err2 != nil {
			return err2
		}
		if cx.reply != nil {
			cx.reply.Append(s1)
		}
		return nil
	}
	s.OnTargetIDRequest = func(c *Context) (string, error) {
		fmt.Println("OnTargetIDRequest()\n")
		return "0", nil // M18
	}
	s.TargetIDMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		if s.OnTargetIDRequest == nil {
			return ErrNoHandler
		}
		s1, err2 := s.OnTargetIDRequest(cx)
		if err2 != nil {
			return err2
		}
		if cx.reply != nil {
			cx.reply.Append(s1)
		}
		return nil
	}
	s.OnSerialNumberRequest = func(c *Context) (string, error) {
		fmt.Println("OnSerialNumberRequest()\n")
		return "fake m18", nil
	}
	s.SerialNumberMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		if s.OnSerialNumberRequest == nil {
			return ErrNoHandler
		}
		s1, err2 := s.OnSerialNumberRequest(cx)
		if err2 != nil {
			return err2
		}
		if cx.reply != nil {
			cx.reply.Append(s1)
		}
		return nil
	}
	s.OnOptimizedMeterSendingSet = func(c *Context, b1 bool) error {
		fmt.Printf("OnOptimizedMeterSendingSet(%v)\n", b1)
		c.Connection().OptimizedMeterSending = b1
		c.DontReply()
		return nil
	}
	s.OptimizedMeterSendingMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnOptimizedMeterSendingSet == nil {
			return ErrNoHandler
		}
		err2 := s.OnOptimizedMeterSendingSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnDateTimeSet = func(c *Context, s1 string) error {
		fmt.Printf("OnDateTimeSet(%v)\n", s1)
		// %Y%m%d%H%M.%s
		c.DontReply()
		return nil
	}
	s.DateTimeMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		s1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnDateTimeSet == nil {
			return ErrNoHandler
		}
		err2 := s.OnDateTimeSet(cx, s1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnLinkedStereoChannelsRequest = func(c *Context, s1 string) error {
		fmt.Printf("OnLinkedStereoChannelsRequest(%v)\n", s1)
		c.DontReply() // TBD
		return nil
	}
	s.LinkedStereoChannelsMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		s1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnLinkedStereoChannelsRequest == nil {
			return ErrNoHandler
		}
		err2 := s.OnLinkedStereoChannelsRequest(cx, s1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnLinkStereoChannelSet = func(c *Context, control int, b1 bool) error {
		fmt.Printf("OnLinkStereoChannelSet(%v, %v)\n", control, b1)
		return nil
	}
	s.LinkStereoChannelMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		control := cx.control
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnLinkStereoChannelSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnLinkStereoChannelSet(cx, control, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnGlobalMuteSet = func(c *Context, b1 bool) error {
		fmt.Printf("OnGlobalMuteSet(%v)\n", b1)
		return nil
	}
	s.GlobalMuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnGlobalMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnGlobalMuteSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnMainMuteSet = func(c *Context, b1 bool) error {
		fmt.Printf("OnMainMuteSet(%v)\n", b1)
		return nil
	}
	s.MainMuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnMainMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnMainMuteSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnMainOutSet = func(c *Context, f1 float32) error {
		fmt.Printf("OnMainOutSet(%v)\n", f1)
		return nil
	}
	s.MainOutMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnMainOutSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnMainOutSet(cx, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnFaderSet = func(c *Context, page int, channel int, f1 float32) error {
		fmt.Printf("OnFaderSet(%v, %v, %v)\n", page, channel, f1)
		return nil
	}
	s.FaderMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		page := cx.page
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnFaderSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnFaderSet(cx, page, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnMuteSet = func(c *Context, page int, channel int, b1 bool) error {
		fmt.Printf("OnMuteSet(%v, %v, %v)\n", page, channel, b1)
		return nil
	}
	s.MuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		page := cx.page
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnMuteSet(cx, page, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnInputPanSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnInputPanSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.InputPanMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnInputPanSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnInputPanSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnPhonesSet = func(c *Context, channel int, b1 bool) error {
		fmt.Printf("OnPhonesSet(%v, %v)\n", channel, b1)
		return nil
	}
	s.PhonesMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnPhonesSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnPhonesSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnInputPreGainSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnInputPreGainSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.InputPreGainMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnInputPreGainSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnInputPreGainSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnInputGainLevelAndBoostSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnInputGainLevelAndBoostSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.InputGainLevelAndBoostMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnInputGainLevelAndBoostSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnInputGainLevelAndBoostSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnInputGainLevelSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnInputGainLevelSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.InputGainLevelMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnInputGainLevelSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnInputGainLevelSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnInputPreTrimSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnInputPreTrimSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.InputPreTrimMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnInputPreTrimSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnInputPreTrimSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnInputInvertPhaseSet = func(c *Context, channel int, b1 bool) error {
		fmt.Printf("OnInputInvertPhaseSet(%v, %v)\n", channel, b1)
		return nil
	}
	s.InputInvertPhaseMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnInputInvertPhaseSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnInputInvertPhaseSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnInputPreFreqSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnInputPreFreqSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.InputPreFreqMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnInputPreFreqSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnInputPreFreqSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnInputHPFSet = func(c *Context, channel int, b1 bool) error {
		fmt.Printf("OnInputHPFSet(%v, %v)\n", channel, b1)
		return nil
	}
	s.InputHPFMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnInputHPFSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnInputHPFSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnInputPrePostGainSet = func(c *Context, page int, channel int, b1 bool) error {
		fmt.Printf("OnInputPrePostGainSet(%v, %v, %v)\n", page, channel, b1)
		return nil
	}
	s.InputPrePostGainMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		page := cx.page
		channel := cx.channel
		page -= 4
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnInputPrePostGainSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnInputPrePostGainSet(cx, page, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnEQModeSet = func(c *Context, channel int, E1 EQMode) error {
		fmt.Printf("OnEQModeSet(%v, %v)\n", channel, E1)
		return nil
	}
	s.EQModeMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fE1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		E1, err := ToEQMode(fE1)
		if err != nil {
			return fmt.Errorf("invalid value for EQMode: %v", err)
		}
		if s.OnEQModeSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnEQModeSet(cx, channel, E1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnAdvancedEQSet = func(c *Context, channel int, b1 bool) error {
		fmt.Printf("OnAdvancedEQSet(%v, %v)\n", channel, b1)
		return nil
	}
	s.AdvancedEQMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnAdvancedEQSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnAdvancedEQSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnEQSet = func(c *Context, channel int, b1 bool) error {
		fmt.Printf("OnEQSet(%v, %v)\n", channel, b1)
		return nil
	}
	s.EQMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnEQSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnEQSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnEQBand1GainSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnEQBand1GainSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.EQBand1GainMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnEQBand1GainSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnEQBand1GainSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnEQBand1FreqSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnEQBand1FreqSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.EQBand1FreqMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnEQBand1FreqSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnEQBand1FreqSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnEQBand1SmoothnessSet = func(c *Context, channel int, E1 EQSmoothness) error {
		fmt.Printf("OnEQBand1SmoothnessSet(%v, %v)\n", channel, E1)
		return nil
	}
	s.EQBand1SmoothnessMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fE1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		E1, err := ToEQSmoothness(fE1)
		if err != nil {
			return fmt.Errorf("invalid value for EQSmoothness: %v", err)
		}
		if s.OnEQBand1SmoothnessSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnEQBand1SmoothnessSet(cx, channel, E1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnEQBand2GainSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnEQBand2GainSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.EQBand2GainMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnEQBand2GainSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnEQBand2GainSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnEQBand2FreqSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnEQBand2FreqSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.EQBand2FreqMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnEQBand2FreqSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnEQBand2FreqSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnEQBand2QSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnEQBand2QSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.EQBand2QMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnEQBand2QSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnEQBand2QSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnEQBand3GainSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnEQBand3GainSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.EQBand3GainMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnEQBand3GainSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnEQBand3GainSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnEQBand3FreqSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnEQBand3FreqSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.EQBand3FreqMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnEQBand3FreqSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnEQBand3FreqSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnEQBand3QSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnEQBand3QSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.EQBand3QMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnEQBand3QSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnEQBand3QSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnEQBand4GainSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnEQBand4GainSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.EQBand4GainMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnEQBand4GainSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnEQBand4GainSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnEQBand4FreqSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnEQBand4FreqSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.EQBand4FreqMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnEQBand4FreqSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnEQBand4FreqSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnEQBand4SmoothnessSet = func(c *Context, channel int, E1 EQSmoothness) error {
		fmt.Printf("OnEQBand4SmoothnessSet(%v, %v)\n", channel, E1)
		return nil
	}
	s.EQBand4SmoothnessMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fE1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		E1, err := ToEQSmoothness(fE1)
		if err != nil {
			return fmt.Errorf("invalid value for EQSmoothness: %v", err)
		}
		if s.OnEQBand4SmoothnessSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnEQBand4SmoothnessSet(cx, channel, E1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnAuxOutSet = func(c *Context, channel int, f1 float32) error {
		fmt.Printf("OnAuxOutSet(%v, %v)\n", channel, f1)
		return nil
	}
	s.AuxOutMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnAuxOutSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnAuxOutSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnAuxOutMuteSet = func(c *Context, channel int, b1 bool) error {
		fmt.Printf("OnAuxOutMuteSet(%v, %v)\n", channel, b1)
		return nil
	}
	s.AuxOutMuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnAuxOutMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnAuxOutMuteSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnAuxOutEQSet = func(c *Context, channel int, b1 bool) error {
		fmt.Printf("OnAuxOutEQSet(%v, %v)\n", channel, b1)
		return nil
	}
	s.AuxOutEQMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnAuxOutEQSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnAuxOutEQSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnPhonesOutSet = func(c *Context, f1 float32) error {
		fmt.Printf("OnPhonesOutSet(%v)\n", f1)
		return nil
	}
	s.PhonesOutMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnPhonesOutSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnPhonesOutSet(cx, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnPhonesMuteSet = func(c *Context, b1 bool) error {
		fmt.Printf("OnPhonesMuteSet(%v)\n", b1)
		return nil
	}
	s.PhonesMuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnPhonesMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnPhonesMuteSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnFXsMuteSet = func(c *Context, b1 bool) error {
		fmt.Printf("OnFXsMuteSet(%v)\n", b1)
		return nil
	}
	s.FXsMuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnFXsMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnFXsMuteSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnFXSendSet = func(c *Context, page int, f1 float32) error {
		fmt.Printf("OnFXSendSet(%v, %v)\n", page, f1)
		return nil
	}
	s.FXSendMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		page := cx.page
		page -= 1
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnFXSendSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnFXSendSet(cx, page, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnFXSendMuteSet = func(c *Context, page int, b1 bool) error {
		fmt.Printf("OnFXSendMuteSet(%v, %v)\n", page, b1)
		return nil
	}
	s.FXSendMuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		page := cx.page
		page -= 1
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnFXSendMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnFXSendMuteSet(cx, page, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnInputNameSet = func(c *Context, control int, s1 string) error {
		fmt.Printf("OnInputNameSet(%v, %v)\n", control, s1)
		return nil
	}
	s.InputNameMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		control := cx.control
		control -= 100
		s1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnInputNameSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnInputNameSet(cx, control, s1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnAuxOutNameSet = func(c *Context, control int, s1 string) error {
		fmt.Printf("OnAuxOutNameSet(%v, %v)\n", control, s1)
		return nil
	}
	s.AuxOutNameMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		control := cx.control
		control -= 120
		s1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		if s.OnAuxOutNameSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnAuxOutNameSet(cx, control, s1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnStartupModeSet = func(c *Context, S1 StartupMode) error {
		fmt.Printf("OnStartupModeSet(%v)\n", S1)
		return nil
	}
	s.StartupModeMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sS1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		S1, err := ToStartupMode(sS1)
		if err != nil {
			return fmt.Errorf("invalid value for StartupMode: %v", err)
		}
		if s.OnStartupModeSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnStartupModeSet(cx, S1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnRestoreDefaultSet = func(c *Context) error {
		fmt.Println("OnRestoreDefaultSet()\n")
		return nil
	}
	s.RestoreDefaultMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		if s.OnRestoreDefaultSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnRestoreDefaultSet(cx)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnPhonesModeSet = func(c *Context, P1 PhonesMode) error {
		fmt.Printf("OnPhonesModeSet(%v)\n", P1)
		return nil
	}
	s.PhonesModeMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sP1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		P1, err := ToPhonesMode(sP1)
		if err != nil {
			return fmt.Errorf("invalid value for PhonesMode: %v", err)
		}
		if s.OnPhonesModeSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnPhonesModeSet(cx, P1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnBridgeModeSet = func(c *Context, B1 BridgeMode) error {
		fmt.Printf("OnBridgeModeSet(%v)\n", B1)
		return nil
	}
	s.BridgeModeMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sB1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		B1, err := ToBridgeMode(sB1)
		if err != nil {
			return fmt.Errorf("invalid value for BridgeMode: %v", err)
		}
		if s.OnBridgeModeSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnBridgeModeSet(cx, B1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnPhantom1To4Set = func(c *Context, b1 bool) error {
		fmt.Printf("OnPhantom1To4Set(%v)\n", b1)
		return nil
	}
	s.Phantom1To4MessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnPhantom1To4Set == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnPhantom1To4Set(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnPhantom5To8Set = func(c *Context, b1 bool) error {
		fmt.Printf("OnPhantom5To8Set(%v)\n", b1)
		return nil
	}
	s.Phantom5To8MessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnPhantom5To8Set == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnPhantom5To8Set(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnSnapshotIncludesMFX1PatchRecallSet = func(c *Context, b1 bool) error {
		fmt.Printf("OnSnapshotIncludesMFX1PatchRecallSet(%v)\n", b1)
		return nil
	}
	s.SnapshotIncludesMFX1PatchRecallMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnSnapshotIncludesMFX1PatchRecallSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnSnapshotIncludesMFX1PatchRecallSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnSnapshotIncludesMFX2PatchRecallSet = func(c *Context, b1 bool) error {
		fmt.Printf("OnSnapshotIncludesMFX2PatchRecallSet(%v)\n", b1)
		return nil
	}
	s.SnapshotIncludesMFX2PatchRecallMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnSnapshotIncludesMFX2PatchRecallSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnSnapshotIncludesMFX2PatchRecallSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnSnapshotIncludesMFX3PatchRecallSet = func(c *Context, b1 bool) error {
		fmt.Printf("OnSnapshotIncludesMFX3PatchRecallSet(%v)\n", b1)
		return nil
	}
	s.SnapshotIncludesMFX3PatchRecallMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnSnapshotIncludesMFX3PatchRecallSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnSnapshotIncludesMFX3PatchRecallSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnSnapshotIncludesMFX4PatchRecallSet = func(c *Context, b1 bool) error {
		fmt.Printf("OnSnapshotIncludesMFX4PatchRecallSet(%v)\n", b1)
		return nil
	}
	s.SnapshotIncludesMFX4PatchRecallMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnSnapshotIncludesMFX4PatchRecallSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnSnapshotIncludesMFX4PatchRecallSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnMFX1And2ChannelsSet = func(c *Context, M1 MFX1And2Channels) error {
		fmt.Printf("OnMFX1And2ChannelsSet(%v)\n", M1)
		return nil
	}
	s.MFX1And2ChannelsMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sM1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		M1, err := ToMFX1And2Channels(sM1)
		if err != nil {
			return fmt.Errorf("invalid value for MFX1And2Channels: %v", err)
		}
		if s.OnMFX1And2ChannelsSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnMFX1And2ChannelsSet(cx, M1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnMFX1BoostSet = func(c *Context, b1 bool) error {
		fmt.Printf("OnMFX1BoostSet(%v)\n", b1)
		return nil
	}
	s.MFX1BoostMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnMFX1BoostSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnMFX1BoostSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnMFX2BoostSet = func(c *Context, b1 bool) error {
		fmt.Printf("OnMFX2BoostSet(%v)\n", b1)
		return nil
	}
	s.MFX2BoostMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnMFX2BoostSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnMFX2BoostSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnMFX3BoostSet = func(c *Context, b1 bool) error {
		fmt.Printf("OnMFX3BoostSet(%v)\n", b1)
		return nil
	}
	s.MFX3BoostMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnMFX3BoostSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnMFX3BoostSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnMFX4BoostSet = func(c *Context, b1 bool) error {
		fmt.Printf("OnMFX4BoostSet(%v)\n", b1)
		return nil
	}
	s.MFX4BoostMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnMFX4BoostSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnMFX4BoostSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnCh10SourceSet = func(c *Context, C1 Channel10Source) error {
		fmt.Printf("OnCh10SourceSet(%v)\n", C1)
		return nil
	}
	s.Ch10SourceMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sC1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		C1, err := ToChannel10Source(sC1)
		if err != nil {
			return fmt.Errorf("invalid value for Channel10Source: %v", err)
		}
		if s.OnCh10SourceSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnCh10SourceSet(cx, C1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnCh17SourceSet = func(c *Context, C1 Channel17Source) error {
		fmt.Printf("OnCh17SourceSet(%v)\n", C1)
		return nil
	}
	s.Ch17SourceMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sC1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		C1, err := ToChannel17Source(sC1)
		if err != nil {
			return fmt.Errorf("invalid value for Channel17Source: %v", err)
		}
		if s.OnCh17SourceSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnCh17SourceSet(cx, C1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnCh18SourceSet = func(c *Context, C1 Channel18Source) error {
		fmt.Printf("OnCh18SourceSet(%v)\n", C1)
		return nil
	}
	s.Ch18SourceMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sC1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		C1, err := ToChannel18Source(sC1)
		if err != nil {
			return fmt.Errorf("invalid value for Channel18Source: %v", err)
		}
		if s.OnCh18SourceSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnCh18SourceSet(cx, C1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnDCAMuteSet = func(c *Context, channel int, b1 bool) error {
		fmt.Printf("OnDCAMuteSet(%v, %v)\n", channel, b1)
		return nil
	}
	s.DCAMuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if s.OnDCAMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnDCAMuteSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnAssignDCAGroupSet = func(c *Context, channel int, A1 AssignDCAGroup) error {
		fmt.Printf("OnAssignDCAGroupSet(%v, %v)\n", channel, A1)
		return nil
	}
	s.AssignDCAGroupMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fA1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		A1, err := ToAssignDCAGroup(fA1)
		if err != nil {
			return fmt.Errorf("invalid value for AssignDCAGroup: %v", err)
		}
		if s.OnAssignDCAGroupSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := s.OnAssignDCAGroupSet(cx, channel, A1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	s.OnPollRequest = func(c *Context) error {
		fmt.Println("OnPollRequest()\n")
		c.DontReply()
		return nil
	}
	s.PollMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		if s.OnPollRequest == nil {
			return ErrNoHandler
		}
		err2 := s.OnPollRequest(cx)
		if err2 != nil {
			return err2
		}
		return nil
	}
	return s
}

func (s *RCFM18Server) Dispatch(cx *Context) error {
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 100 {
		return s.FirmwareVersionMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 146 {
		return s.TargetIDMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 275 {
		return s.SerialNumberMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 172 {
		return s.OptimizedMeterSendingMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 272 {
		return s.DateTimeMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 27 && cx.channel == 0 && cx.control == 78 {
		return s.LinkedStereoChannelsMessageHandler(cx)
	}
	if cx.addressType == "ut" && cx.page == 27 && cx.channel == 0 && cx.control >= 0 && cx.control <= 9 {
		return s.LinkStereoChannelMessageHandler(cx)
	}
	if cx.addressType == "mg" && cx.page == 0 && cx.channel == 0 && cx.control == 0 {
		return s.GlobalMuteMessageHandler(cx)
	}
	if cx.addressType == "mt" && cx.page == 0 && cx.channel == 0 && cx.control == 0 {
		return s.MainMuteMessageHandler(cx)
	}
	if cx.addressType == "fd" && cx.page == 0 && cx.channel == 0 && cx.control == 0 {
		return s.MainOutMessageHandler(cx)
	}
	if cx.addressType == "fd" && cx.page >= 1 && cx.page <= 10 && cx.channel >= 1 && cx.channel <= 23 && cx.control == 0 {
		return s.FaderMessageHandler(cx)
	}
	if cx.addressType == "mt" && cx.page >= 1 && cx.page <= 10 && cx.channel >= 1 && cx.channel <= 23 && cx.control == 0 {
		return s.MuteMessageHandler(cx)
	}
	if cx.addressType == "bl" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 0 {
		return s.InputPanMessageHandler(cx)
	}
	if cx.addressType == "sl" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 0 {
		return s.PhonesMessageHandler(cx)
	}
	if cx.addressType == "pg" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 8 && cx.control == 0 {
		return s.InputPreGainMessageHandler(cx)
	}
	if cx.addressType == "pg" && cx.page == 1 && cx.channel >= 9 && cx.channel <= 10 && cx.control == 0 {
		return s.InputGainLevelAndBoostMessageHandler(cx)
	}
	if cx.addressType == "pg" && cx.page == 1 && cx.channel >= 11 && cx.channel <= 18 && cx.control == 0 {
		return s.InputGainLevelMessageHandler(cx)
	}
	if cx.addressType == "pg" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 10 && cx.control == 1 {
		return s.InputPreTrimMessageHandler(cx)
	}
	if cx.addressType == "ph" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 18 && cx.control == 0 {
		return s.InputInvertPhaseMessageHandler(cx)
	}
	if cx.addressType == "hp" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 18 && cx.control == 0 {
		return s.InputPreFreqMessageHandler(cx)
	}
	if cx.addressType == "hp" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 18 && cx.control == 1 {
		return s.InputHPFMessageHandler(cx)
	}
	if cx.addressType == "pr" && cx.page >= 9 && cx.page <= 10 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 0 {
		return s.InputPrePostGainMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 0 {
		return s.EQModeMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 1 {
		return s.AdvancedEQMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 2 {
		return s.EQMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 3 {
		return s.EQBand1GainMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 4 {
		return s.EQBand1FreqMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 5 {
		return s.EQBand1SmoothnessMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 7 {
		return s.EQBand2GainMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 8 {
		return s.EQBand2FreqMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 9 {
		return s.EQBand2QMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 10 {
		return s.EQBand3GainMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 11 {
		return s.EQBand3FreqMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 12 {
		return s.EQBand3QMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 13 {
		return s.EQBand4GainMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 14 {
		return s.EQBand4FreqMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 15 {
		return s.EQBand4SmoothnessMessageHandler(cx)
	}
	if cx.addressType == "fd" && cx.page == 26 && cx.channel >= 1 && cx.channel <= 6 && cx.control == 0 {
		return s.AuxOutMessageHandler(cx)
	}
	if cx.addressType == "mt" && cx.page == 26 && cx.channel >= 1 && cx.channel <= 6 && cx.control == 0 {
		return s.AuxOutMuteMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 26 && cx.channel >= 1 && cx.channel <= 6 && cx.control == 2 {
		return s.AuxOutEQMessageHandler(cx)
	}
	if cx.addressType == "fd" && cx.page == 13 && cx.channel == 0 && cx.control == 0 {
		return s.PhonesOutMessageHandler(cx)
	}
	if cx.addressType == "mt" && cx.page == 13 && cx.channel == 0 && cx.control == 0 {
		return s.PhonesMuteMessageHandler(cx)
	}
	if cx.addressType == "mg" && cx.page == 0 && cx.channel == 1 && cx.control == 0 {
		return s.FXsMuteMessageHandler(cx)
	}
	if cx.addressType == "fd" && cx.page >= 2 && cx.page <= 3 && cx.channel == 0 && cx.control == 0 {
		return s.FXSendMessageHandler(cx)
	}
	if cx.addressType == "mt" && cx.page >= 2 && cx.page <= 3 && cx.channel == 0 && cx.control == 0 {
		return s.FXSendMuteMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control >= 101 && cx.control <= 120 {
		return s.InputNameMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control >= 121 && cx.control <= 126 {
		return s.AuxOutNameMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 209 {
		return s.StartupModeMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 222 {
		return s.RestoreDefaultMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 70 {
		return s.PhonesModeMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 22 {
		return s.BridgeModeMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 65 {
		return s.Phantom1To4MessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 66 {
		return s.Phantom5To8MessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 241 {
		return s.SnapshotIncludesMFX1PatchRecallMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 242 {
		return s.SnapshotIncludesMFX2PatchRecallMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 243 {
		return s.SnapshotIncludesMFX3PatchRecallMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 244 {
		return s.SnapshotIncludesMFX4PatchRecallMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 245 {
		return s.MFX1And2ChannelsMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 246 {
		return s.MFX1BoostMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 247 {
		return s.MFX2BoostMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 248 {
		return s.MFX3BoostMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 249 {
		return s.MFX4BoostMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 300 {
		return s.Ch10SourceMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 282 {
		return s.Ch17SourceMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 283 {
		return s.Ch18SourceMessageHandler(cx)
	}
	if cx.addressType == "mg" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 4 && cx.control == 0 {
		return s.DCAMuteMessageHandler(cx)
	}
	if cx.addressType == "sg" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 23 && cx.control == 0 {
		return s.AssignDCAGroupMessageHandler(cx)
	}
	if cx.addressType == "pl" && cx.page == 0 && cx.channel == 0 && cx.control == 0 {
		return s.PollMessageHandler(cx)
	}
	return ErrNoHandler
}

type RCFM18Client struct {
	conn                                          *Connection
	Send                                          func(p osc.Packet) error
	OnFirmwareVersionResponse                     func(c *Context, s1 string) error
	FirmwareVersionMessageHandler                 func(c *Context) error
	OnTargetIDResponse                            func(c *Context, s1 string) error
	TargetIDMessageHandler                        func(c *Context) error
	OnSerialNumberResponse                        func(c *Context, s1 string) error
	SerialNumberMessageHandler                    func(c *Context) error
	OnOptimizedMeterSendingSet                    func(c *Context) error
	OptimizedMeterSendingMessageHandler           func(c *Context) error
	OnDateTimeSet                                 func(c *Context) error
	DateTimeMessageHandler                        func(c *Context) error
	OnLinkedStereoChannelsResponse                func(c *Context) error
	LinkedStereoChannelsMessageHandler            func(c *Context) error
	OnLinkStereoChannelSet                        func(c *Context, control int, b1 bool) error
	LinkStereoChannelMessageHandler               func(c *Context) error
	OnGlobalMuteSet                               func(c *Context, b1 bool) error
	GlobalMuteMessageHandler                      func(c *Context) error
	OnMainMuteSet                                 func(c *Context, b1 bool) error
	MainMuteMessageHandler                        func(c *Context) error
	OnMainOutSet                                  func(c *Context, f1 float32) error
	MainOutMessageHandler                         func(c *Context) error
	OnFaderSet                                    func(c *Context, page int, channel int, f1 float32) error
	FaderMessageHandler                           func(c *Context) error
	OnMuteSet                                     func(c *Context, page int, channel int, b1 bool) error
	MuteMessageHandler                            func(c *Context) error
	OnInputPanSet                                 func(c *Context, channel int, f1 float32) error
	InputPanMessageHandler                        func(c *Context) error
	OnPhonesSet                                   func(c *Context, channel int, b1 bool) error
	PhonesMessageHandler                          func(c *Context) error
	OnInputPreGainSet                             func(c *Context, channel int, f1 float32) error
	InputPreGainMessageHandler                    func(c *Context) error
	OnInputGainLevelAndBoostSet                   func(c *Context, channel int, f1 float32) error
	InputGainLevelAndBoostMessageHandler          func(c *Context) error
	OnInputGainLevelSet                           func(c *Context, channel int, f1 float32) error
	InputGainLevelMessageHandler                  func(c *Context) error
	OnInputPreTrimSet                             func(c *Context, channel int, f1 float32) error
	InputPreTrimMessageHandler                    func(c *Context) error
	OnInputInvertPhaseSet                         func(c *Context, channel int, b1 bool) error
	InputInvertPhaseMessageHandler                func(c *Context) error
	OnInputPreFreqSet                             func(c *Context, channel int, f1 float32) error
	InputPreFreqMessageHandler                    func(c *Context) error
	OnInputHPFSet                                 func(c *Context, channel int, b1 bool) error
	InputHPFMessageHandler                        func(c *Context) error
	OnInputPrePostGainSet                         func(c *Context, page int, channel int, b1 bool) error
	InputPrePostGainMessageHandler                func(c *Context) error
	OnEQModeSet                                   func(c *Context, channel int, E1 EQMode) error
	EQModeMessageHandler                          func(c *Context) error
	OnAdvancedEQSet                               func(c *Context, channel int, b1 bool) error
	AdvancedEQMessageHandler                      func(c *Context) error
	OnEQSet                                       func(c *Context, channel int, b1 bool) error
	EQMessageHandler                              func(c *Context) error
	OnEQBand1GainSet                              func(c *Context, channel int, f1 float32) error
	EQBand1GainMessageHandler                     func(c *Context) error
	OnEQBand1FreqSet                              func(c *Context, channel int, f1 float32) error
	EQBand1FreqMessageHandler                     func(c *Context) error
	OnEQBand1SmoothnessSet                        func(c *Context, channel int, E1 EQSmoothness) error
	EQBand1SmoothnessMessageHandler               func(c *Context) error
	OnEQBand2GainSet                              func(c *Context, channel int, f1 float32) error
	EQBand2GainMessageHandler                     func(c *Context) error
	OnEQBand2FreqSet                              func(c *Context, channel int, f1 float32) error
	EQBand2FreqMessageHandler                     func(c *Context) error
	OnEQBand2QSet                                 func(c *Context, channel int, f1 float32) error
	EQBand2QMessageHandler                        func(c *Context) error
	OnEQBand3GainSet                              func(c *Context, channel int, f1 float32) error
	EQBand3GainMessageHandler                     func(c *Context) error
	OnEQBand3FreqSet                              func(c *Context, channel int, f1 float32) error
	EQBand3FreqMessageHandler                     func(c *Context) error
	OnEQBand3QSet                                 func(c *Context, channel int, f1 float32) error
	EQBand3QMessageHandler                        func(c *Context) error
	OnEQBand4GainSet                              func(c *Context, channel int, f1 float32) error
	EQBand4GainMessageHandler                     func(c *Context) error
	OnEQBand4FreqSet                              func(c *Context, channel int, f1 float32) error
	EQBand4FreqMessageHandler                     func(c *Context) error
	OnEQBand4SmoothnessSet                        func(c *Context, channel int, E1 EQSmoothness) error
	EQBand4SmoothnessMessageHandler               func(c *Context) error
	OnAuxOutSet                                   func(c *Context, channel int, f1 float32) error
	AuxOutMessageHandler                          func(c *Context) error
	OnAuxOutMuteSet                               func(c *Context, channel int, b1 bool) error
	AuxOutMuteMessageHandler                      func(c *Context) error
	OnAuxOutEQSet                                 func(c *Context, channel int, b1 bool) error
	AuxOutEQMessageHandler                        func(c *Context) error
	OnPhonesOutSet                                func(c *Context, f1 float32) error
	PhonesOutMessageHandler                       func(c *Context) error
	OnPhonesMuteSet                               func(c *Context, b1 bool) error
	PhonesMuteMessageHandler                      func(c *Context) error
	OnFXsMuteSet                                  func(c *Context, b1 bool) error
	FXsMuteMessageHandler                         func(c *Context) error
	OnFXSendSet                                   func(c *Context, page int, f1 float32) error
	FXSendMessageHandler                          func(c *Context) error
	OnFXSendMuteSet                               func(c *Context, page int, b1 bool) error
	FXSendMuteMessageHandler                      func(c *Context) error
	OnInputNameSet                                func(c *Context, control int, s1 string) error
	InputNameMessageHandler                       func(c *Context) error
	OnAuxOutNameSet                               func(c *Context, control int, s1 string) error
	AuxOutNameMessageHandler                      func(c *Context) error
	OnStartupModeSet                              func(c *Context, S1 StartupMode) error
	StartupModeMessageHandler                     func(c *Context) error
	OnRestoreDefaultSet                           func(c *Context) error
	RestoreDefaultMessageHandler                  func(c *Context) error
	OnPhonesModeSet                               func(c *Context, P1 PhonesMode) error
	PhonesModeMessageHandler                      func(c *Context) error
	OnBridgeModeSet                               func(c *Context, B1 BridgeMode) error
	BridgeModeMessageHandler                      func(c *Context) error
	OnPhantom1To4Set                              func(c *Context, b1 bool) error
	Phantom1To4MessageHandler                     func(c *Context) error
	OnPhantom5To8Set                              func(c *Context, b1 bool) error
	Phantom5To8MessageHandler                     func(c *Context) error
	OnSnapshotIncludesMFX1PatchRecallSet          func(c *Context, b1 bool) error
	SnapshotIncludesMFX1PatchRecallMessageHandler func(c *Context) error
	OnSnapshotIncludesMFX2PatchRecallSet          func(c *Context, b1 bool) error
	SnapshotIncludesMFX2PatchRecallMessageHandler func(c *Context) error
	OnSnapshotIncludesMFX3PatchRecallSet          func(c *Context, b1 bool) error
	SnapshotIncludesMFX3PatchRecallMessageHandler func(c *Context) error
	OnSnapshotIncludesMFX4PatchRecallSet          func(c *Context, b1 bool) error
	SnapshotIncludesMFX4PatchRecallMessageHandler func(c *Context) error
	OnMFX1And2ChannelsSet                         func(c *Context, M1 MFX1And2Channels) error
	MFX1And2ChannelsMessageHandler                func(c *Context) error
	OnMFX1BoostSet                                func(c *Context, b1 bool) error
	MFX1BoostMessageHandler                       func(c *Context) error
	OnMFX2BoostSet                                func(c *Context, b1 bool) error
	MFX2BoostMessageHandler                       func(c *Context) error
	OnMFX3BoostSet                                func(c *Context, b1 bool) error
	MFX3BoostMessageHandler                       func(c *Context) error
	OnMFX4BoostSet                                func(c *Context, b1 bool) error
	MFX4BoostMessageHandler                       func(c *Context) error
	OnCh10SourceSet                               func(c *Context, C1 Channel10Source) error
	Ch10SourceMessageHandler                      func(c *Context) error
	OnCh17SourceSet                               func(c *Context, C1 Channel17Source) error
	Ch17SourceMessageHandler                      func(c *Context) error
	OnCh18SourceSet                               func(c *Context, C1 Channel18Source) error
	Ch18SourceMessageHandler                      func(c *Context) error
	OnDCAMuteSet                                  func(c *Context, channel int, b1 bool) error
	DCAMuteMessageHandler                         func(c *Context) error
	OnAssignDCAGroupSet                           func(c *Context, channel int, A1 AssignDCAGroup) error
	AssignDCAGroupMessageHandler                  func(c *Context) error
	OnPollResponse                                func(c *Context) error
	PollMessageHandler                            func(c *Context) error
	OnMeters                                      func(c *Context, inputs, player, fxReturns, main, fxSends, aux, phones, recording []float32) error
	MetersMessageHandler                          func(c *Context) error
}

func (c *RCFM18Client) setupCallbacks() {
	c.FirmwareVersionMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		s1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnFirmwareVersionResponse == nil {
			return ErrNoHandler
		}
		err2 := c.OnFirmwareVersionResponse(cx, s1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.TargetIDMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		s1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnTargetIDResponse == nil {
			return ErrNoHandler
		}
		err2 := c.OnTargetIDResponse(cx, s1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.SerialNumberMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		s1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnSerialNumberResponse == nil {
			return ErrNoHandler
		}
		err2 := c.OnSerialNumberResponse(cx, s1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.OptimizedMeterSendingMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		if c.OnOptimizedMeterSendingSet == nil {
			return ErrNoHandler
		}
		err2 := c.OnOptimizedMeterSendingSet(cx)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.DateTimeMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		if c.OnDateTimeSet == nil {
			return ErrNoHandler
		}
		err2 := c.OnDateTimeSet(cx)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.LinkedStereoChannelsMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		if c.OnLinkedStereoChannelsResponse == nil {
			return ErrNoHandler
		}
		err2 := c.OnLinkedStereoChannelsResponse(cx)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.LinkStereoChannelMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		control := cx.control
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnLinkStereoChannelSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnLinkStereoChannelSet(cx, control, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.GlobalMuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnGlobalMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnGlobalMuteSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.MainMuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnMainMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnMainMuteSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.MainOutMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnMainOutSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnMainOutSet(cx, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.FaderMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		page := cx.page
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnFaderSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnFaderSet(cx, page, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.MuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		page := cx.page
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnMuteSet(cx, page, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.InputPanMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnInputPanSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnInputPanSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.PhonesMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnPhonesSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnPhonesSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.InputPreGainMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnInputPreGainSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnInputPreGainSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.InputGainLevelAndBoostMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnInputGainLevelAndBoostSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnInputGainLevelAndBoostSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.InputGainLevelMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnInputGainLevelSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnInputGainLevelSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.InputPreTrimMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnInputPreTrimSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnInputPreTrimSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.InputInvertPhaseMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnInputInvertPhaseSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnInputInvertPhaseSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.InputPreFreqMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnInputPreFreqSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnInputPreFreqSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.InputHPFMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnInputHPFSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnInputHPFSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.InputPrePostGainMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		page := cx.page
		channel := cx.channel
		page -= 4
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnInputPrePostGainSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnInputPrePostGainSet(cx, page, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.EQModeMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fE1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		E1, err := ToEQMode(fE1)
		if err != nil {
			return fmt.Errorf("invalid value for EQMode: %v", err)
		}
		if c.OnEQModeSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnEQModeSet(cx, channel, E1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.AdvancedEQMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnAdvancedEQSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnAdvancedEQSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.EQMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnEQSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnEQSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.EQBand1GainMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnEQBand1GainSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnEQBand1GainSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.EQBand1FreqMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnEQBand1FreqSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnEQBand1FreqSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.EQBand1SmoothnessMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fE1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		E1, err := ToEQSmoothness(fE1)
		if err != nil {
			return fmt.Errorf("invalid value for EQSmoothness: %v", err)
		}
		if c.OnEQBand1SmoothnessSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnEQBand1SmoothnessSet(cx, channel, E1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.EQBand2GainMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnEQBand2GainSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnEQBand2GainSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.EQBand2FreqMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnEQBand2FreqSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnEQBand2FreqSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.EQBand2QMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnEQBand2QSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnEQBand2QSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.EQBand3GainMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnEQBand3GainSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnEQBand3GainSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.EQBand3FreqMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnEQBand3FreqSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnEQBand3FreqSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.EQBand3QMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnEQBand3QSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnEQBand3QSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.EQBand4GainMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnEQBand4GainSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnEQBand4GainSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.EQBand4FreqMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnEQBand4FreqSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnEQBand4FreqSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.EQBand4SmoothnessMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fE1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		E1, err := ToEQSmoothness(fE1)
		if err != nil {
			return fmt.Errorf("invalid value for EQSmoothness: %v", err)
		}
		if c.OnEQBand4SmoothnessSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnEQBand4SmoothnessSet(cx, channel, E1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.AuxOutMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnAuxOutSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnAuxOutSet(cx, channel, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.AuxOutMuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnAuxOutMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnAuxOutMuteSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.AuxOutEQMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnAuxOutEQSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnAuxOutEQSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.PhonesOutMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnPhonesOutSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnPhonesOutSet(cx, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.PhonesMuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnPhonesMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnPhonesMuteSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.FXsMuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnFXsMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnFXsMuteSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.FXSendMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		page := cx.page
		page -= 1
		f1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnFXSendSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnFXSendSet(cx, page, f1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.FXSendMuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		page := cx.page
		page -= 1
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnFXSendMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnFXSendMuteSet(cx, page, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.InputNameMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		control := cx.control
		control -= 100
		s1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnInputNameSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnInputNameSet(cx, control, s1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.AuxOutNameMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		control := cx.control
		control -= 120
		s1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		if c.OnAuxOutNameSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnAuxOutNameSet(cx, control, s1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.StartupModeMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sS1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		S1, err := ToStartupMode(sS1)
		if err != nil {
			return fmt.Errorf("invalid value for StartupMode: %v", err)
		}
		if c.OnStartupModeSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnStartupModeSet(cx, S1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.RestoreDefaultMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		if c.OnRestoreDefaultSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnRestoreDefaultSet(cx)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.PhonesModeMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sP1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		P1, err := ToPhonesMode(sP1)
		if err != nil {
			return fmt.Errorf("invalid value for PhonesMode: %v", err)
		}
		if c.OnPhonesModeSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnPhonesModeSet(cx, P1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.BridgeModeMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sB1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		B1, err := ToBridgeMode(sB1)
		if err != nil {
			return fmt.Errorf("invalid value for BridgeMode: %v", err)
		}
		if c.OnBridgeModeSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnBridgeModeSet(cx, B1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.Phantom1To4MessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnPhantom1To4Set == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnPhantom1To4Set(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.Phantom5To8MessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnPhantom5To8Set == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnPhantom5To8Set(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.SnapshotIncludesMFX1PatchRecallMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnSnapshotIncludesMFX1PatchRecallSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnSnapshotIncludesMFX1PatchRecallSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.SnapshotIncludesMFX2PatchRecallMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnSnapshotIncludesMFX2PatchRecallSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnSnapshotIncludesMFX2PatchRecallSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.SnapshotIncludesMFX3PatchRecallMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnSnapshotIncludesMFX3PatchRecallSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnSnapshotIncludesMFX3PatchRecallSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.SnapshotIncludesMFX4PatchRecallMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnSnapshotIncludesMFX4PatchRecallSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnSnapshotIncludesMFX4PatchRecallSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.MFX1And2ChannelsMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sM1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		M1, err := ToMFX1And2Channels(sM1)
		if err != nil {
			return fmt.Errorf("invalid value for MFX1And2Channels: %v", err)
		}
		if c.OnMFX1And2ChannelsSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnMFX1And2ChannelsSet(cx, M1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.MFX1BoostMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnMFX1BoostSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnMFX1BoostSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.MFX2BoostMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnMFX2BoostSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnMFX2BoostSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.MFX3BoostMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnMFX3BoostSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnMFX3BoostSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.MFX4BoostMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sb1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := stringToBool(sb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnMFX4BoostSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnMFX4BoostSet(cx, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.Ch10SourceMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sC1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		C1, err := ToChannel10Source(sC1)
		if err != nil {
			return fmt.Errorf("invalid value for Channel10Source: %v", err)
		}
		if c.OnCh10SourceSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnCh10SourceSet(cx, C1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.Ch17SourceMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sC1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		C1, err := ToChannel17Source(sC1)
		if err != nil {
			return fmt.Errorf("invalid value for Channel17Source: %v", err)
		}
		if c.OnCh17SourceSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnCh17SourceSet(cx, C1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.Ch18SourceMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		sC1, ok := cx.msg.Arguments[0].(string)
		if !ok {
			return fmt.Errorf("message should have had string argument, had %T", cx.msg.Arguments[0])
		}
		C1, err := ToChannel18Source(sC1)
		if err != nil {
			return fmt.Errorf("invalid value for Channel18Source: %v", err)
		}
		if c.OnCh18SourceSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnCh18SourceSet(cx, C1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.DCAMuteMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fb1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		b1, err := floatToBool(fb1)
		if err != nil {
			return fmt.Errorf("invalid value for bool: %v", err)
		}
		if c.OnDCAMuteSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnDCAMuteSet(cx, channel, b1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.AssignDCAGroupMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		channel := cx.channel
		fA1, ok := cx.msg.Arguments[0].(float32)
		if !ok {
			return fmt.Errorf("message should have had float32 argument, had %T", cx.msg.Arguments[0])
		}
		A1, err := ToAssignDCAGroup(fA1)
		if err != nil {
			return fmt.Errorf("invalid value for AssignDCAGroup: %v", err)
		}
		if c.OnAssignDCAGroupSet == nil {
			return ErrNoHandler
		}
		cx.broadcast = true
		cx.reply = cx.msg
		err2 := c.OnAssignDCAGroupSet(cx, channel, A1)
		if err2 != nil {
			return err2
		}
		return nil
	}
	c.PollMessageHandler = func(cx *Context) error {
		if len(cx.msg.Arguments) != 1 {
			return fmt.Errorf("message had %d argument(s), expected 1", len(cx.msg.Arguments))
		}
		if c.OnPollResponse == nil {
			return ErrNoHandler
		}
		err2 := c.OnPollResponse(cx)
		if err2 != nil {
			return err2
		}
		return nil
	}
}

func (c *RCFM18Client) Dispatch(cx *Context) error {
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 100 {
		return c.FirmwareVersionMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 146 {
		return c.TargetIDMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 275 {
		return c.SerialNumberMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 172 {
		return c.OptimizedMeterSendingMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 272 {
		return c.DateTimeMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 27 && cx.channel == 0 && cx.control == 78 {
		return c.LinkedStereoChannelsMessageHandler(cx)
	}
	if cx.addressType == "ut" && cx.page == 27 && cx.channel == 0 && cx.control >= 0 && cx.control <= 9 {
		return c.LinkStereoChannelMessageHandler(cx)
	}
	if cx.addressType == "mg" && cx.page == 0 && cx.channel == 0 && cx.control == 0 {
		return c.GlobalMuteMessageHandler(cx)
	}
	if cx.addressType == "mt" && cx.page == 0 && cx.channel == 0 && cx.control == 0 {
		return c.MainMuteMessageHandler(cx)
	}
	if cx.addressType == "fd" && cx.page == 0 && cx.channel == 0 && cx.control == 0 {
		return c.MainOutMessageHandler(cx)
	}
	if cx.addressType == "fd" && cx.page >= 1 && cx.page <= 10 && cx.channel >= 1 && cx.channel <= 23 && cx.control == 0 {
		return c.FaderMessageHandler(cx)
	}
	if cx.addressType == "mt" && cx.page >= 1 && cx.page <= 10 && cx.channel >= 1 && cx.channel <= 23 && cx.control == 0 {
		return c.MuteMessageHandler(cx)
	}
	if cx.addressType == "bl" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 0 {
		return c.InputPanMessageHandler(cx)
	}
	if cx.addressType == "sl" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 0 {
		return c.PhonesMessageHandler(cx)
	}
	if cx.addressType == "pg" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 8 && cx.control == 0 {
		return c.InputPreGainMessageHandler(cx)
	}
	if cx.addressType == "pg" && cx.page == 1 && cx.channel >= 9 && cx.channel <= 10 && cx.control == 0 {
		return c.InputGainLevelAndBoostMessageHandler(cx)
	}
	if cx.addressType == "pg" && cx.page == 1 && cx.channel >= 11 && cx.channel <= 18 && cx.control == 0 {
		return c.InputGainLevelMessageHandler(cx)
	}
	if cx.addressType == "pg" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 10 && cx.control == 1 {
		return c.InputPreTrimMessageHandler(cx)
	}
	if cx.addressType == "ph" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 18 && cx.control == 0 {
		return c.InputInvertPhaseMessageHandler(cx)
	}
	if cx.addressType == "hp" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 18 && cx.control == 0 {
		return c.InputPreFreqMessageHandler(cx)
	}
	if cx.addressType == "hp" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 18 && cx.control == 1 {
		return c.InputHPFMessageHandler(cx)
	}
	if cx.addressType == "pr" && cx.page >= 9 && cx.page <= 10 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 0 {
		return c.InputPrePostGainMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 0 {
		return c.EQModeMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 1 {
		return c.AdvancedEQMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 2 {
		return c.EQMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 3 {
		return c.EQBand1GainMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 4 {
		return c.EQBand1FreqMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 5 {
		return c.EQBand1SmoothnessMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 7 {
		return c.EQBand2GainMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 8 {
		return c.EQBand2FreqMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 9 {
		return c.EQBand2QMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 10 {
		return c.EQBand3GainMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 11 {
		return c.EQBand3FreqMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 12 {
		return c.EQBand3QMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 13 {
		return c.EQBand4GainMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 14 {
		return c.EQBand4FreqMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 20 && cx.control == 15 {
		return c.EQBand4SmoothnessMessageHandler(cx)
	}
	if cx.addressType == "fd" && cx.page == 26 && cx.channel >= 1 && cx.channel <= 6 && cx.control == 0 {
		return c.AuxOutMessageHandler(cx)
	}
	if cx.addressType == "mt" && cx.page == 26 && cx.channel >= 1 && cx.channel <= 6 && cx.control == 0 {
		return c.AuxOutMuteMessageHandler(cx)
	}
	if cx.addressType == "eq" && cx.page == 26 && cx.channel >= 1 && cx.channel <= 6 && cx.control == 2 {
		return c.AuxOutEQMessageHandler(cx)
	}
	if cx.addressType == "fd" && cx.page == 13 && cx.channel == 0 && cx.control == 0 {
		return c.PhonesOutMessageHandler(cx)
	}
	if cx.addressType == "mt" && cx.page == 13 && cx.channel == 0 && cx.control == 0 {
		return c.PhonesMuteMessageHandler(cx)
	}
	if cx.addressType == "mg" && cx.page == 0 && cx.channel == 1 && cx.control == 0 {
		return c.FXsMuteMessageHandler(cx)
	}
	if cx.addressType == "fd" && cx.page >= 2 && cx.page <= 3 && cx.channel == 0 && cx.control == 0 {
		return c.FXSendMessageHandler(cx)
	}
	if cx.addressType == "mt" && cx.page >= 2 && cx.page <= 3 && cx.channel == 0 && cx.control == 0 {
		return c.FXSendMuteMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control >= 101 && cx.control <= 120 {
		return c.InputNameMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control >= 121 && cx.control <= 126 {
		return c.AuxOutNameMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 209 {
		return c.StartupModeMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 222 {
		return c.RestoreDefaultMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 70 {
		return c.PhonesModeMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 22 {
		return c.BridgeModeMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 65 {
		return c.Phantom1To4MessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 66 {
		return c.Phantom5To8MessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 241 {
		return c.SnapshotIncludesMFX1PatchRecallMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 242 {
		return c.SnapshotIncludesMFX2PatchRecallMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 243 {
		return c.SnapshotIncludesMFX3PatchRecallMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 244 {
		return c.SnapshotIncludesMFX4PatchRecallMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 245 {
		return c.MFX1And2ChannelsMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 246 {
		return c.MFX1BoostMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 247 {
		return c.MFX2BoostMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 248 {
		return c.MFX3BoostMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 249 {
		return c.MFX4BoostMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 300 {
		return c.Ch10SourceMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 282 {
		return c.Ch17SourceMessageHandler(cx)
	}
	if cx.addressType == "gb" && cx.page == 22 && cx.channel == 0 && cx.control == 283 {
		return c.Ch18SourceMessageHandler(cx)
	}
	if cx.addressType == "mg" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 4 && cx.control == 0 {
		return c.DCAMuteMessageHandler(cx)
	}
	if cx.addressType == "sg" && cx.page == 1 && cx.channel >= 1 && cx.channel <= 23 && cx.control == 0 {
		return c.AssignDCAGroupMessageHandler(cx)
	}
	if cx.addressType == "pl" && cx.page == 0 && cx.channel == 0 && cx.control == 0 {
		return c.PollMessageHandler(cx)
	}
	if cx.addressType == "vmeter" && cx.page == 0 && cx.channel == 0 && cx.control == 0 {
		return c.MetersMessageHandler(cx)
	}
	return ErrNoHandler
}

func (c *RCFM18Client) RequestFirmwareVersion() error {
	msg := osc.NewMessage("/22/00/gb_100")
	msg.Append(float32(0))
	return c.Send(msg)
}

func (c *RCFM18Client) RequestTargetID() error {
	msg := osc.NewMessage("/22/00/gb_146")
	msg.Append(" ")
	return c.Send(msg)
}

func (c *RCFM18Client) RequestSerialNumber() error {
	msg := osc.NewMessage("/22/00/gb_275")
	msg.Append(" ")
	return c.Send(msg)
}

func (c *RCFM18Client) SetOptimizedMeterSending(b1 bool) error {
	msg := osc.NewMessage("/22/00/gb_172")
	msg.Append(boolToString(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetDateTime(s1 string) error {
	msg := osc.NewMessage("/22/00/gb_272")
	msg.Append(s1)
	return c.Send(msg)
}

func (c *RCFM18Client) RequestLinkedStereoChannels(s1 string) error {
	msg := osc.NewMessage("/27/00/gb_78")
	msg.Append(s1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetLinkStereoChannel(control int, b1 bool) error {
	if control < 0 || control > 9 {
		return fmt.Errorf("control out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/27/00/ut_%03d", control))
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetGlobalMute(b1 bool) error {
	msg := osc.NewMessage("/00/00/mg_000")
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetMainMute(b1 bool) error {
	msg := osc.NewMessage("/00/00/mt_000")
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetMainOut(f1 float32) error {
	msg := osc.NewMessage("/00/00/fd_000")
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetFader(page int, channel int, f1 float32) error {
	if page < 1 || page > 10 {
		return fmt.Errorf("page out of range")
	}
	if channel < 1 || channel > 23 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/%02d/%02d/fd_000", page, channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetMute(page int, channel int, b1 bool) error {
	if page < 1 || page > 10 {
		return fmt.Errorf("page out of range")
	}
	if channel < 1 || channel > 23 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/%02d/%02d/mt_000", page, channel))
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetInputPan(channel int, f1 float32) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/bl_000", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetPhones(channel int, b1 bool) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/sl_000", channel))
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetInputPreGain(channel int, f1 float32) error {
	if channel < 1 || channel > 8 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/pg_000", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetInputGainLevelAndBoost(channel int, f1 float32) error {
	if channel < 9 || channel > 10 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/pg_000", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetInputGainLevel(channel int, f1 float32) error {
	if channel < 11 || channel > 18 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/pg_000", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetInputPreTrim(channel int, f1 float32) error {
	if channel < 1 || channel > 10 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/pg_001", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetInputInvertPhase(channel int, b1 bool) error {
	if channel < 1 || channel > 18 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/ph_000", channel))
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetInputPreFreq(channel int, f1 float32) error {
	if channel < 1 || channel > 18 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/hp_000", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetInputHPF(channel int, b1 bool) error {
	if channel < 1 || channel > 18 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/hp_001", channel))
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetInputPrePostGain(page int, channel int, b1 bool) error {
	page += 4
	if page < 9 || page > 10 {
		return fmt.Errorf("page out of range")
	}
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/%02d/%02d/pr_000", page, channel))
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetEQMode(channel int, E1 EQMode) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/eq_000", channel))
	msg.Append(E1.ToWire())
	return c.Send(msg)
}

func (c *RCFM18Client) SetAdvancedEQ(channel int, b1 bool) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/eq_001", channel))
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetEQ(channel int, b1 bool) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/eq_002", channel))
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetEQBand1Gain(channel int, f1 float32) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/eq_003", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetEQBand1Freq(channel int, f1 float32) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/eq_004", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetEQBand1Smoothness(channel int, E1 EQSmoothness) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/eq_005", channel))
	msg.Append(E1.ToWire())
	return c.Send(msg)
}

func (c *RCFM18Client) SetEQBand2Gain(channel int, f1 float32) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/eq_007", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetEQBand2Freq(channel int, f1 float32) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/eq_008", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetEQBand2Q(channel int, f1 float32) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/eq_009", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetEQBand3Gain(channel int, f1 float32) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/eq_010", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetEQBand3Freq(channel int, f1 float32) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/eq_011", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetEQBand3Q(channel int, f1 float32) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/eq_012", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetEQBand4Gain(channel int, f1 float32) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/eq_013", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetEQBand4Freq(channel int, f1 float32) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/eq_014", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetEQBand4Smoothness(channel int, E1 EQSmoothness) error {
	if channel < 1 || channel > 20 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/eq_015", channel))
	msg.Append(E1.ToWire())
	return c.Send(msg)
}

func (c *RCFM18Client) SetAuxOut(channel int, f1 float32) error {
	if channel < 1 || channel > 6 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/26/%02d/fd_000", channel))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetAuxOutMute(channel int, b1 bool) error {
	if channel < 1 || channel > 6 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/26/%02d/mt_000", channel))
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetAuxOutEQ(channel int, b1 bool) error {
	if channel < 1 || channel > 6 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/26/%02d/eq_002", channel))
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetPhonesOut(f1 float32) error {
	msg := osc.NewMessage("/13/00/fd_000")
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetPhonesMute(b1 bool) error {
	msg := osc.NewMessage("/13/00/mt_000")
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetFXsMute(b1 bool) error {
	msg := osc.NewMessage("/00/01/mg_000")
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetFXSend(page int, f1 float32) error {
	page += 1
	if page < 2 || page > 3 {
		return fmt.Errorf("page out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/%02d/00/fd_000", page))
	msg.Append(f1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetFXSendMute(page int, b1 bool) error {
	page += 1
	if page < 2 || page > 3 {
		return fmt.Errorf("page out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/%02d/00/mt_000", page))
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetInputName(control int, s1 string) error {
	control += 100
	if control < 101 || control > 120 {
		return fmt.Errorf("control out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/22/00/gb_%03d", control))
	msg.Append(s1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetAuxOutName(control int, s1 string) error {
	control += 120
	if control < 121 || control > 126 {
		return fmt.Errorf("control out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/22/00/gb_%03d", control))
	msg.Append(s1)
	return c.Send(msg)
}

func (c *RCFM18Client) SetStartupMode(S1 StartupMode) error {
	msg := osc.NewMessage("/22/00/gb_209")
	msg.Append(S1.ToWire())
	return c.Send(msg)
}

func (c *RCFM18Client) SetRestoreDefault() error {
	msg := osc.NewMessage("/22/00/gb_222")
	msg.Append("1")
	return c.Send(msg)
}

func (c *RCFM18Client) SetPhonesMode(P1 PhonesMode) error {
	msg := osc.NewMessage("/22/00/gb_070")
	msg.Append(P1.ToWire())
	return c.Send(msg)
}

func (c *RCFM18Client) SetBridgeMode(B1 BridgeMode) error {
	msg := osc.NewMessage("/22/00/gb_022")
	msg.Append(B1.ToWire())
	return c.Send(msg)
}

func (c *RCFM18Client) SetPhantom1To4(b1 bool) error {
	msg := osc.NewMessage("/22/00/gb_065")
	msg.Append(boolToString(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetPhantom5To8(b1 bool) error {
	msg := osc.NewMessage("/22/00/gb_066")
	msg.Append(boolToString(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetSnapshotIncludesMFX1PatchRecall(b1 bool) error {
	msg := osc.NewMessage("/22/00/gb_241")
	msg.Append(boolToString(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetSnapshotIncludesMFX2PatchRecall(b1 bool) error {
	msg := osc.NewMessage("/22/00/gb_242")
	msg.Append(boolToString(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetSnapshotIncludesMFX3PatchRecall(b1 bool) error {
	msg := osc.NewMessage("/22/00/gb_243")
	msg.Append(boolToString(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetSnapshotIncludesMFX4PatchRecall(b1 bool) error {
	msg := osc.NewMessage("/22/00/gb_244")
	msg.Append(boolToString(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetMFX1And2Channels(M1 MFX1And2Channels) error {
	msg := osc.NewMessage("/22/00/gb_245")
	msg.Append(M1.ToWire())
	return c.Send(msg)
}

func (c *RCFM18Client) SetMFX1Boost(b1 bool) error {
	msg := osc.NewMessage("/22/00/gb_246")
	msg.Append(boolToString(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetMFX2Boost(b1 bool) error {
	msg := osc.NewMessage("/22/00/gb_247")
	msg.Append(boolToString(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetMFX3Boost(b1 bool) error {
	msg := osc.NewMessage("/22/00/gb_248")
	msg.Append(boolToString(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetMFX4Boost(b1 bool) error {
	msg := osc.NewMessage("/22/00/gb_249")
	msg.Append(boolToString(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetCh10Source(C1 Channel10Source) error {
	msg := osc.NewMessage("/22/00/gb_300")
	msg.Append(C1.ToWire())
	return c.Send(msg)
}

func (c *RCFM18Client) SetCh17Source(C1 Channel17Source) error {
	msg := osc.NewMessage("/22/00/gb_282")
	msg.Append(C1.ToWire())
	return c.Send(msg)
}

func (c *RCFM18Client) SetCh18Source(C1 Channel18Source) error {
	msg := osc.NewMessage("/22/00/gb_283")
	msg.Append(C1.ToWire())
	return c.Send(msg)
}

func (c *RCFM18Client) SetDCAMute(channel int, b1 bool) error {
	if channel < 1 || channel > 4 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/mg_000", channel))
	msg.Append(boolToFloat(b1))
	return c.Send(msg)
}

func (c *RCFM18Client) SetAssignDCAGroup(channel int, A1 AssignDCAGroup) error {
	if channel < 1 || channel > 23 {
		return fmt.Errorf("channel out of range")
	}
	msg := osc.NewMessage(fmt.Sprintf("/01/%02d/sg_000", channel))
	msg.Append(A1.ToWire())
	return c.Send(msg)
}

func (c *RCFM18Client) RequestPoll() error {
	msg := osc.NewMessage("/00/00/pl_000")
	msg.Append(float32(1))
	return c.Send(msg)
}
