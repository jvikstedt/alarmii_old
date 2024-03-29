package models

import (
	"log"
	"os"

	"github.com/boltdb/bolt"
)

// Database bolt database
var Database *bolt.DB

// OpenDatabase opens database connection to provided file
func OpenDatabase(dbPath string) {
	db, err := bolt.Open(dbPath, 0600, nil)
	Database = db
	if err != nil {
		log.Fatal(err)
	}
}

// CloseDatabase closes database
func CloseDatabase() {
	Database.Close()
}

// ClearDatabase removes db file and recreates it
func ClearDatabase(dbPath string) {
	CloseDatabase()
	os.Remove(dbPath)
	OpenDatabase(dbPath)
}
