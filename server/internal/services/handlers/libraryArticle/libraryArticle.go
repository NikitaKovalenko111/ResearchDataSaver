package libraryArticleService

import (
	"research-data-saver/internal/models"
	libraryArticleRepo "research-data-saver/internal/storage/repositories/libraryArticle"
)

type LibraryArticleService struct {
	repo *libraryArticleRepo.LibraryArticleRepo
}

func Init(repo *libraryArticleRepo.LibraryArticleRepo) *LibraryArticleService {
	return &LibraryArticleService{
		repo: repo,
	}
}

func (s *LibraryArticleService) GetAll(queryName string) (*[]models.LibraryArticle, error) {
	var articles *[]models.LibraryArticle

	articles, err := s.repo.GetArticles(queryName)

	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *LibraryArticleService) AddArticle(name string, annotation string, link string, publishingDate string, lang string, udk string, publisherObject string, publisher string, supervisor string, authors []string) (*models.LibraryArticle, error) {
	var addedArticle *models.LibraryArticle

	addedArticle, err := s.repo.AddArticle(name, annotation, link, publishingDate, lang, udk, publisherObject, publisher, supervisor, authors)

	if err != nil {
		return nil, err
	}

	return addedArticle, nil
}
