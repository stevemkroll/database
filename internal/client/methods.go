package client

import (
	"fmt"
	"strings"
)

func (c *Client) buildConnStr() string {
	var connStr strings.Builder
	connStr.WriteString(fmt.Sprintf("host=%s\n", c.Host))
	connStr.WriteString(fmt.Sprintf("port=%s\n", c.Port))
	connStr.WriteString(fmt.Sprintf("user=%s\n", c.User))
	connStr.WriteString(fmt.Sprintf("password=%s\n", c.Password))
	connStr.WriteString(fmt.Sprintf("dbname=%s\n", c.DBName))
	connStr.WriteString(fmt.Sprintf("sslmode=%s\n", c.SSLMode))
	return connStr.String()
}
