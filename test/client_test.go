package test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stevemkroll/database/internal/client"
	"github.com/stevemkroll/database/pkg/database"
)

var (
	err          error
	testClient   *client.Client
	testDatabase *sqlx.DB
)

type testAccount struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	ID        int64     `db:"id"`
}

func TestNewClient(t *testing.T) {
	testClient = database.NewClient()
	switch fmt.Sprintf("%T", testClient) {
	case "*client.Client":
		break
	default:
		t.Fatal("invalid type")
	}
}

func TestConnect(t *testing.T) {
	TestNewClient(t)
	testDatabase, err = testClient.Connect()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	TestConnect(t)
	q := "SELECT * FROM accounts WHERE email='test@test.com'"
	var acnt testAccount
	if err := testClient.Get(&acnt, q); err != nil {
		t.Fatal(err)
	}
	if acnt.Email != "test@test.com" {
		t.Fatal(sql.ErrNoRows)
	}
}

func TestSelect(t *testing.T) {
	TestConnect(t)
	q := "SELECT * FROM accounts"
	var acnts []testAccount
	if err := testClient.Select(&acnts, q); err != nil {
		t.Fatal(err)
	}
	if len(acnts) < 1 {
		t.Fatal(sql.ErrNoRows)
	}
}
