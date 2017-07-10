package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/go-errors/errors"
	"fmt"
)
/*
 sql orm thirdParty library
 sqlx

 */

type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}

	err = Db.QueryRow("INSERT INTO comments (content,author,post_id) values ($1,$2,$3) returning id",
		comment.Content, comment.Author, comment.Post.Id).
		Scan(&comment.Id)
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = Db.QueryRow("SELECT id,content,author FROM posts WHERE id = $1", id).
		Scan(&post.Id, &post.Content, &post.Author)
	rows, err := Db.Query("SELECT id,content,author from comments")
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}

func (post *Post) Create() (err error) {
	Db.QueryRow("INSERT INTO posts (content,author) values ($1,$2) returning id", post.Content, post.Author).
		Scan(&post.Id)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	post.Create()

	comment := Comment{Content: "Good post!", Author: "Joe", Post: &post}
	comment.Create()
	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)
}