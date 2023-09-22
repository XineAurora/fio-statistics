package main

import (
	"github.com/XineAurora/fio-statistics/intrernal/app"
	_ "github.com/joho/godotenv/autoload"
)

// consume message
// verify message
// enrich message
// put in db

func main() {
	a := app.New()
	a.Run()
}
