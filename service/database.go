package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"../proto"
	_ "github.com/lib/pq"
)

// Errors
const (
	errCreatePost = "Can not create the post"
)

//Databse Info
const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "password"
	dbname   = "blog"
)

//databaseService default db used in blog service
var databaseService *sql.DB

//InitDB set the db
func InitDB() {

	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		host, user, password, dbname)
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
	query := "Insert INTO posts (author, title, content, published)"
	query += "VALUES ($1, $2, $3, $4)"

	_, err := databaseService.ExecContext(ctx, query,
		&post.Author, &post.Title, &post.Content, &post.Published)

	if err != nil {
		log.Println(err)
		return errors.New(errCreatePost)
	}
	return nil
}
