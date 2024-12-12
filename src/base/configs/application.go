package configs

import (
	"dainxor/we/base/logger"

	"github.com/joho/godotenv"
)

type app struct{}
type fields struct {
	Env     string
	Address string

	DBType     string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	ProxyURL string
	FrontURL string
}

var App app

func (app) LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		logger.Error("Error loading .env file")

	}
}

func (app) createDefaults() {

}

func (app) setConfigEnvDefault() {

}

func (app) Enviroment() string {
	return "development"
}
