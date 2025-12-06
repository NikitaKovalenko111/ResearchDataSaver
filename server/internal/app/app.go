package app

import (
	"log"
	"research-data-saver/internal/config"
	"research-data-saver/internal/logger"
	"research-data-saver/internal/services"
	"research-data-saver/internal/storage"
	"research-data-saver/internal/transport/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Run() {
	cfg := config.Init()

	logger := logger.Init(cfg)

	logger.Info("Logger is enabled")
	logger.Debug("Debug logger is enabled")

	db := storage.Connect(cfg)

	logger.Info("Successfully connected to database")

	storage := storage.Init(db)

	logger.Info("Storage initialized")

	services := services.Init(storage)

	logger.Info("Services initialized")

	app := fiber.New()

	app.Use(cors.New())

	http := http.Init(app, services)

	http.LibraryArticleController.Start("/library-article")
	http.InternetArticleController.Start("/internet-article")
	http.FipsController.Start("/fips")
	http.DocumentController.Start("/document")

	log.Fatal(app.Listen(cfg.Address))

	logger.Info("HTTP server started on " + cfg.Address)
}
