package db

import (
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewMySQLiteStorage(databaseUrl string, authToken string) (*gorm.DB, error) {
	dsn := databaseUrl + "?authToken=" + authToken

	db, err := gorm.Open(sqlite.New(sqlite.Config{
		DriverName: "libsql",
		DSN:        dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
