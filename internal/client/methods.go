package client

import (
	"fmt"
	"strings"
)

func (c *Client) buildConnStr() string {
	var connStr strings.Builder
	connStr.WriteString(fmt.Sprintf("host=%s", c.Host))
	connStr.WriteString(fmt.Sprintf("port=%s", c.Port))
	connStr.WriteString(fmt.Sprintf("user=%s", c.User))
	connStr.WriteString(fmt.Sprintf("password=%s", c.Password))
	connStr.WriteString(fmt.Sprintf("dbname=%s", c.DBName))
	connStr.WriteString(fmt.Sprintf("sslmode=%s", c.SSLMode))
	return connStr.String()
}
