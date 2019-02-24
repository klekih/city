package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Hello, world. I'm the city.")

	l, err := net.Listen("tcp", ":7450")
	if err != nil {
		log.Fatal(err)
	}
	go startWebSocket()
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connection accepted")
		go comm(conn)
	}
}
