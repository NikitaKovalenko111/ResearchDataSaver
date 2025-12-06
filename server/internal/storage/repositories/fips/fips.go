package fipsRepository

import (
	"database/sql"
	"fmt"
	"research-data-saver/internal/models"
)

type FipsRepository struct {
	db *sql.DB
}

func Init(db *sql.DB) *FipsRepository {
	return &FipsRepository{db: db}
}

func authorValues(fipsId int, authors []string) string {
	values := ""

	for i, author := range authors {
		values += fmt.Sprintf("(%d, '%s')", fipsId, author)

		if i < len(authors)-1 {
			values += ", "
		}
	}

	return values
}

func (r *FipsRepository) Create(name string, link string, fipsType string, annotation string, registration string, publishingDate string, applicant string, address string, authors []string) (*models.FipsContent, error) {
	var createdFips models.FipsContent
	query := `
	INSERT INTO fips_content (content_name, link, content_type, content_annotation, content_registration, content_publishing_date, content_applicant, content_address)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, content_name, link, content_type, content_annotation, content_registration, content_publishing_date, content_applicant, content_address`

	tx, err := r.db.Begin()

	if err != nil {
		return nil, err
	}

	err = tx.QueryRow(query, name, link, fipsType, annotation, registration, publishingDate, applicant, address).Scan(&createdFips.Id, &createdFips.Name, &createdFips.Link, &createdFips.Type, &createdFips.Annotation, &createdFips.Registration, &createdFips.PublishingDate, &createdFips.Applicant, &createdFips.Address)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	queryAuthors := `INSERT INTO fips_content_authors (content_id, author_fullname) VALUES ` + authorValues(createdFips.Id, authors)

	_, err = tx.Exec(queryAuthors)

	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &createdFips, nil
}

func (r *FipsRepository) GetAll() (*[]models.FipsContent, error) {
	var fipsContents []models.FipsContent
	query := `SELECT id, content_name, link, content_type, content_annotation, content_registration, content_publishing_date, content_applicant, content_address FROM fips_content`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var fips models.FipsContent
		err := rows.Scan(&fips.Id, &fips.Name, &fips.Link, &fips.Type, &fips.Annotation, &fips.Registration, &fips.PublishingDate, &fips.Applicant, &fips.Address)

		if err != nil {
			return nil, err
		}

		fipsContents = append(fipsContents, fips)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	for i, fips := range fipsContents {
		authorQuery := `SELECT author_fullname FROM fips_content_authors WHERE content_id = $1`
		authorRows, err := r.db.Query(authorQuery, fips.Id)

		if err != nil {
			return nil, err
		}

		var authors []string

		for authorRows.Next() {
			var author string
			err := authorRows.Scan(&author)

			if err != nil {
				return nil, err
			}

			authors = append(authors, author)
		}

		fipsContents[i].Authors = authors
	}

	return &fipsContents, nil
}
