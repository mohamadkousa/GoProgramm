package main

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

var (
	start time.Time
)

// StartClient
func StartClient(hostname string) {
	//connection aufbauen
	serverAddr, err := net.ResolveUDPAddr("udp", hostname+":6543")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Problem with Connection: ", err)
	}
	//anmelde Anfrage
	Sendmsg(conn) //sende hi

	// am endeSchließe die Verbindung
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)
	//just receive response
	start = time.Now()
	for {
		receive(conn)
	}

}

func receive(conn *net.UDPConn) {
	// Empfange Daten vom ServerUDP
	buffer := make([]byte, 1024)
	// Setze ein Receive-Timeout von 3 Sekunden
	timeoutDuration := time.Second * 3
	_ = conn.SetReadDeadline(time.Now().Add(timeoutDuration))

	//fmt.Fprintf(conn, "hi VS server") //send to server
	n, _, err := conn.ReadFromUDP(buffer) //n ist wichtig!!
	if err != nil {
		if counter != 0 {
			//for test
			fmt.Println("counter: ", counter)
			messeRequestperSecond()
			saveData()
			saveInventor()
			start = time.Now()
			counter = 0
		}
		//fmt.Println("Receive-Timeout ", err)
		//geht es nicht mehr da conn ist closed
		Sendmsg(conn)
		//chickPriceOFwertPapier(conn)
		return
	}
	//read just n byte from  buffer also Payload buffer[:n]
	calculateWin(buffer[:n])
	//add 1 to counter um die Anzahl von Request zu rechnen
	atomic.AddUint64(&counter, 1) //counter++

}

func Sendmsg(conn *net.UDPConn) {
	// Nachricht, die gesendet werden soll
	message := []byte("hi")

	// Sende die Nachricht
	_, err := conn.Write(message)
	if err != nil {
		fmt.Println("Fehler beim Senden der Nachricht:", err)
	}
}

//func chickPriceOFwertPapier(conn *net.UDPConn) {
//	layout := "2006-01-02 15:04:05"
//	for index, lastMessageTime := range myAktienTime {
//		t, _ := time.Parse(layout, lastMessageTime)
//		if time.Since(t) > (time.Second)*10 {
//			SendAnfrage(conn, index)
//		}
//	}
//}
