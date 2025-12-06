package storage

import (
	"database/sql"
	"fmt"
	"research-data-saver/internal/config"
	documentRepository "research-data-saver/internal/storage/repositories/document"
	fipsRepository "research-data-saver/internal/storage/repositories/fips"
	internetArticleRepo "research-data-saver/internal/storage/repositories/internetArticle"
	libraryArticleRepo "research-data-saver/internal/storage/repositories/libraryArticle"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db           *sql.DB
	Repositories *Repositories
}

type Repositories struct {
	InternetArticleRepo *internetArticleRepo.InternetArticleRepo
	LibraryArticleRepo  *libraryArticleRepo.LibraryArticleRepo
	FipsRepo            *fipsRepository.FipsRepository
	DocumentRepo        *documentRepository.DocumentRepository
}

func Connect(cfg *config.Config) *sql.DB {
	connString := fmt.Sprintf("host=%s user=%s port=%d password=%s dbname=%s sslmode=disable", cfg.Storage.Host, cfg.Username, cfg.Storage.Port, cfg.Storage.Password, cfg.Storage.Name)

	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}

	return db
}

func Init(db *sql.DB) *Storage {
	internetArticleRepo := internetArticleRepo.Init(db)
	libraryArticleRepo := libraryArticleRepo.Init(db)
	fipsRepo := fipsRepository.Init(db)
	documentRepo := documentRepository.Init(db)

	repos := Repositories{
		InternetArticleRepo: internetArticleRepo,
		LibraryArticleRepo:  libraryArticleRepo,
		FipsRepo:            fipsRepo,
		DocumentRepo:        documentRepo,
	}

	storage := Storage{
		Db:           db,
		Repositories: &repos,
	}

	return &storage
}
