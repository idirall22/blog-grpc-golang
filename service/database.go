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
	errCreatePost   = "Can not create the post"
	errUpdatePost   = "Can not update the post"
	errPostNotExist = "Post does not exist"

//Databse Info
)
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

//createPost create a post in database
func createPost(ctx context.Context, post *proto.Post) error {
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

//UpdatePost update a post in database
func UpdatePost(ctx context.Context, post *proto.Post) error {

	tx, err := databaseService.Begin()
	if err != nil {
		return errors.New(errUpdatePost)
	}

	queryUpdate := "UPDATE posts SET title=$1, content=$2, published=$3 WHERE id=$4"
	_, errP := tx.ExecContext(ctx, queryUpdate,
		&post.Title, &post.Content, &post.Published, &post.Id)

	if errP != nil {
		log.Fatal(errP)
	}
	tx.Commit()
	return nil
}
