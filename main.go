package main

import (
	"context"

	"./proto"
	"./service"
)

func main() {
	service.InitDB()

	post := &proto.Post{Title: "post1", Author: 1,
		Content: "content01", Published: false}
	service.CreatePost(context.Background(), post)

}
