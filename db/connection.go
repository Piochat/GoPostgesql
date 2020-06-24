package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Driver para interactuar
)

const (
	dsn string = "postgres://piocha:koke@127.0.0.1:5432/gocrud?sslmode=disable"
)

//DbCon Connection variable
var DbCon = connection()

//connection Get a connection to the database
func connection() *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln("Error [file connection]", err.Error())
	}

	if db.Ping() != nil {
		log.Println("DB Die", db.Ping().Error())
	}

	return db
}

//PingToDB check the connection to DB
func PingToDB() bool {
	return DbCon.Ping() != nil
}
