package initializer

import (
	"fmt"
	"os"
	"url_shortener/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbinitializer struct {
	DB *gorm.DB
	L  *logger.Logger
}

func ConnectDb(dbi *Dbinitializer) {
	dbi.L.LogMessage("Entering into connecting db")
	var err error
	dbstring := os.Getenv("DATABASE_URL")
	fmt.Println("dbstring is",dbstring)
	if dbstring == "" {
		dbi.L.LogFatalMessage("Database details are not loaded from the env file")
	}
	dbi.DB, err = gorm.Open(postgres.Open(dbstring), &gorm.Config{})
	if err != nil {
		dbi.L.LogFatalMessage("Failed to connect to db host")
	}
}
