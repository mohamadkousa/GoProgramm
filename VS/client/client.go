package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// for test the Bank to send him a data

var (
	start   time.Time
	counter = 0
)

func main() {
	start = time.Now()
	StartTCPclient()
}

func StartTCPclient() {
	for {
		conn, err := net.Dial("tcp", ":6543")
		if err != nil {
			println(counter)
			log.Fatal("can't connect to server ", err)
		}
		chunkSize := 2000
		delay := time.Second * 4
		data := make([]byte, 20000)
		text := []byte("HTTP /example HTTP1.1")
		copy(data, text) // Kopiere den Text in den Buffer

		dataWithende := addEndOfMessage(data)
		err = sendData(conn, dataWithende, chunkSize, delay)
		if err != nil {
			fmt.Println("Fehler beim Senden der Daten:", err)
			return
		}
		buffer := make([]byte, 2048)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			closeConnection(conn)
		}
		fmt.Println(string(buffer[:n]))
		defer closeConnection(conn)
		//messeRequestperSecond()
	}
}

func closeConnection(conn net.Conn) {
	counter++
	err := conn.Close()
	if err != nil {
		return
	}
}

func messeRequestperSecond() {
	// Berechne Dauer in Sekunden
	duration := time.Since(start).Seconds()
	// Berechne Anzahl der empfangenen Pakete pro Sekunde
	rps := float64(counter) / duration

	file, err := os.Create("TCPwithUDP.txt")
	if err != nil {
		panic(err)
	}
	defer closeFile(file)
	// Schreibe eine Zeile in die Datei
	line := fmt.Sprintf("Anzahl gesendete Request: %d\n", counter)
	_, err = fmt.Fprintln(file, line)
	line = fmt.Sprintf("dauer: %v\n", duration)
	_, err = fmt.Fprintln(file, line)
	line = fmt.Sprintf("Request pro Sekunde: %.2f\n", rps)
	_, err = fmt.Fprintln(file, line)
	if err != nil {
		panic(err)
	}
}
func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		return
	}
}
func sendData(conn net.Conn, data []byte, chunkSize int, delay time.Duration) error {
	totalSize := len(data)
	println("totalSize #####", totalSize)
	for i := 0; i < totalSize; i += chunkSize {
		end := i + chunkSize
		if end > totalSize {
			end = totalSize
		}

		// Sende Datenchunk
		_, err := conn.Write(data[i:end])
		if err != nil {
			return err
		}

		// Warte 1 Sekunde nach jedem 5.000 Byte
		if (i+1)%5000 == 0 {
			time.Sleep(delay)
		}
	}
	return nil
}
func addEndOfMessage(buffer []byte) []byte {
	endOfMessage := []byte("\r\n\r\n")
	return bytes.Join([][]byte{buffer, endOfMessage}, nil)
}

// Nachricht an Server senden
//message := "Hello, Server!"
//_, err = conn.Write([]byte(message))
//
//if err != nil {
//	panic(err)
//}
//
//fmt.Println("Message sent to server:", message)
