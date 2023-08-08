package main

import (
	"encoding/json"
	"sync/atomic"

	"fmt"
	"log"
	"net"
	"time"
)

// var mut sync.Mutex

func StartServer() {
	//start VS connection on Port 6543
	udpserver, err := net.ListenUDP("udp", &addr)
	fmt.Println("ServerUDP listening on: ", addr.Port)
	if err != nil {
		addr.Port++
		StartServer()
		log.Fatal(err)
	}
	//defer udpServer.Close()
	banks = make(map[string]time.Time) //erstellen
	done := make(chan bool)
	buf := make([]byte, 1024)
	for {
		//receive bleibt blockiert bis "Hi" from banks kommt
		n, addr1, err1 := udpserver.ReadFromUDP(buf) //_ ist die große anzahl of byte ip:addr, +err
		if err1 != nil {
			continue
		}
		fmt.Printf("Received message from %s: %s with length %v\n", addr.String(), buf[:n], n)

		if _, isMapContainsKey := banks[addr1.String()]; !isMapContainsKey {
			//key does not exist
			fmt.Println("new Bank joind")
			//für messung koennen wir das hier starten
			go response(udpserver, addr1, done)
			//go sleepAndKill(done)
		}
		//go response(udpserver, addr, done)
		banks[addr1.String()] = time.Now()

		fmt.Printf("aktive Banks: %v\n", len(banks))

		//check()// Check for inactive clients and remove them from the list

	}

}
func sleepAndKill(done chan bool) {
	//nur zum messen
	time.Sleep(time.Second * 10)
	done <- true
}

func response(udpServer *net.UDPConn, addr *net.UDPAddr, done chan bool) {
	//while true
	for {
		select {
		case <-done:
			fmt.Printf("go routine beendet anzahl geschickte Packets:%v\n", counter)
			counter = 0
			return
		default:
			//zahl between 0 and 3
			zahl := randomInt()
			//var p Msg
			p := Msg{Aktie: kuerzel[zahl], Anzahl: anzahlAktien(), Price: aktienkuse(zahl)}
			//convert msg in byte
			buf, err := json.Marshal(p)

			_, err = udpServer.WriteToUDP(buf, addr)
			if err != nil {
				fmt.Printf("couldnt send response %v", err)
			}
			atomic.AddUint64(&counter, 1) //counter fuer PacketsaNumber
		}
		//wait a second
		time.Sleep(time.Second)
	}
}

//func check() {
//
//	for {
//		if len(banks) != 0 {
//			for clientAddr, lastMessageTime := range banks {
//				if time.Since(lastMessageTime) > (time.Second)*6 {
//					fmt.Println("Removing inactive bank:", clientAddr)
//					delete(banks, clientAddr)
//					fmt.Printf("es sind %v banks verbunden\n", len(banks))
//				}
//			}
//		}
//		time.Sleep(time.Second * 3)
//	}
//
//}
