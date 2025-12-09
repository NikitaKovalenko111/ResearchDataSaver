package documentRepository

import (
	"database/sql"
	"research-data-saver/internal/models"
)

type DocumentRepository struct {
	db *sql.DB
}

func Init(db *sql.DB) *DocumentRepository {
	return &DocumentRepository{db: db}
}

func (r *DocumentRepository) Create(name string, annotation string, link string, publishingDate string, author string) (*models.Document, error) {
	var createdDocument models.Document
	query := `INSERT INTO documents (document_name, document_annotation, document_link, document_publishing_date, document_author) VALUES ($1, $2, $3, $4, $5) RETURNING id, document_name, document_annotation, document_link, document_publishing_date, document_author`
	err := r.db.QueryRow(query, name, annotation, link, publishingDate, author).Scan(&createdDocument.Id, &createdDocument.Name, &createdDocument.Annotation, &createdDocument.Link, &createdDocument.PublishingDate, &createdDocument.Author)

	if err != nil {
		return nil, err
	}

	return &createdDocument, nil
}

func (r *DocumentRepository) GetAll(queryName string) (*[]models.Document, error) {
	var document []models.Document
	query := `SELECT id, document_name, document_annotation, document_link, document_publishing_date, document_author FROM documents WHERE document_name LIKE $1`
	rows, err := r.db.Query(query, "%"+queryName+"%")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var doc models.Document
		err := rows.Scan(&doc.Id, &doc.Name, &doc.Annotation, &doc.Link, &doc.PublishingDate, &doc.Author)

		if err != nil {
			return nil, err
		}
		document = append(document, doc)
	}

	return &document, nil
}
