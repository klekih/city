package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"github.com/tomagb/city/common"
)

func comm(c net.Conn) {
	defer c.Close()

	gob.Register(common.Report{})
	gob.Register(common.LineInfo{})

	dec := gob.NewDecoder(c)
	env := new(common.Envelope)
	err := dec.Decode(&env)
	if err != nil {
		fmt.Println("Error on decoding", err)
		return
	}

	switch env.MessageType {
	case common.SendReport:
		payload := env.Payload.(common.Report)
		if payload.ReportDetail == common.ReportOnTheLine {
			deliverLineData(payload.CurrentLine)
		} else {
			deleteLineData(payload.CurrentLine)
		}
	case common.AskForLine:
		enc := gob.NewEncoder(c)
		err := enc.Encode(common.Envelope{MessageType: common.RespondWithLine})
		payload := env.Payload.(common.LineInfo)
		fmt.Println("Recieved ask for line", payload)
		if err != nil {
			fmt.Println("Error on decoding", err)
			return
		}
	}
}
