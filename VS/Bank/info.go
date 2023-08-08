package main

import (
	"fmt"
	"os"
	"time"
)

type Msg struct {
	Aktie  string  `json:"Aktie"`
	Anzahl int     `json:"Anzahl"`
	Price  float64 `json:"Price"`
}
type Data struct {
	Aktie      string
	Price      float64
	LastUpdate string
}
type Kunde struct {
	Kontostand float64
}

var myAktienPrice = map[string]float64{
	"AAPL": 300.0,
	"AMZN": 200.0,
	"MSFT": 180.0,
	"TSLA": 400.0,
}
var myAktienTime = map[string]string{
	"AAPL": "2006-01-02 15:04:05",
	"AMZN": "2006-01-02 15:04:05",
	"MSFT": "2006-01-02 15:04:05",
	"TSLA": "2006-01-02 15:04:05",
}

var (
	//100K
	counter          uint64
	Inventor         = 10000.0
	listInventor     []float64
	KontoStand       = 0.0
	AnzahlTCPrequest = 0
	headers          map[string]string
	method           string
	path             string
)

const (
	Topic          = "topic/transaction"
	CommitAction   = "commit"
	RollbackAction = "rollback"
)

func messeRequestperSecond() {
	// Berechne Dauer in Sekunden
	duration := time.Since(start).Seconds()
	// Berechne Anzahl der empfangenen Pakete pro Sekunde
	rps := float64(counter) / duration

	//fmt.Printf("Anzahl empfangener Pakete: %d\n", counter)
	//fmt.Printf("dauer: %v\n", duration)
	//fmt.Printf("Pakete pro Sekunde: %.2f\n", rps)

	file, err := os.Create("RequestPerSecond.txt")
	if err != nil {
		panic(err)
	}
	defer closeFile(file)
	// Schreibe eine Zeile in die Datei
	line := fmt.Sprintf("Anzahl empfangener Pakete: %d\n", counter)
	_, err = fmt.Fprintln(file, line)
	line = fmt.Sprintf("dauer: %v\n", duration)
	_, err = fmt.Fprintln(file, line)
	line = fmt.Sprintf("Pakete pro Sekunde: %.2f\n", rps)
	_, err = fmt.Fprintln(file, line)
	if err != nil {
		panic(err)
	}
}

func saveData() {
	file, err := os.Create("AktienData.txt")
	if err != nil {
		panic(err)
	}
	for index, value := range myAktienPrice {
		line := "Aktien:" + index + " Price: " + fmt.Sprintf("%.2f", value) + " last time: " + myAktienTime[index] + "\n"
		_, err = fmt.Fprint(file, line)
		if err != nil {
			panic(err)
		}
	}
	//line := fmt.Sprintf("Inventor: %.2f", Inventor)
	//_, err = fmt.Fprintln(file, line)
	//am ende schließe die file
	defer closeFile(file)
}

func saveInventor() {
	file, err := os.Create("Inventar.txt")
	if err != nil {
		panic(err)
	}
	for _, elemt := range listInventor {
		line := fmt.Sprintf("Inventar: %.2f€\n", elemt)
		_, err = file.WriteString(line)
		if err != nil {
			panic(err)
		}
	}
	defer closeFile(file)
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		return
	}
}
