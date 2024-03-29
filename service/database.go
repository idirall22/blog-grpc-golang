package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"../proto"
	_ "github.com/lib/pq"
)

// Errors
const (
	errBeginTX    = "Can not begin a transaction"
	errCommitTX   = "Can not commit a transaction"
	errCreatePost = "Can not create the post"
	errUpdatePost = "Can not update the post"
	errDeletePost = "Can not delete the post"
	errScanPost   = "Can not Scan the post"
	errUnMarshal  = "Can not unmarshal the row"

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

//GetPosts get a list of posts from database by limit index
func GetPosts(ctx context.Context, limitPosts, page int) ([]*proto.Post, error) {
	if limitPosts < 10 {
		limitPosts = 10
	}
	if page < 1 {
		page = 1
	}
	limit := limitPosts
	offset := limit*page - limit
	query := "SELECT * FROM posts LIMIT $1 OFFSET $2"

	rows, err := databaseService.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, sql.ErrNoRows
	}
	defer rows.Close()

	var posts []*proto.Post
	if errR := rows.Next(); errR {
		var comments []uint8
		var post proto.Post
		if err := rows.Scan(&post.Id, &post.Author, &post.Title,
			&post.Content, &post.Timestemp, &post.Published,
			&comments); err != nil {
			fmt.Println(err)
			return nil, errors.New(errScanPost)
		}
		if errM := json.Unmarshal(comments, &post.Comments); errM != nil {
			return nil, errors.New(errUnMarshal)
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

//GetSinglePost get a post from database
func GetSinglePost(ctx context.Context, postID int) (*proto.Post, error) {
	query := "SELECT * FROM posts WHERE id=$1"
	post := &proto.Post{}
	var comments []uint8
	if err := databaseService.QueryRowContext(ctx, query, postID).
		Scan(&post.Id, &post.Author, &post.Title, &post.Content,
			&post.Timestemp, &post.Published, &comments); err != nil {

		fmt.Println(err)
		return nil, sql.ErrNoRows
	}
	json.Unmarshal(comments, &post.Comments)
	return post, nil
}

//CreatePostDB create a post in database
func CreatePostDB(ctx context.Context, post *proto.Post) error {
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
		return errors.New(errBeginTX)
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

// DeletePost delete apost in database
func DeletePost(ctx context.Context, postID int) error {
	tx, errTX := databaseService.Begin()

	if errTX != nil {
		return errors.New(errBeginTX)
	}

	queryDelete := "DELETE FROM posts WHERE id=$1"
	_, errD := tx.ExecContext(ctx, queryDelete, postID)

	if errD != nil {
		print("2")
		print(errD)
		return errors.New(errDeletePost)
	}
	if errTX = tx.Commit(); errTX != nil {
		return errors.New(errCommitTX)
	}
	return nil
}
