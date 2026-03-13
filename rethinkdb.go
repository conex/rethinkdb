package rethinkdb

import (
	"testing"
	"time"

	gorethink "gopkg.in/gorethink/gorethink.v3"

	"github.com/omeid/conex"
)

var (
	// Image to use for the box.
	Image = "rethinkdb:2"
	// Port used for connecting to RethinkDB.
	Port = "28015"

	// RethinkUpWaitTime dictates how long we should wait for RethinkDB to accept connections.
	RethinkUpWaitTime = 10 * time.Second
)

func init() {
	conex.Require(func() string { return Image })
}

// Box returns a RethinkDB client connect to a RethinkDB
// container based on your provided tags.
func Box(t testing.TB, db string) (*gorethink.Session, conex.Container) {
	c := conex.Box(t, &conex.Config{
		Image:  Image,
		Expose: []string{Port},
	})

	t.Log("Waiting for RethinkDB to accept connections")

	err := c.Wait(Port, RethinkUpWaitTime)
	if err != nil {
		c.Drop()
		t.Fatal("RethinkDB failed to start:", err)
	}

	t.Log("RethinkDB is now accepting connections")

	opts := gorethink.ConnectOpts{
		Address:  c.Address() + ":" + Port,
		Database: db,
	}

	sesh, err := gorethink.Connect(opts)
	if err != nil {
		c.Drop()
		t.Fatal(err)
	}

	return sesh, c
}
