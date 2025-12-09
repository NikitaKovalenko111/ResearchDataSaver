package internetArticleService

import (
	"research-data-saver/internal/models"
	internetArticleRepo "research-data-saver/internal/storage/repositories/internetArticle"
)

type InternetArticleService struct {
	repo *internetArticleRepo.InternetArticleRepo
}

func Init(repo *internetArticleRepo.InternetArticleRepo) *InternetArticleService {
	return &InternetArticleService{
		repo: repo,
	}
}

func (s *InternetArticleService) GetAll(queryName string) (*[]models.InternetArticle, error) {
	var articles *[]models.InternetArticle

	articles, err := s.repo.GetArticles(queryName)

	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *InternetArticleService) AddArticle(name string, annotation string, link string, publishingDate string, author string, searchingMachine string) (*models.InternetArticle, error) {
	var addedArticle *models.InternetArticle

	addedArticle, err := s.repo.AddArticle(name, annotation, link, publishingDate, author, searchingMachine)

	if err != nil {
		return nil, err
	}

	return addedArticle, nil
}
