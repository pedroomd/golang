package config
 
import (
    "database/sql"
	"log"
)

func Connect() *sql.DB {
    dbDriver := "mysql"
    dbUser := "user"
    dbPass := "password"
    dbName := "getground"
 
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
		log.Fatal(err)
	}
    return db
}