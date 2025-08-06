package main

import (
	"log"
	"os"

	"github.com/lpernett/godotenv"
	"github.com/phamaiden/blog-platform-api/internal/handlers"
	"github.com/phamaiden/blog-platform-api/internal/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading env: %s", err)
	}

	cfg := &config{
		addr: os.Getenv("ADDR"),
		//dbUrl: os.Getenv("DB_URL"),
	}

	app := &application{
		config: *cfg,
	}

	services := services.

	handler := handlers.NewBlogHandler()

	r := app.mount()

	log.Fatal(app.run(r))
}
