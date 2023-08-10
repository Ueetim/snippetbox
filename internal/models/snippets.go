package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID 		int
	Title	string
	Content	string
	Created	time.Time
	Expires	time.Time
}

// wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// execute db command. Exec is used for commands that dont return from db, like INSERT and DELETE
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// get id of newly inserted record
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// convert id from int64 to int
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}