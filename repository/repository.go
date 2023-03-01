package repository

import (
	"database/sql"
	"github.com/database"
	"log"
)

func GetConnection() *sql.DB {
	db, error := database.Connect()
	if error != nil {
		log.Fatal("Error connecting to database", error)
	}

	return db
}
