package main

import (
	"os"
	"strconv"
)

func main() {
	strBankGeld := os.Getenv("BANKRESERVE")
	if len(strBankGeld) != 0 {
		f, _ := strconv.ParseFloat(strBankGeld, 8)
		Inventor = f
	}
	go StartClient(os.Getenv("HOSTNAME"))
	go StartTCPServer() //Praktikum 2
	go StartMQTT()      //Praktikum 4
	//go RPC()
	//start = time.Now()
	//RPCtest() //Ping Test RPC
	RPC() //Praktikum 3

}
