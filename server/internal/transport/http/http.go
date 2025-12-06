package http

import (
	"research-data-saver/internal/services"
	documentController "research-data-saver/internal/transport/http/controllers/document"
	fipsController "research-data-saver/internal/transport/http/controllers/fips"
	internetArticleController "research-data-saver/internal/transport/http/controllers/internetArticle"
	libraryArticleController "research-data-saver/internal/transport/http/controllers/libraryArticle"

	"github.com/gofiber/fiber/v2"
)

type HTTP struct {
	LibraryArticleController  *libraryArticleController.LibraryArticleController
	InternetArticleController *internetArticleController.InternetArticleController
	FipsController            *fipsController.FipsController
	DocumentController        *documentController.DocumentController
}

func Init(router *fiber.App, services *services.Services) *HTTP {
	return &HTTP{
		LibraryArticleController:  libraryArticleController.Init(router, services.LibraryArticleService),
		InternetArticleController: internetArticleController.Init(router, services.InternetArticleService),
		DocumentController:        documentController.Init(router, services.DocumentService),
		FipsController:            fipsController.Init(router, services.FipsService),
	}
}
