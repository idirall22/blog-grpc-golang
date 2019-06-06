package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

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

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "your-password"
	dbname   = "Blog"
)

//databaseService default db used in blog service
var databaseService *sql.DB

//InitDB set the db
func InitDB(db *sql.DB) {

	dataSourceName := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable",
		host, user, dbname)
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		log.Fatalln("Could not Connect to Database: %s", err)
		return
	}

	if err = db.Ping(); err != nil {
		log.Fatalln("Could not Ping to Database: %s", err)
		return
	}
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
