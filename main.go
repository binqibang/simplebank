package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	api "simplebank/api/v1"
	db "simplebank/db/sqlc"
	"simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	err = server.Start(config.Address)
	if err != nil {
		log.Fatal("can not start http server:", err)
	}
}
