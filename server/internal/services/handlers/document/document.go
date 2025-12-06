package documentService

import (
	"research-data-saver/internal/models"
	documentRepo "research-data-saver/internal/storage/repositories/document"
)

type DocumentService struct {
	repo *documentRepo.DocumentRepository
}

func Init(repo *documentRepo.DocumentRepository) *DocumentService {
	return &DocumentService{
		repo: repo,
	}
}

func (s *DocumentService) AddDocument(name string, annotation string, link string, publishingDate string, author string) (*models.Document, error) {
	var document *models.Document

	document, err := s.repo.Create(name, annotation, link, publishingDate, author)

	if err != nil {
		return nil, err
	}

	return document, nil
}

func (s *DocumentService) GetAll() (*[]models.Document, error) {
	var documents *[]models.Document

	documents, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return documents, nil
}
