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

func Prepare(db *sql.DB) error {
	_, err := db.Exec(
		`
		CREATE TABLE IF NOT EXISTS public.documents
		(
			id integer NOT NULL DEFAULT nextval('documents_id_seq'::regclass),
			document_name text COLLATE pg_catalog."default" NOT NULL,
			document_annotation text COLLATE pg_catalog."default" NOT NULL,
			document_link text COLLATE pg_catalog."default" NOT NULL,
			document_publishing_date date,
			document_author character varying(256) COLLATE pg_catalog."default",
			CONSTRAINT documents_pkey PRIMARY KEY (id),
			CONSTRAINT documents_document_name_key UNIQUE (document_name)
		)

		CREATE TABLE IF NOT EXISTS public.fips_content
		(
			id integer NOT NULL DEFAULT nextval('fips_content_id_seq'::regclass),
			content_type character varying(64) COLLATE pg_catalog."default" NOT NULL,
			content_name character varying(256) COLLATE pg_catalog."default" NOT NULL,
			content_annotation text COLLATE pg_catalog."default",
			content_registration character varying(32) COLLATE pg_catalog."default" NOT NULL,
			content_publishing_date date NOT NULL,
			content_applicant character varying(64) COLLATE pg_catalog."default",
			content_address character varying(256) COLLATE pg_catalog."default" NOT NULL,
			link text COLLATE pg_catalog."default",
			CONSTRAINT fips_content_pkey PRIMARY KEY (id),
			CONSTRAINT fips_content_content_name_key UNIQUE (content_name),
			CONSTRAINT fips_content_content_registration_key UNIQUE (content_registration)
		)

		CREATE TABLE IF NOT EXISTS public.fips_content_authors
		(
			id integer NOT NULL DEFAULT nextval('fips_content_authors_id_seq'::regclass),
			author_fullname character varying(128) COLLATE pg_catalog."default" NOT NULL,
			content_id integer,
			CONSTRAINT fips_content_authors_pkey PRIMARY KEY (id),
			CONSTRAINT fips_content_authors_content_id_fkey FOREIGN KEY (content_id)
				REFERENCES public.fips_content (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION
		)

		CREATE TABLE IF NOT EXISTS public.internet_articles
		(
			id integer NOT NULL DEFAULT nextval('internet_articles_id_seq'::regclass),
			article_name character varying(256) COLLATE pg_catalog."default" NOT NULL,
			article_annotation text COLLATE pg_catalog."default" NOT NULL,
			article_link text COLLATE pg_catalog."default" NOT NULL,
			article_publishing_date date,
			article_author character varying(256) COLLATE pg_catalog."default",
			searching_machine character varying(64) COLLATE pg_catalog."default",
			CONSTRAINT internet_articles_pkey PRIMARY KEY (id),
			CONSTRAINT internet_articles_article_name_key UNIQUE (article_name)
		)

		CREATE TABLE IF NOT EXISTS public.library_articles
		(
			id integer NOT NULL DEFAULT nextval('library_articles_id_seq'::regclass),
			article_name character varying(256) COLLATE pg_catalog."default" NOT NULL,
			article_annotation text COLLATE pg_catalog."default" NOT NULL,
			article_link text COLLATE pg_catalog."default" NOT NULL,
			article_publishing_date date,
			article_lang character varying(96) COLLATE pg_catalog."default" NOT NULL,
			article_udk character varying(30) COLLATE pg_catalog."default" NOT NULL,
			article_publisher_object character varying(128) COLLATE pg_catalog."default",
			article_publisher character varying(256) COLLATE pg_catalog."default",
			article_supervisor character varying(128) COLLATE pg_catalog."default",
			CONSTRAINT library_articles_pkey PRIMARY KEY (id),
			CONSTRAINT library_articles_article_link_key UNIQUE (article_link),
			CONSTRAINT library_articles_article_name_key UNIQUE (article_name)
		)

		CREATE TABLE IF NOT EXISTS public.library_articles_authors
		(
			id integer NOT NULL DEFAULT nextval('library_articles_authors_id_seq'::regclass),
			author_fullname character varying(128) COLLATE pg_catalog."default" NOT NULL,
			article_id integer,
			CONSTRAINT library_articles_authors_pkey PRIMARY KEY (id),
			CONSTRAINT library_articles_authors_article_id_fkey FOREIGN KEY (article_id)
				REFERENCES public.library_articles (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION
		)
		`,
	)

	return err
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
