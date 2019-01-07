package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"github.com/klekih/city/common"
)

// Connect is the typical method used for connecting to
// a city.
func Connect() (chan common.Report, chan common.Line) {

	var sendReportChan = make(chan common.Report)
	var lineChan = make(chan common.Line)

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
			case j := <-lineChan:
				conn, err := net.Dial("tcp", "localhost:7450")
				if err != nil {
					fmt.Println("Error on dialing", err)
					break
				}
				defer conn.Close()
				env := common.Envelope{
					MessageType: common.AskForLine,
					Payload:     j}

				gob.Register(common.Line{})
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
				fmt.Println("Received response with line", env)
				lineChan <- common.Line{}
			}
		}
	}()

	return sendReportChan, lineChan
}
