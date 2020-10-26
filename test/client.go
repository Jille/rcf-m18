package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Jille/rcf-m18/rcfm18"
)

func main() {
	log.SetFlags(log.Ltime)
	rand.Seed(time.Now().UnixNano())
	c, err := rcfm18.Dial("127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	err = c.SetInputName(3, fmt.Sprintf("Wh%dt", rand.Intn(9)))
	if err != nil {
		log.Fatal(err)
	}
}
