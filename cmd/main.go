package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/05blue04/Flow/internal/config"
	"github.com/05blue04/Flow/internal/database"
	"github.com/05blue04/Flow/internal/handlers"
	"github.com/05blue04/Flow/internal/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration", "error", err)
	}

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to open database", "error", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	log.Println("Database connection established")

	app := &config.App{
		Cfg: cfg,
		DB:  database.New(db),
	}

	userHandler := handlers.NewUserHandler(app)
	healthHandler := handlers.NewHealthHandler(app)

	r := router.Setup(&router.Handlers{
		User:   userHandler,
		Health: healthHandler,
	})

	server := &http.Server{
		Handler: r,
		Addr:    ":" + cfg.Port,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
