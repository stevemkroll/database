package test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

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

func TestScan(t *testing.T) {
	TestQuery(t)
	type Account struct {
		CreatedAt time.Time
		UpdatedAt time.Time
		Email     string
		Password  string
		ID        int
	}
	var Accounts []Account
	for testRows.Next() {
		a := new(Account)
		if err := testRows.Scan(
			&a.ID,
			&a.Email,
			&a.Password,
			&a.CreatedAt,
			&a.UpdatedAt,
		); err != nil {
			t.Fatal(err)
		}
		Accounts = append(Accounts, *a)
	}
	if len(Accounts) == 0 {
		t.Fatal("no accounts found")
	}
}
