package client

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type ClientProvider interface {
	Open() (*sql.DB, error)
}

type Client struct {
	ClientProvider
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (c *Client) Open() (*sql.DB, error) {
	connStr := c.buildConnStr()
	return sql.Open("postgres", connStr)
}
