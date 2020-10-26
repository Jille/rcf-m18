package main

import (
	"flag"
	"log"
	"net"
	"strings"
	"sync"

	"github.com/Jille/go-osc/osc"
)

var (
	ignoreTypes = flag.String("ignore_types", "", "Ignored types")
)

func main() {
	flag.Parse()
	addr := "0.0.0.0:9000"

	fm18 := &Dumper{}
	fm18.ListenAndServe(addr)
}

type Dumper struct {
	s                     *osc.Server
	mtx                   sync.Mutex
	OptimizedMeterSending bool
	Faders                [24][11]float32
}

func shouldIgnore(addr string) bool {
	typ := addr[7:9]
	for _, i := range strings.Split(*ignoreTypes, ",") {
		if typ == i {
			return true
		}
	}
	return false
}

func (f *Dumper) ListenAndServe(addr string) error {
	f.s = &osc.Server{Addr: addr}

	f.s.Handle("*", f.Handle)

	return f.s.ListenAndServe()
}

func (f *Dumper) Handle(msg *osc.Message, src net.Addr) {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	if !shouldIgnore(msg.Address) {
		log.Printf("[%s] Packet: %s", src, msg)
	}
}
