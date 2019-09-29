package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// LinesData is the response for the front end service
// with lines and their data.
type LinesData struct {
	Coords  [][]float64 `json:"coords"`
	Density int         `json:"density"`
}

var addr = flag.String("addr", ":9000", "http service address")

var upgrader = websocket.Upgrader{} // use default options

var sendLineDataChan <-chan LinesData

func waitForIncomingMessages(conn *websocket.Conn) {
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading from front end:", err)
			break
		}
		log.Printf("Received from front end. type: %d  msg: %s", messageType, message)
	}
}

func city(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	go waitForIncomingMessages(c)
	for {
		select {
		case data := <-sendLineDataChan:
			respBytes := new(bytes.Buffer)
			json.NewEncoder(respBytes).Encode(data)
			err := c.WriteMessage(1, respBytes.Bytes())
			if err != nil {
				log.Println("Error writing to front end:", err)
			}
		}
	}
}

func startWebSocket(sendLineData <-chan LinesData) {
	sendLineDataChan = sendLineData
	http.HandleFunc("/city", city)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
