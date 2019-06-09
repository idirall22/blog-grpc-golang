package service

import (
	"context"

	"../proto"
)

//BlogService model used with grpc
type BlogService struct{}

//GetPost return single post
func (s *BlogService) GetPost(ctx context.Context,
	reqPost *proto.RequestPost) (*proto.ResponsePost, error) {
	post, err := GetSinglePost(ctx, int(reqPost.PostID))

	if err != nil {
		return &proto.ResponsePost{StatusCode: 404,
			Message: "Does not exist", Post: nil}, err
	}
	return &proto.ResponsePost{StatusCode: 200,
		Message: "Success", Post: post}, nil
}

//GetPosts return list of posts
func (s *BlogService) GetPosts(ctx context.Context,
	postSearch *proto.PostsSearch) (*proto.Posts, error) {

	posts, err := GetPosts(ctx, 10, int(postSearch.Page))

	if err != nil {
		return nil, err
	}
	return &proto.Posts{Posts: posts}, nil
}

//CreatePost allow to add a post
func (s *BlogService) CreatePost(ctx context.Context,
	post *proto.Post) (*proto.ResponsePost, error) {

	err := CreatePostDB(ctx, post)
	if err != nil {
		return &proto.ResponsePost{StatusCode: 400,
			Message: "Bad Request", Post: nil}, err
	}
	return &proto.ResponsePost{StatusCode: 200,
		Message: "Success", Post: post}, nil
}

//UpdatePost allow to update a post
func (s *BlogService) UpdatePost(ctx context.Context,
	post *proto.Post) (*proto.ResponsePost, error) {

	err := UpdatePost(ctx, post)
	if err != nil {
		return &proto.ResponsePost{StatusCode: 400,
			Message: "Bad Request", Post: nil}, err
	}
	return &proto.ResponsePost{StatusCode: 200,
		Message: "Success", Post: post}, nil
}

//DeletePost allow to delete a post
func (s *BlogService) DeletePost(ctx context.Context,
	reqPost *proto.RequestPost) (*proto.ResponsePost, error) {

	err := DeletePost(ctx, int(reqPost.PostID))
	if err != nil {
		return &proto.ResponsePost{StatusCode: 400,
			Message: "Bad Request", Post: nil}, err
	}
	return &proto.ResponsePost{StatusCode: 200,
		Message: "Success", Post: nil}, nil
}
