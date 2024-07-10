package conn

import "github.com/go-pg/pg/v10"

func GetDB() *pg.DB {
	var db = pg.Connect(&pg.Options{
		Addr:     "localhost:15432",
		User:     "app",
		Password: "app",
		Database: "app",
	})
	return db
}
