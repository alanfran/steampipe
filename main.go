package main

import (
	"log"
	"net"
	"time"

	"github.com/alanfran/steampipe/protocol"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9407")
	if err != nil {
		log.Fatalf("Failed to start grpc server on port 9407.")
	}

	app := newApp(time.Second * 9)

	grpcServer := grpc.NewServer()
	protocol.RegisterSteamPipeServer(grpcServer, &QueryService{app})
	go grpcServer.Serve(lis)

	app.run(":80")
}
