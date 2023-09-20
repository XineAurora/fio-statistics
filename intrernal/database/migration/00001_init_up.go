package main

import (
	"fmt"
	"log"
	"os"

	"github.com/XineAurora/fio-statistics/intrernal/database/models"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	POSTGRES_HOST     = "POSTGRES_HOST"
	POSTGRES_USER     = "POSTGRES_USER"
	POSTGRES_PASSWORD = "POSTGRES_PASSWORD"
	POSTGRES_DBNAME   = "POSTGRES_DBNAME"
	POSTGRES_PORT     = "POSTGRES_PORT"
	POSTGRES_SSLMODE  = "POSTGRES_SSLMODE"
	POSTGRES_TIMEZONE = "POSTGRES_TIMEZONE"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv(POSTGRES_HOST), os.Getenv(POSTGRES_USER), os.Getenv(POSTGRES_PASSWORD),
		os.Getenv(POSTGRES_DBNAME), os.Getenv(POSTGRES_PORT),
		os.Getenv(POSTGRES_SSLMODE), os.Getenv(POSTGRES_TIMEZONE))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error on opening db connection: %s", err.Error())
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		log.Println("transaction began")
		m := tx.Migrator()
		if m.HasTable(models.FIO{}) {
			log.Println("table FIO exists, dropping it")
			if err := m.DropTable(models.FIO{}); err != nil {
				return err
			}
		}
		log.Println("creating table FIO")
		if err := m.CreateTable(models.FIO{}); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Printf("error occured  during transaction: %s\n", err.Error())
	} else {
		log.Println("transaction commited successfully")
	}

}
