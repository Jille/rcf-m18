package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/Jille/go-osc/osc"
	"github.com/Jille/rcf-m18/gen/out"
)

type Input struct {
	Name  string
	Left  int
	Right int
	Mic   bool
}

type Output struct {
	Name      string
	WantLeft  bool
	WantRight bool
	WantMic   bool
	AUX       int
}

var (
	inputs = []Input{
		{
			Name:  "MIC1",
			Left:  6,
			Right: 6,
			Mic:   true,
		},
		{
			Name:  "MIC2",
			Left:  5,
			Right: 5,
			Mic:   true,
		},
		{
			Name:  "Projector",
			Left:  9,
			Right: 10,
		},
		{
			Name:  "Chromecast",
			Left:  11,
			Right: 12,
		},
		{
			Name:  "Laptop",
			Left:  13,
			Right: 14,
		},
		{
			Name:  "Pulse 1",
			Left:  17,
			Right: 18,
		},
		{
			Name:  "Pulse 2",
			Left:  15,
			Right: 16,
		},
	}
	outputs = []Output{
		{
			Name:      "Marmitek L",
			WantLeft:  true,
			WantRight: true,
			AUX:       1,
		},
		{
			Name:      "Marmitek R",
			WantLeft:  true,
			WantRight: true,
			AUX:       2,
		},
		{
			Name:     "Bedroom L",
			WantLeft: true,
			AUX:      3,
		},
		{
			Name:      "Bedroom R",
			WantRight: true,
			AUX:       4,
		},
		{
			Name:     "Pulseaudio L",
			WantLeft: true,
			WantMic:  true,
			AUX:      5,
		},
		{
			Name:      "Pulseaudio R",
			WantRight: true,
			WantMic:   true,
			AUX:       6,
		},
	}
)

func setup(c *rcfm18.RCFM18Client) error {
	realSend := c.Send
	c.Send = func(p osc.Packet) error {
		time.Sleep(10 * time.Millisecond)
		return realSend(p)
	}
	if err := c.SetStartupMode(rcfm18.LastShow); err != nil {
		return err
	}
	if err := c.SetBridgeMode(rcfm18.BridgeLANAndNAT); err != nil {
		return err
	}
	if err := c.SetMFX1And2Channels(rcfm18.Channel5And6); err != nil {
		return err
	}
	defaultVolume := float32(0.7647059)
	for _, in := range inputs {
		/*
			if err := c.SetLinkStereoChannel(in.Left, false); err != nil {
				return err
			}
		*/
		if err := c.SetInputName(in.Left, in.Name); err != nil {
			return err
		}
		if !in.Mic {
			if in.Left == 9 || in.Left == 10 {
				if err := c.SetInputGainLevelAndBoost(in.Left, 2); err != nil {
					return err
				}
			} else {
				if err := c.SetInputGainLevel(in.Left, 1); err != nil {
					return err
				}
			}
		}
		for _, out := range outputs {
			vol := defaultVolume
			if in.Mic && !out.WantMic {
				vol = 0
			}
			if out.WantLeft {
				if err := c.SetFader(4+out.AUX, in.Left, vol); err != nil {
					return err
				}
			}
			if out.WantRight {
				if err := c.SetFader(4+out.AUX, in.Right, vol); err != nil {
					return err
				}
			}
		}
		vol := defaultVolume
		if in.Mic {
			vol = 0
		}
		if err := c.SetFader(1, in.Left, vol); err != nil {
			return err
		}
		if in.Left != in.Right {
			if err := c.SetInputName(in.Right, in.Name); err != nil {
				return err
			}
			if !in.Mic {
				if in.Right == 9 || in.Right == 10 {
					if err := c.SetInputGainLevelAndBoost(in.Right, 2); err != nil {
						return err
					}
				} else {
					if err := c.SetInputGainLevel(in.Right, 1); err != nil {
						return err
					}
				}
			}
			if err := c.SetFader(1, in.Right, vol); err != nil {
				return err
			}
			if err := c.SetInputPan(in.Left, 0); err != nil {
				return err
			}
			if err := c.SetInputPan(in.Right, 1); err != nil {
				return err
			}
		}
	}
	for _, out := range outputs {
		if err := c.SetAuxOutName(out.AUX, out.Name); err != nil {
			return err
		}
		if err := c.SetAuxOut(out.AUX, defaultVolume); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	log.SetFlags(log.Ltime)
	rand.Seed(time.Now().UnixNano())
	c, err := rcfm18.Dial("rcfm18.skynet.quis.cx:8000")
	if err != nil {
		log.Fatal(err)
	}
	if err := setup(c); err != nil {
		log.Fatal(err)
	}
}
