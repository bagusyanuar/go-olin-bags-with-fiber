package main

import (
	"github.com/bagusyanuar/go-olin-bags-with-fiber/app/server"
	"github.com/bagusyanuar/go-olin-bags-with-fiber/config"
)

func main() {
	cfg, err := config.NewConfig(".env")
	if err != nil {
		panic(err)
	}

	db, err := config.NewMySQLConnection(&cfg.MySQL)

	if err != nil {
		panic(err)
	}

	if cfg.AppMode == "dev" {
		db = db.Debug()
	}

	server.Listen(cfg, db)

}
