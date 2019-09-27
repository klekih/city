package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Hello, world. I'm the city simulator.")

	l, err := net.Listen("tcp", ":7450")
	if err != nil {
		log.Fatal(err)
	}
	sendLineDataChan := make(chan LinesData)
	go startWebSocket(sendLineDataChan)
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go comm(conn)
	}
}
