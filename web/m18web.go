package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type void = struct{}

var (
	devAddr   = flag.String("device_addr", "192.168.0.18", "IP of the RCF M18")
	serveAddr = flag.String("serve_addr", "0.0.0.0:8000", "IP/port to listen on")

	upgrader = websocket.Upgrader{}

	clients = map[*WsClient]void{}
	mtx     sync.Mutex
)

type WsClient struct {
	VolCh chan []int
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Websocket upgrade failed", 500)
		return
	}
	defer ws.Close()

	wsc := &WsClient{
		VolCh: make(chan []int),
	}

	mtx.Lock()
	clients[wsc] = void{}
	mtx.Unlock()
	defer func() {
		mtx.Lock()
		delete(clients, wsc)
		mtx.Unlock()
	}()

	done := make(chan void)
	defer close(done)

	go func() {
		for {
			select {
			case msg := <-wsc.VolCh:
				p, err := json.Marshal(msg)
				if err != nil {
					panic(err)
				}
				ws.WriteMessage(websocket.TextMessage, p)
			case <-done:
				return
			}
		}
	}()

	for {
		msgType, p, err := ws.ReadMessage()
		_, _ = msgType, p
		if err != nil {
			panic(err)
		}
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request, t *template.Template) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := map[string]interface{}{}
	data["Host"] = r.Host
	data["Meters"] = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "L", "R", "FX1", "FX2", "FX2", "L", "R", "FX1", "FX2", "FX2", "AUX1", "AUX2", "AUX3", "AUX4", "AUX5", "AUX6"}
	t.Execute(w, data)
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	default:
		http.Error(w, "Method not allowed", 405)
		return
	}
	switch r.URL.Path {
	case "/":
		index := template.Must(template.ParseFiles("index.html"))
		serveTemplate(w, r, index)
	}
}

func main() {
	flag.Parse()

	go func() {
		for range time.NewTicker(time.Second).C {
			meters := make([]int, 39)
			for i := range meters {
				meters[i] = rand.Intn(100)
			}
			mtx.Lock()
			for c := range clients {
				select {
				case c.VolCh <- meters:
				default:
				}
			}
			mtx.Unlock()
		}
	}()

	http.HandleFunc("/", serveFile)
	http.HandleFunc("/ws", serveWs)
	log.Fatal(http.ListenAndServe(*serveAddr, nil))
}
