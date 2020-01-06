package main

import (
	"github.com/goburrow/modbus"
	"github.com/goburrow/serial"
	"github.com/houaq/mbserver"
	"log"
	"time"
)

var mbs mbserver.Server
var mbc modbus.Client

func initMBServer() {
	mbs := mbserver.NewServer()

	err := mbs.ListenTCP("0.0.0.0:1502")
	if err != nil {
		log.Fatalf("listen TCP error: %s\n", err.Error())
	}

	err = mbs.ListenRTU(&serial.Config{
		Address:  "/dev/tty.Bluetooth-Incoming-Port",
		BaudRate: 9600,
		DataBits: 8,
		StopBits: 1,
		Parity:   "N",
		Timeout:  10 * time.Second})
	if err != nil {
		log.Fatalf("listen RTU error: %s\n", err.Error())
	}
}

func initMBClient() {
	handler := modbus.NewTCPClientHandler("localhost:1502")
	err := handler.Connect()
	if err != nil {
		log.Fatalf("client connect error: %s\n", err.Error())
	}
	mbc = modbus.NewClient(handler)
}

func main() {
	initMBServer()
	initMBClient()

	go watchInput()

	for i := 0; i < 120; i++ {
		time.Sleep(time.Second)
	}

	mbs.Close()
}
