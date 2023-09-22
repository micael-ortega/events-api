package internals

import (
	"database/sql"
	"log"
)

func OpenDb() *sql.DB {
	db, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return db
}