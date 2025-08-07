package main

import (
	"context"
	"log"
	"os"

	"github.com/lpernett/godotenv"
	"github.com/phamaiden/blog-platform-api/internal/db"
	"github.com/phamaiden/blog-platform-api/internal/handlers"
	"github.com/phamaiden/blog-platform-api/internal/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading env: %s", err)
	}

	cfg := config{
		addr:  os.Getenv("ADDR"),
		dbUrl: os.Getenv("DB_URL"),
	}

	app := &application{
		config: cfg,
	}

	conn, err := db.Init(cfg.dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	queries := db.New(conn)

	services := services.NewBlogService(queries)

	handler := handlers.NewBlogHandler(services)

	r := app.mount(handler)

	log.Fatal(app.run(r))
}
