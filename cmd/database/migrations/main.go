package main

import (
	"fmt"

	"github.com/bagusyanuar/go-olin-bags-with-fiber/cmd/database/scheme"
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

	scheme.Migrate(db)
	fmt.Println("successfully migrating database")
}
