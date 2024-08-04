package db

import (
	"database/sql"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func NewMySQLiteStorage(databaseUrl string, authToken string) (*sql.DB, error) {
	url := "libsql://" + databaseUrl + "?authToken=" + authToken

	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
