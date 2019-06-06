package service

import (
	"context"

	"../proto"
)

//BlogService model used with grpc
type BlogService struct{}

//GetPost return single post
func (s *BlogService) GetPost(context.Context,
	*proto.RequestPost) (*proto.ResponsePost, error) {
	return nil, nil
}

//GetPosts return list of posts
func (s *BlogService) GetPosts(context.Context,
	*proto.PostsSearch) (*proto.Posts, error) {
	return nil, nil
}

//CreatePost allow to add a post
func (s *BlogService) CreatePost(context.Context,
	*proto.Post) (*proto.ResponsePost, error) {
	return nil, nil
}

//UpdatePost allow to update a post
func (s *BlogService) UpdatePost(context.Context,
	*proto.Post) (*proto.ResponsePost, error) {
	return nil, nil
}

//DeletePost allow to delete a post
func (s *BlogService) DeletePost(context.Context,
	*proto.RequestPost) (*proto.ResponsePost, error) {
	return nil, nil
}
