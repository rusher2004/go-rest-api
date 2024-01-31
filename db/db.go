package db

import (
	"database/sql"
	"fmt"
)

type Client struct {
	db *sql.DB
}

func NewDB(driver, dsn string) (Client, error) {
	pool, err := sql.Open(driver, dsn)
	if err != nil {
		return Client{}, fmt.Errorf("error opening db connection: %w", err)
	}

	return Client{pool}, nil
}

func (c *Client) Close() error {
	return c.db.Close()
}
