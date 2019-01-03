package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

// Connect is the typical method used for connecting to
// a city.
func Connect() (chan Report, chan Junction) {

	var sendReportChan = make(chan Report)
	var junctionChan = make(chan Junction)

	go func() {
		for {
			select {
			case r := <-sendReportChan:
				conn, err := net.Dial("tcp", "localhost:7450")
				if err != nil {
					fmt.Println("Error on dialing", err)
					break
				}
				defer conn.Close()

				env := Envelope{
					MessageType: SendReport,
					Payload:     r}

				gob.Register(r)
				enc := gob.NewEncoder(conn)
				err = enc.Encode(env)
				if err != nil {
					fmt.Println("Error on sending data", err)
				}
			case j := <-junctionChan:
				conn, err := net.Dial("tcp", "localhost:7450")
				if err != nil {
					fmt.Println("Error on dialing", err)
					break
				}
				defer conn.Close()
				env := Envelope{
					MessageType: AskForJunction,
					Payload:     j}

				gob.Register(Junction{})
				enc := gob.NewEncoder(conn)
				err = enc.Encode(env)
				if err != nil {
					fmt.Println("Error on sending data", err)
				}

				dec := gob.NewDecoder(conn)
				env = Envelope{}
				err = dec.Decode(&env)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("Received response with junction", env)
				junctionChan <- Junction{}
			}
		}
	}()

	return sendReportChan, junctionChan
}
