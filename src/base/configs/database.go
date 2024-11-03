package configs

import (
	"dainxor/we/base/logger"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type db struct{}

var DB db
var DataBase *gorm.DB

func (db) Get() *gorm.DB {
	return DataBase
}

func (db) EnvInit() {
	dbType, exist := os.LookupEnv("DB_TYPE")

	if exist && dbType == "POSTGRES" {
		logger.Info("Connecting to Postgres database")
		DB.ConnectPostgresEnv()
	} else {
		logger.Info("Connecting to SQLite database")
		DB.ConnectSQLite()
	}
}

func (db) ConnectPostgresEnv() {
	useTesting, exist := os.LookupEnv("DB_TESTING")
	if exist && useTesting != "TRUE" {
		DB.ConnectPostgres(
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
	} else {
		logger.Info("Using testing database")
		DB.ConnectPostgres(
			os.Getenv("DB_HOST_TEST"),
			os.Getenv("DB_USER_TEST"),
			os.Getenv("DB_PASSWORD_TEST"),
			os.Getenv("DB_NAME_TEST"),
			os.Getenv("DB_PORT_TEST"),
		)
	}
}
func (db) ConnectPostgres(host string, user string, password string, dbname string, port string) {
	var err error
	dsn := "host=" + host +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" port=" + port +
		" sslmode=disable"
	logger.Info("Connecting to database: ", dsn)
	DataBase, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Fatal(err)
	}
}

func (db) ConnectSQLite() {
	var err error
	DataBase, err = gorm.Open(sqlite.Open("worldengsqlite.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
