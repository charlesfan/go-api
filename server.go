package main

import (
	"flag"

	"github.com/jinzhu/gorm"

	"github.com/charlesfan/go-api/repository"
	"github.com/charlesfan/go-api/repository/sqlite"
	"github.com/charlesfan/go-api/route"
	"github.com/charlesfan/go-api/service/rsi"
)

var (
	d string
	h string
	m bool
)

func main() {
	flag.StringVar(&d, "db", "sqlite", "db type")
	flag.StringVar(&h, "host", "", "the host of db")
	flag.BoolVar(&m, "migrate", false, "do db migrate")

	flag.Parse()

	// DB init
	database := newDatabase(d, h, m)
	defer database.Close()

	// rsi.Services init
	rsi.Init(database)
	router := route.Init()

	router.Run(":8080")
}

func newDatabase(s, host string, m bool) *repository.Database {

	var db *gorm.DB
	switch s {
	case "sqlite":
		db, _ = sqlite.NewSqlite(host, m)
	}
	return repository.NewDatabase(db)
}
