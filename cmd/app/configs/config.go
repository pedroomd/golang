package config

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
    db, err := sql.Open("mysql", "user:password@tcp(getgroundDB)/database")
    if err != nil {
        
		log.Fatal(err)
	}
    return db
}