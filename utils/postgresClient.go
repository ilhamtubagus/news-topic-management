package utils

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseClient() *gorm.DB {
	DSN := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	dbClient, err := gorm.Open(postgres.New(postgres.Config{DSN: DSN}))
	if err != nil {
		log.Fatal(err)
	}
	return dbClient
}
