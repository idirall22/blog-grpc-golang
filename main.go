package main

import (
	"context"

	"./proto"
	"./service"
)

func main() {
	service.InitDB()

	post := &proto.Post{Id: 1, Title: "edited post1", Author: 1,
		Content: "edited content01", Published: true}
	// service.CreatePost(context.Background(), post)

	service.DeletePost(context.Background(), int(post.Id))

}
