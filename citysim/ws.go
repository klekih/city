package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// LinesData is the response for the front end service
// with lines and their data.
type LinesData struct {
}

var addr = flag.String("addr", "localhost:9000", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func city(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Error reading from front end:", err)
			break
		}
		log.Printf("Received from front end: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("Write to front end:", err)
			break
		}
	}
}

func startWebSocket() {
	http.HandleFunc("/city", city)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
