package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

// Location aggregates the information about the place
// where an actor is at a certain moment in time
type Location struct {
	Long float64
	Lat  float64
}

// Report is the base type for reporting status and vectors
// to a city entity
type Report struct {
	Loc Location
}

// Junction is the message send back from the city with information
// about a junction
type Junction struct {
	Loc Location
}

const (
	// SendReport is a message passed from an actor to the city
	// indicating its status (e.g. location).
	SendReport = iota

	// AskForJunction is a message passed from an actor to the city.
	// A response is awaited.
	AskForJunction = iota

	// RespondWithJunction is a message passed from the city to
	// an actor and it contains junction data.
	RespondWithJunction = iota
)

// Envelope is the container for different messages sent back
// and forth between an actor and a city
type Envelope struct {
	MessageType int
	Payload     interface{}
}

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
