package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"../proto"
)

//Databse Info
const (
	databaseName = "blog"
	tableName    = "posts"
)

// Errors
const (
	errCreatePost = "Can not create the post"
)

//databaseService default db used in blog service
var databaseService *sql.DB

//InitDB set the db
func InitDB(db *sql.DB) {
	databaseService = db
}

//CreatePost create a post in database
func CreatePost(ctx context.Context, post *proto.Post) error {
	dataIn := "(author, title, content, published)"
	query := fmt.Sprintf("Insert INTO $1 $2", tableName, dataIn)

	_, err := databaseService.ExecContext(ctx, query,
		&post.Author, &post.Title, &post.Content, &post.Published)

	if err != nil {
		return errors.New(errCreatePost)
	}

	return nil
}
