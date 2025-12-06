package services

import (
	documentService "research-data-saver/internal/services/handlers/document"
	fipsService "research-data-saver/internal/services/handlers/fips"
	internetArticleService "research-data-saver/internal/services/handlers/internetArticle"
	libraryArticleService "research-data-saver/internal/services/handlers/libraryArticle"
	"research-data-saver/internal/storage"
)

type Services struct {
	LibraryArticleService  *libraryArticleService.LibraryArticleService
	DocumentService        *documentService.DocumentService
	InternetArticleService *internetArticleService.InternetArticleService
	FipsService            *fipsService.FipsService
}

func Init(storage *storage.Storage) *Services {
	return &Services{
		LibraryArticleService:  libraryArticleService.Init(storage.Repositories.LibraryArticleRepo),
		DocumentService:        documentService.Init(storage.Repositories.DocumentRepo),
		InternetArticleService: internetArticleService.Init(storage.Repositories.InternetArticleRepo),
		FipsService:            fipsService.Init(storage.Repositories.FipsRepo),
	}
}
