package database

import (
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() error {
	var err error
	db, err = gorm.Open(sqlite.Open("db/store.db"), &gorm.Config{})
	return err
}
