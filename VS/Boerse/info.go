package main

import (
	"net"
	"time"
)

type Msg struct {
	Aktie  string  `json:"Aktie"`
	Anzahl int     `json:"Anzahl"`
	Price  float64 `json:"Price"`
}

var (
	//ch      chan string
	//wg      sync.WaitGroup
	kuerzel = []string{"AAPL", "AMZN", "MSFT", "TSLA"}

	price = []float64{300.0, 200.0, 180.0, 400.0}

	addr = net.UDPAddr{Port: 6543, IP: net.ParseIP("localhost")}

	banks   map[string]time.Time
	counter uint64
)

// Default werte from
// aktienkurse(0) means price[0]
//var defaultPrice = map[string]float64{
//	"AAPL": 300.0,
//	"AMZN": 200.0,
//	"MSFT": 180.0,
//	"TSLA": 400.0,
//}
