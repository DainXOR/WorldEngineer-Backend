package configs

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgresTest() {
	var err error
	dsn := "host=localhost user=postgres password=2407 dbname=worldeng-test port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}

func ConnectPostgres() {
	var err error
	dsn := "host=localhost user=postgres password=2407 dbname=worldeng port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}

func ConnectSQLite() {
	var err error
	DB, err = gorm.Open(sqlite.Open("worldengsqlite.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
