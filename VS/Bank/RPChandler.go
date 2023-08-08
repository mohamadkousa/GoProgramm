package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
)

func RPCuberweisung(betrag string, bank string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(bank+":6544", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	client := NewChatServiceClient(conn)
	f, _ := strconv.ParseFloat(betrag, 8) //convert string to float
	massage := Uberweisung{Geld: f}
	ok, err := client.Uberweissung(context.Background(), &massage)
	if err != nil {
		log.Fatalf("Error when calling RPCuberweisung: %v", err)
	}
	if ok.Ok {
		println("uberweisung wurde erfolgreich ausgefuehrt ", ok.Ok)
	}
}
func RPCStonierung(bank string, betrag string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(bank+":6544", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	client := NewChatServiceClient(conn)
	f, _ := strconv.ParseFloat(betrag, 8)
	massage := Message{Body: "Gib mir mein Geld zurueck", HowMuch: f}
	respons, err1 := client.Stonieren(context.Background(), &massage)
	if err1 != nil {
		log.Fatalf("Error when calling RPCuberweisung: %v", err)
	}
	fmt.Printf("Money has been sent: %v\n", respons.Ok)
	//fur test
	//_, err = client.Ping(context.Background(), &Empty{})
	//if err != nil {
	//	log.Fatalf("Error when calling Ping-Methode: %v", err)
	//}
}
