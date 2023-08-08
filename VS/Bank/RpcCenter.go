package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RPC() {
	listener, err := net.Listen("tcp", ":6544")
	if err != nil {
		log.Fatal("cannt connect to Port :6544 ", err)
	}
	fmt.Println("RPC Server started. Listening on port 6544...")
	s := Server{}

	grpcServer := grpc.NewServer()

	RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server over port :6543 : %v", err)
	}
	CloseListener(listener)
}
