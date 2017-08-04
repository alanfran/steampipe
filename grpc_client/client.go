package main

/*
	This is an example client for the SteamPipe grpc service.
*/

import (
	"context"
	"log"

	"github.com/alanfran/steampipe/protocol"
	"google.golang.org/grpc"
)

const (
	address   = "localhost:9407"
	queryaddr = "zs.nekonet.xyz:27015"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := protocol.NewSteamPipeClient(conn)

	resp, err := client.Query(context.Background(), &protocol.Address{queryaddr})
	if err != nil {
		log.Fatalf("Error querying... %v", err)
	}

	log.Println(resp)
}
