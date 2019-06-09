package main

import (
	"log"
	"net"

	"./proto"
	"./service"
	"google.golang.org/grpc"
)

func main() {
	listner, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
		return
	}
	server := grpc.NewServer()
	proto.RegisterBlogServicesServer(server, &service.BlogService{})
	service.InitDB()
	if errS := server.Serve(listner); errS != nil {
		log.Fatal(errS)
	}
}
