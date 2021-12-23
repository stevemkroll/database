package database

import "github.com/stevemkroll/database/internal/client"

func NewClient() *client.Client {
	c := new(client.Client)
	c.Host = "localhost"
	c.Port = "5432"
	c.User = "user"
	c.Password = "password"
	c.DBName = "db"
	c.SSLMode = "disable"
	return c
}
