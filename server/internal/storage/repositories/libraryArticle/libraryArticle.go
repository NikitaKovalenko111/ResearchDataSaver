package libraryArticleRepo

import (
	"database/sql"
	"fmt"
	"research-data-saver/internal/models"
)

type LibraryArticleRepo struct {
	db *sql.DB
}

func Init(db *sql.DB) *LibraryArticleRepo {
	repo := LibraryArticleRepo{
		db: db,
	}

	return &repo
}

func authorValues(authors []string, id int) string {
	values := ""
	for i, author := range authors {
		values += fmt.Sprintf("('%s', %d)", author, id)

		if i < len(authors)-1 {
			values += ", "
		}
	}
	return values
}

func (repo *LibraryArticleRepo) AddArticle(name string, annotation string, link string, publishingDate string, lang string, udk string, publisherObject string, publisher string, supervisor string, authors []string) (*models.LibraryArticle, error) {
	var article models.LibraryArticle

	tx, err := repo.db.Begin()

	if err != nil {
		return nil, err
	}

	queryAddArticle :=
		`
		INSERT INTO library_articles 
		(article_name, article_annotation, article_link, article_publishing_date, article_lang, article_udk, article_publisher_object, article_publisher, article_supervisor) 
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, article_name, article_annotation, article_link, article_publishing_date, article_lang, article_udk, article_publisher_object, article_publisher, article_supervisor
	`

	err = tx.QueryRow(queryAddArticle, name, annotation, link, publishingDate, lang, udk, publisherObject, publisher, supervisor).Scan(&article.Id, &article.Name, &article.Annotation, &article.Link, &article.PublishingDate, &article.Lang, &article.UDK, &article.PublisherObject, &article.Publisher, &article.Supervisor)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	queryAddAuthors := `INSERT INTO library_articles_authors (author_fullname, article_id) VALUES ` + authorValues(authors, article.Id)

	_, err = tx.Exec(queryAddAuthors)

	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &article, nil
}

func (repo *LibraryArticleRepo) GetArticles(queryName string, queryDate string, queryLang string, querySupervisor string) (*[]models.LibraryArticle, error) {
	var articles []models.LibraryArticle

	rows, err := repo.db.Query(
		`
		SELECT
		id, article_name, article_annotation, article_link, article_publishing_date, article_lang, article_udk, article_publisher_object, article_publisher, article_supervisor
		FROM library_articles
		WHERE article_name LIKE $1 AND TO_CHAR(article_publishing_date, 'YYYY-MM-DD') LIKE $2 AND article_lang LIKE $3 AND article_supervisor LIKE $4
		`,
		"%"+queryName+"%", "%"+queryDate+"%", "%"+queryLang+"%", "%"+querySupervisor+"%",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var article models.LibraryArticle

		err = rows.Scan(&article.Id, &article.Name, &article.Annotation, &article.Link, &article.PublishingDate, &article.Lang, &article.UDK, &article.PublisherObject, &article.Publisher, &article.Supervisor)

		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	for i, article := range articles {
		authorRows, err := repo.db.Query(
			`
			SELECT author_fullname FROM library_articles_authors WHERE article_id = $1
			`,
			article.Id,
		)

		if err != nil {
			return nil, err
		}

		defer authorRows.Close()

		var authors []string

		for authorRows.Next() {
			var author string

			err = authorRows.Scan(&author)

			if err != nil {
				return nil, err
			}

			authors = append(authors, author)
		}

		articles[i].Authors = authors
	}

	return &articles, nil
}
