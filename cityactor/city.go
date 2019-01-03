package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"github.com/klekih/city/common"
)

// Connect is the typical method used for connecting to
// a city.
func Connect() (chan common.Report, chan common.Junction) {

	var sendReportChan = make(chan common.Report)
	var junctionChan = make(chan common.Junction)

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

				env := common.Envelope{
					MessageType: common.SendReport,
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
				env := common.Envelope{
					MessageType: common.AskForJunction,
					Payload:     j}

				gob.Register(common.Junction{})
				enc := gob.NewEncoder(conn)
				err = enc.Encode(env)
				if err != nil {
					fmt.Println("Error on sending data", err)
				}

				dec := gob.NewDecoder(conn)
				env = common.Envelope{}
				err = dec.Decode(&env)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("Received response with junction", env)
				junctionChan <- common.Junction{}
			}
		}
	}()

	return sendReportChan, junctionChan
}
