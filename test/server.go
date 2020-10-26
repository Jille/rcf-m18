package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/Jille/rcf-m18/rcfm18"
)

var (
	meters = flag.Bool("meters", false, "Enable random meter updates")
)

func main() {
	log.SetFlags(log.Ltime)
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	s := rcfm18.NewServer()
	if *meters {
		go func() {
			for range time.NewTicker(200 * time.Millisecond).C {
				v := make([]float32, 41)
				for i := range v {
					v[i] = rand.Float32()
				}
				s.BroadcastMeters(v)
			}
		}()
	}
	log.Fatal(s.ListenAndServe("0.0.0.0:8000"))
}
