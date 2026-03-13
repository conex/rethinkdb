package rethinkdb_test

import (
	"os"
	"testing"

	"github.com/conex/rethinkdb"
	"github.com/omeid/conex"
	gorethink "gopkg.in/gorethink/gorethink.v3"
)

func TestMain(m *testing.M) {
	os.Exit(conex.Run(m))
}

func TestRethinkDB(t *testing.T) {
	sesh, con := rethinkdb.Box(t, "test")
	defer con.Drop()

	// Test basic connectivity with a simple server info query
	cursor, err := gorethink.Expr(1).Run(sesh)
	if err != nil {
		t.Fatal(err)
	}
	defer cursor.Close()

	var result int
	if err := cursor.One(&result); err != nil {
		t.Fatal(err)
	}

	if result != 1 {
		t.Fatalf("expected 1, got %d", result)
	}
}
