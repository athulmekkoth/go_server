package main

import (
	"database/sql"
	"log"

	"github.com/athulmekkoth/go_server.git/cmd/api"
	"github.com/athulmekkoth/go_server.git/db"

	"github.com/athulmekkoth/go_server.git/config"
	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.NewConfig()
	cfg.User = config.Envs.DBUser
	cfg.Passwd = config.Envs.DBPassword
	cfg.Addr = config.Envs.DBAddress
	cfg.DBName = config.Envs.DBName
	cfg.Net = "tcp"
	cfg.AllowNativePasswords = true
	cfg.ParseTime = true

	// Pass the configured mysql.Config to NewMySQLStorage
	db, err := db.NewMyQLStorage(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)
	server := api.NewAPIServer(":8000", nil)
	if err := server.RUN(); err != nil {
		log.Fatal()
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db connected")
}
