package internetArticleRepo

import (
	"database/sql"
	"research-data-saver/internal/models"
)

type InternetArticleRepo struct {
	db *sql.DB
}

func Init(db *sql.DB) *InternetArticleRepo {
	articles := InternetArticleRepo{
		db: db,
	}

	return &articles
}

func (repo *InternetArticleRepo) AddArticle(name string, annotation string, link string, publishingDate string, author string, searchingMachine string) (*models.InternetArticle, error) {
	var InternetArticle models.InternetArticle

	err := repo.db.QueryRow(
		`
		INSERT INTO internet_articles
		(article_name, article_annotation, article_link, article_publishing_date, article_author, searching_machine) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, article_name, article_annotation, article_link, article_publishing_date, article_author, searching_machine
		`,
		name,
		annotation,
		link,
		publishingDate,
		author,
		searchingMachine,
	).Scan(&InternetArticle.Id, &InternetArticle.Name, &InternetArticle.Annotation, &InternetArticle.Link, &InternetArticle.PublishingDate, &InternetArticle.Author, &InternetArticle.SearchingMachine)

	if err != nil {
		return nil, err
	}

	return &InternetArticle, nil
}

func (repo *InternetArticleRepo) GetArticles(queryName string) (*[]models.InternetArticle, error) {
	var articles []models.InternetArticle

	rows, err := repo.db.Query(
		`
		SELECT id, article_name, article_annotation, article_link, article_publishing_date, article_author, searching_machine
		FROM internet_articles
		WHERE article_name LIKE $1
		`,
		"%"+queryName+"%",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var article models.InternetArticle

		err = rows.Scan(&article.Id, &article.Name, &article.Annotation, &article.Link, &article.PublishingDate, &article.Author, &article.SearchingMachine)

		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &articles, nil
}
