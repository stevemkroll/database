package test

import (
	"fmt"
	"testing"

	"github.com/stevemkroll/database/internal/client"
	"github.com/stevemkroll/database/pkg/database"
)

var (
	testClient *client.Client
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

func TestClientOpen(t *testing.T) {
	TestNewClient(t)
	db, err := testClient.Open()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", db)
}
