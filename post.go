package main

import (
	"database/sql"
	"time"
)

type DBManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) *DBManager {
	return &DBManager{
		db: db,
	}
}

type Posts struct {
	Id        int
	Post      string
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (c *DBManager) CreatePost(post *Posts) (*Posts, error) {
	query := `
		insert into posts(
		 	Post,
		    UserId
		) values ($1,$2)
			returning Post,UserId
	`

	row := c.db.QueryRow(
		query,
		post.Post,
		post.UserId,
	)

	var result Posts

	err := row.Scan(
		&result.Post,
		&result.UserId,
	)

	if err != nil {
		return nil, err
	}
	return &result, nil

}
