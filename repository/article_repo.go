package repository

import (
	"database/sql"
	"fmt"
	"api-articles/models"
)

type ArticleRepository interface {
	Create(article models.ArticleRequest) error
	FindAll(searchQuery, authorName string, limit, offset int) ([]models.Article, error)
}

type articleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) Create(req models.ArticleRequest) error {
	query := `INSERT INTO articles (title, body, author_id) VALUES ($1, $2, $3)`
	// blank identifier agar (_) agar tidak terjadi eror dan hanya perlu mengembalikan err saja
	_, err := r.db.Exec(query, req.Title, req.Body, req.AuthorID)
	return err
}

func (r *articleRepository) FindAll(searchQuery, authorName string, limit, offset int) ([]models.Article, error) {
	baseQuery := `
		SELECT a.id, a.title, a.body, a.created_at, a.author_id, u.name 
		FROM articles a 
		JOIN authors u ON a.author_id = u.id 
		WHERE 1=1
	`
	var args []interface{}
	counter := 1

	if searchQuery != "" {
		// CREATE INDEX articles_fts_idx ON articles USING GIN (to_tsvector('simple', title || ' ' || body));
		// menggunakan fitur full text search milik postgresql
		baseQuery += fmt.Sprintf(" AND to_tsvector('simple', a.title || ' ' || a.body) @@ plainto_tsquery('simple', $%d)", counter)
		args = append(args, searchQuery)
		counter ++
	}

	// Filter pencarian berdasarkan nama author
	if authorName != "" {
		// nama tidak bisa menggunakan fts, dikarenakan gaada bentuk dasarnya
		baseQuery += fmt.Sprintf(" AND u.name ILIKE $%d", counter)
		args = append(args, "%"+authorName+"%")
		counter++
	}

	baseQuery += fmt.Sprintf(" ORDER BY a.created_at DESC LIMIT $%d OFFSET $%d", counter, counter+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []models.Article{}

	for rows.Next() {
		var art models.Article
		err := rows.Scan(&art.ID, &art.Title, &art.Body, &art.CreatedAt, &art.AuthorID, &art.Author.Name)
		if err != nil {
			return nil, err
		}
		art.Author.ID = art.AuthorID
		articles = append(articles, art)
	}

	return articles, nil
}