package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/XineAurora/fio-statistics/intrernal/api"
	"github.com/XineAurora/fio-statistics/intrernal/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// open db
// run kafka consumer
// run api

type App struct {
	repo database.FIORepository
	api  *http.Server
}

func New() *App {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_SSLMODE"),
		os.Getenv("POSTGRES_TIMEZONE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	repo := database.NewDBFIORepository(db)
	return &App{
		repo: repo,
		api:  api.NewApiServer(repo),
		// kafka worker
	}
}

func (app *App) Run() {

}
