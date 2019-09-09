package templates

// DatabaseGO ...
func DatabaseGO() (string, string) {
	filename := "database/database.go"
	body := `package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB ...
var DB *gorm.DB

// Init - func for connection to db
func Init(dbUser string, dbPassword string, dbName string, dbHost string, dbPort string) (*gorm.DB, error) {
	var connectionString = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		dbHost, dbPort, dbUser, dbName, dbPassword,
	)
	db, err := gorm.Open("postgres", connectionString)
	DB = db
	return DB, err
}

// GetDB - return db instance
func GetDB() *gorm.DB {
	return DB
}

`
	return filename, body
}
