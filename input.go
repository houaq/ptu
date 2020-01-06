package main

import (
	"log"
	"time"
)

var dummy int

func updateIOState() {
	var states [8]uint16
	dummy++
	states[0] = uint16(dummy % 2)
	_, err := mbc.WriteSingleRegister(0x11b, states[0])
	if err != nil {
		log.Printf("update io state: %s", err.Error())
	}
}

func watchInput() {
	for {
		updateIOState()
		time.Sleep(time.Second)
	}
}
