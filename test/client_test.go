package test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stevemkroll/database/internal/client"
	"github.com/stevemkroll/database/pkg/database"
)

var (
	err          error
	testClient   *client.Client
	testDatabase *sql.DB
	testRows     *sql.Rows
)

func TestNewClient(t *testing.T) {
	testClient = database.NewClient()
	switch fmt.Sprintf("%T", testClient) {
	case "*client.Client":
		break
	default:
		t.Fatal("invalid type")
	}
}

func TestOpen(t *testing.T) {
	TestNewClient(t)
	testDatabase, err = testClient.Open()
	if err != nil {
		t.Fatal(err)
	}
}

func TestQuery(t *testing.T) {
	TestNewClient(t)
	query := "SELECT * FROM accounts"
	testRows, err = testClient.Query(query)
	if err != nil {
		t.Fatal(err)
	}
}
