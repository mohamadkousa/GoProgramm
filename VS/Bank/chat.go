package main

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {} //diese Mehtode wird von RPC generiert um sicher zu stellen dass alle Methode implementiert sind

func (s *Server) Uberweissung(ctx context.Context, uberweisung *Uberweisung) (*Ok, error) {

	log.Println("UberweissungMethode uberweisung.Geld: ", uberweisung.Geld)
	KontoStand += uberweisung.Geld
	Inventor += uberweisung.Geld
	messeRPC1()
	return &Ok{Ok: true}, nil
}

func (s *Server) Stonieren(ctx context.Context, message *Message) (*Ok, error) {
	log.Println("Stonieren message.Body: ", message.Body, "	message.HowMuch:", message.HowMuch)
	//if Inventor >= message.HowMuch {
	//	if KontoStand < message.HowMuch {
	//		return &Ok{Ok: false}, nil
	//	} else {
	//		KontoStand = KontoStand - message.HowMuch
	//		Inventor = Inventor - message.HowMuch
	//		return &Ok{Ok: true}, nil
	//	}
	//} else {
	//	return &Ok{Ok: false}, nil
	//}// konnte man machen
	KontoStand = KontoStand - message.HowMuch
	Inventor = Inventor - message.HowMuch
	return &Ok{Ok: true}, nil

}

func (s *Server) Ping(cts context.Context, req *Empty) (*Empty, error) {

	// Diese Methode hat den Rückgabetyp void, daher gibt es nichts zurückzugeben
	//time.Sleep(time.Second*2)
	return &Empty{}, nil
}
