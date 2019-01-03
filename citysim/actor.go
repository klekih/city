package citysim

import (
	"encoding/gob"
	"fmt"
	"net"
)

func comm(c net.Conn) {
	defer c.Close()
	defer fmt.Println("Connection closed")

	gob.Register(Report{})
	gob.Register(Junction{})

	dec := gob.NewDecoder(c)
	env := new(Envelope)
	err := dec.Decode(&env)
	if err != nil {
		fmt.Println("Error on decoding", err)
		return
	}

	switch env.MessageType {
	case SendReport:
		fmt.Println("Received report", env)
	case AskForJunction:
		fmt.Println("Received query on junction", env)
		enc := gob.NewEncoder(c)
		err := enc.Encode(Envelope{MessageType: RespondWithJunction})
		if err != nil {
			fmt.Println("Error on decoding", err)
			return
		}
	}
}
