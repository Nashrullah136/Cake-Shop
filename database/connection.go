package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"sync"
)

var lock = &sync.Mutex{}

var db *sql.DB

func DefaultConfig() mysql.Config {
	return mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}
}

func SetUpDbConnection(config mysql.Config) *sql.DB {
	log.Printf("Connecting to database with cfg %v", config)
	dbConnection, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	if pingErr := dbConnection.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}
	return dbConnection
}

func GetDbConnection() *sql.DB {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			db = SetUpDbConnection(DefaultConfig())
		} else {
			if pingErr := db.Ping(); pingErr != nil {
				db = SetUpDbConnection(DefaultConfig())
			}
		}
	}
	return db
}
