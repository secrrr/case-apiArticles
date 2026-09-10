package repository

import (
	"database/sql"
	"testing"
	"api-articles/models"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *sql.DB {
	connStr := "postgres://postgres:password@localhost:5432/tes_kumparan?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Gagal koneksi ke database test: %v", err)
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("Database test tidak merespon: %v", err)
	}

	return db
}

func TestArticleRepository_Integration(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewArticleRepository(db)

	req := models.ArticleRequest{
		Title:    "Test Integrasi Judul",
		Body:     "Test integrasi body artikel",
		AuthorID: "11111111-1111-1111-1111-111111111111",
	}

	err := repo.Create(req)
	assert.NoError(t, err)

	articles, err := repo.FindAll("Test Integrasi", "", 10, 0)
	
	assert.NoError(t, err, "Proses SELECT FindAll harusnya tidak error")
	assert.NotEmpty(t, articles, "Data artikel yang baru dibuat harusnya ditemukan")
	assert.Equal(t, "Test Integrasi Judul", articles[0].Title)
}