package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"github.com/tomagb/city/common"
)

// Connect is the typical method used for connecting to
// a city.
func Connect() (chan common.Report, chan common.LineInfo) {

	var sendReportChan = make(chan common.Report)
	var lineChan = make(chan common.LineInfo)

	go func() {
		for {
			select {
			case r := <-sendReportChan:
				conn, err := net.Dial("tcp", "citysim:7450")
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
			case l := <-lineChan:
				conn, err := net.Dial("tcp", "citysim:7450")
				if err != nil {
					fmt.Println("Error on dialing", err)
					break
				}
				defer conn.Close()
				env := common.Envelope{
					MessageType: common.AskForLine,
					Payload:     l}

				gob.Register(common.LineInfo{})
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
				lineChan <- common.LineInfo{}
			}
		}
	}()

	return sendReportChan, lineChan
}
