package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

func Ping() {
	var conn *grpc.ClientConn
	bank := "localhost"
	//startrpc := time.Now()
	conn, err := grpc.Dial(bank+":6544", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	clientTest := NewChatServiceClient(conn)
	//fur test
	//startPing := time.Now()
	//start = time.Now()
	_, err = clientTest.Ping(context.Background(), &Empty{})
	counter++
	messeRPC()
	//duration := time.Since(startrpc).Microseconds()
	//duration1 := time.Since(startPing).Microseconds()
	//saverpctime(duration, duration1, file)
	if err != nil {
		log.Fatalf("Error when calling Ping-Methode: %v", err)
	}
}

func saverpctime(duration int64, duration1 int64, file *os.File) {
	//file, err := os.Create("RPCtime.txt")

	line := fmt.Sprintf("RPC start Request Gesamte-duration: %d, nur Ping nach der Verbindung %v = %v\n", duration, duration1, duration-duration1)
	_, err := file.WriteString(line)
	if err != nil {
		panic(err)
	}
}

func RPCtest() {
	//time.Sleep(time.Second)
	//f, err := os.OpenFile("RPCtime.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	//if err != nil {
	//	panic(err)
	//}
	//t1 := time.Now()
	counter = 0
	for {
		Ping()
		println("ping")
	}
	//duration := time.Since(t1).Microseconds()
	//println(duration)
	//defer closeFile(f)
}

func messeRPC() {
	// Berechne Dauer in Sekunden
	duration := time.Since(start).Seconds()
	// Berechne Anzahl der empfangenen Pakete pro Sekunde
	rps := float64(counter) / float64(duration)

	file, err := os.Create("RPC-Ping.txt")
	if err != nil {
		panic(err)
	}
	defer closeFile(file)
	// Schreibe eine Zeile in die Datei
	line := fmt.Sprintf("Anzahl RPC ping: %d\n", counter)
	_, err = fmt.Fprintln(file, line)
	line = fmt.Sprintf("dauer: %v\n", duration)
	_, err = fmt.Fprintln(file, line)
	line = fmt.Sprintf("Ping pro Sekunde: %.2f\n", rps)
	_, err = fmt.Fprintln(file, line)
	if err != nil {
		panic(err)
	}
}
func messeRPC1() {
	// Berechne Dauer in Sekunden
	duration := time.Since(start).Seconds()
	// Berechne Anzahl der empfangenen Pakete pro Sekunde
	//rps := float64(counter) / float64(duration)

	file, err := os.Create("MQTT.txt")
	if err != nil {
		panic(err)
	}
	defer closeFile(file)
	// Schreibe eine Zeile in die Datei
	//line := fmt.Sprintf("Anzahl RPC ping: %d\n", counter)
	//_, err = fmt.Fprintln(file, line)
	line := fmt.Sprintf("dauer bis der Bank gerettet wurde: %v\n", duration)
	_, err = fmt.Fprintln(file, line)
	//line = fmt.Sprintf("Ping pro Sekunde: %.2f\n", rps)
	//_, err = fmt.Fprintln(file, line)
	if err != nil {
		panic(err)
	}
}
