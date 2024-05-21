package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/atanda0x/e-commerce/util"
	_ "github.com/lib/pq"
)

func DBset() *sql.DB {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to postresql!!!!!!!!!!!!!")
	return db
}
