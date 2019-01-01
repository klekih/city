package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type report struct {
	location struct {
		Long float64
		Lat  float64
	}
}

func comm(c net.Conn) {
	dec := gob.NewDecoder(c)

	for {
		r := new(report)
		err := dec.Decode(&r)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Received report", r)
	}
}
