package fipsService

import (
	"research-data-saver/internal/models"
	fipsRepo "research-data-saver/internal/storage/repositories/fips"
)

type FipsService struct {
	repo *fipsRepo.FipsRepository
}

func Init(repo *fipsRepo.FipsRepository) *FipsService {
	return &FipsService{
		repo: repo,
	}
}

func (s *FipsService) GetAll(queryName string, queryContentType string, queryContentReg string, queryContentDate string) (*[]models.FipsContent, error) {
	var fips *[]models.FipsContent

	fips, err := s.repo.GetAll(queryName, queryContentType, queryContentReg, queryContentDate)

	if err != nil {
		return nil, err
	}

	return fips, nil
}

func (s *FipsService) Create(name string, link string, fipsType string, annotation string, registration string, publishingDate string, applicant string, address string, authors []string) (*models.FipsContent, error) {
	var createdFips *models.FipsContent

	createdFips, err := s.repo.Create(name, link, fipsType, annotation, registration, publishingDate, applicant, address, authors)

	if err != nil {
		return nil, err
	}

	return createdFips, nil
}
