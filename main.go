package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	PostgresUser     = "postgres"
	PostgresPassword = "7"
	PostgresDatabase = "golang"
	PostgresHost     = "localhost"
	PostgresPort     = 5432
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
		PostgresHost,
		PostgresPort,
		PostgresUser,
		PostgresPassword,
		PostgresDatabase,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	dbManager := NewDBManager(db)

	posts, err := dbManager.CreatePost(&Posts{
		Post:   "Motivation",
		UserId: 1,
	})

	if err != nil {
		log.Fatalf("failed to create post: %v", err)
	}

	PrintPost(posts)
}

func PrintPost(post *Posts) {
	fmt.Println("-----------Post----------")
	fmt.Println("Id: ", post.Id)
	fmt.Println("Post: ", post.Post)
	fmt.Println("User Id: ", post.UserId)
	fmt.Println("Created at: ", post.CreatedAt)
	fmt.Println("Updated at: ", post.UpdatedAt)
	fmt.Println("Deleted at: ", post.DeletedAt)
}
