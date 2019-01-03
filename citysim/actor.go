package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"github.com/klekih/city/common"
)

func comm(c net.Conn) {
	defer c.Close()
	defer fmt.Println("Connection closed")

	gob.Register(common.Report{})
	gob.Register(common.Junction{})

	dec := gob.NewDecoder(c)
	env := new(common.Envelope)
	err := dec.Decode(&env)
	if err != nil {
		fmt.Println("Error on decoding", err)
		return
	}

	switch env.MessageType {
	case common.SendReport:
		fmt.Println("Received report", env)
	case common.AskForJunction:
		fmt.Println("Received query on junction", env)
		enc := gob.NewEncoder(c)
		err := enc.Encode(common.Envelope{MessageType: common.RespondWithJunction})
		if err != nil {
			fmt.Println("Error on decoding", err)
			return
		}
	}
}
