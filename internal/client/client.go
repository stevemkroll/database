package client

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ClientProvider interface {
	Connect() (*sqlx.DB, error)
	Get(dest interface{}, query string) error
	Select(dest interface{}, query string) error
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

func (c *Client) Connect() (*sqlx.DB, error) {
	connStr := c.buildConnStr()
	return sqlx.Connect("postgres", connStr)
}

func (c *Client) Get(dest interface{}, query string) error {
	db, err := c.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Get(dest, query)
}

func (c *Client) Select(dest interface{}, query string) error {
	db, err := c.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Select(dest, query)
}
