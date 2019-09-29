package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"github.com/tomagb/city/common"
)

func comm(c net.Conn, frontChan chan<- LinesData) {
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
			density := deliverLineData(payload.CurrentLine)
			frontChan <- LinesData{payload.CurrentLine, density}
		} else {
			density := deleteLineData(payload.CurrentLine)
			frontChan <- LinesData{payload.CurrentLine, density}
		}
	case common.AskForLine:
		enc := gob.NewEncoder(c)
		payload := env.Payload.(common.LineInfo)
		if err != nil {
			fmt.Println("Error on decoding", err)
			return
		}

		data := getLineData(payload.Coordinates)
		payloadBack := common.LineInfo{}.WithLine(payload.Coordinates).WithDensity(data)
		err := enc.Encode(common.Envelope{MessageType: common.RespondWithLine, Payload: payloadBack})
		if err != nil {
			fmt.Println("Error on encoding", err)
			return
		}
	}
}
