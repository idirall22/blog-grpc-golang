package main

import (
	"context"
	"fmt"
	"log"

	"../proto"
	"google.golang.org/grpc"

)
func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	client := proto.NewBlogServicesClient(conn)
	grpc.WaitForReady(true)

	post, err := client.CreatePost(context.Background(), &proto.Post{
		Author: 1, Title: "from client", Content: "yes", Published: true})

	if err != nil {
		fmt.Println("******************")
		log.Fatal(err)
		return
	}
	fmt.Println(post)
}
