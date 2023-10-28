package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func dbConnection() {
	stringConnection := getStringConnection()
	DB, err = gorm.Open(postgres.Open(stringConnection), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Print("... DB Connected ...")
}

func HandleDatabaseConnection() {
	dbConnection()
	dbAutomigrate()
}
