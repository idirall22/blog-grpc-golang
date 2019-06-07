package main

import (
	"context"
	"fmt"

	"./service"
)

func main() {
	service.InitDB()

	// post := &proto.Post{Title: "post1", Author: 1,
	// 	Content: "content01", Published: false}
	// service.CreatePost(context.Background(), post)
	p, _ := service.GetSinglePost(context.Background(), 2)
	fmt.Println(p.Comments)
	// service.DeletePost(context.Background(), int(post.Id))

}
