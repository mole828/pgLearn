package main

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

func main() {
	var db = pg.Connect(&pg.Options{
		Addr:     "localhost:15432",
		User:     "app",
		Password: "app",
		Database: "app",
	})
	var ping_error = db.Ping(context.Background())
	if ping_error != nil {
		logrus.Error(ping_error)
	}
	logrus.Info("done")
}
