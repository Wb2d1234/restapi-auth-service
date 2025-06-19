package main

import (
	"auth-service/internal/config"
	"auth-service/internal/handlers"
	"auth-service/internal/middleware"
	"auth-service/internal/models"
	"auth-service/internal/storage"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or failed to load")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Config error:", err)
	}

	pg, err := storage.NewPostgresStorage(cfg)
	if err != nil {
		log.Fatal("Database error:", err)
	}
	defer pg.DB.Close()

	userRepo := models.NewUserRepository(pg.DB)
	log.Printf("Loaded config: DB_DSN=%s, JWT_SECRET=%s, PORT=%s", cfg.DB_DSN, cfg.JWT_SECRET, cfg.PORT)

	// Маршруты
	http.Handle("/register", handlers.RegisterHandler(userRepo))
	http.Handle("/login", handlers.LoginHandler(userRepo, cfg.JWT_SECRET))

	// Защищённый маршрут
	http.Handle("/profile", middleware.AuthMiddleware(cfg.JWT_SECRET)(handlers.ProfileHandler()))

	log.Printf("Server starting on port %s", cfg.PORT)
	log.Fatal(http.ListenAndServe(":"+cfg.PORT, nil))

}
