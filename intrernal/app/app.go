package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	// run api
	go func() {
		if err := app.api.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error on api server: %s\n", err.Error())
		}
	}()

	// run message handler

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := app.api.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown error: %s\n", err.Error())
	}
	select {
	case <-ctx.Done():
		log.Println("server shutdown timed out")
	}
	log.Println("server closing")
}
