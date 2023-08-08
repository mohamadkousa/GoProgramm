package main

import (
	"math/rand"
	"time"
)

func randomInt() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(4)
}

func anzahlAktien() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return 1 + r.Intn(100)
}

func aktienkuse(startPrice int) float64 {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	MAXSCHWANKUNG := 5.0
	newPrice := price[startPrice]
	//ein floatzahl zwischen 0 und 1
	change := rand.Float64() * MAXSCHWANKUNG
	//entweder 0 oder 1
	if rand.Intn(2) == 1 {
		// Preis erhöhen
		//newPrice += change
		newPrice -= change
	} else {
		// Preis senken
		newPrice -= change
	}
	price[startPrice] = newPrice //diese Zeile ist wichtig damit ich immer die neue Preise aktualisiere aber irgendwann ist die Aktien nicht mehr wert!!
	return newPrice

}
