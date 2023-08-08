package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

func calculateWin(buffer []byte) {
	var msg Msg
	if err := json.Unmarshal(buffer, &msg); err != nil {
		log.Fatal(err)
	}
	//ausgabe hat einen Einfluess auf die Performance
	//fmt.Printf("Aktie %v, Anzahl: %v, Price: %.2f\n", msg.Aktie, msg.Anzahl, msg.Price)
	// berechne das gewinn oder verlost newPrice-AltePrice>0
	gewin := (float64(msg.Anzahl) * msg.Price) - (myAktienPrice[msg.Aktie] * (float64(msg.Anzahl)))
	//if gewin > 0 {
	//	fmt.Printf("Bank gewinn in Hoehe von %.2f€\n", gewin)
	//
	//} else {
	//	fmt.Printf("Bank verlust in Hoehe von %.2f€\n", gewin)
	//}
	//aktualisiere die Preise mit den neuen erhaltene Preise
	myAktienPrice[msg.Aktie] = msg.Price
	//aktualisiere die Zeit für diese Aktien Mit dem Format "2006-01-02 15:04:05"
	t := time.Now()
	layout := "2006-01-02 15:04:05"
	str := t.Format(layout)
	myAktienTime[msg.Aktie] = fmt.Sprint(str)

	//gucke wie viel gewinn oder verlust hat ein Bank
	Inventor += gewin
	listInventor = append(listInventor, Inventor)
	//fmt.Printf("Inventor: %.2f\n", Inventor)

}

func calculate(aktie string, price string, anzahl string) {
	newPrice, _ := strconv.ParseFloat(price, 8)
	numberOfAktien, _ := strconv.ParseFloat(anzahl, 8)
	gewin := (numberOfAktien * newPrice) - (myAktienPrice[aktie] * numberOfAktien)
	if gewin > 0 {
		fmt.Printf("Bank gewinn in Hoehe von %.2f€\n", gewin)

	} else {
		fmt.Printf("Bank verlust in Hoehe von %.2f€\n", gewin)
	}
	//aktualisiere die Preise mit den neuen erhaltene Preise
	myAktienPrice[aktie] = newPrice
	//aktualisiere die Zeit für diese Aktien Mit dem Format "2006-01-02 15:04:05"
	t := time.Now()
	layout := "2006-01-02 15:04:05"
	str := t.Format(layout)
	myAktienTime[aktie] = fmt.Sprint(str)

	//gucke wie viel gewinn oder verlust hat ein Bank
	Inventor += gewin
	listInventor = append(listInventor, Inventor)
}
