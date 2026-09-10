package service

import (
	"testing"
	"api-articles/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockArticleRepository struct {
	mock.Mock
}

func (m *MockArticleRepository) Create(article models.ArticleRequest) error {
	args := m.Called(article)
	return args.Error(0)
}

func (m *MockArticleRepository) FindAll(searchQuery, authorName string, limit, offset int) ([]models.Article, error) {
	args := m.Called(searchQuery, authorName, limit, offset)
	return args.Get(0).([]models.Article), args.Error(1)
}

func TestCreateArticle_Failed(t *testing.T) {
	mockRepo := new(MockArticleRepository)
	serv := NewArticleService(mockRepo)

	req := models.ArticleRequest{
		Title:    "",
		Body:     "YANG semangat!",
		AuthorID: "123",
	}

	err := serv.CreateArticle(req)

	assert.Error(t, err)
	assert.Equal(t, "title dan body wajib diisi keduanya", err.Error())
}

func TestCreateArticle_Success(t *testing.T) {
	mockRepo := new(MockArticleRepository)
	req := models.ArticleRequest{
		Title:    "Golang Testing",
		Body:     "Belajar unit test di Go",
		AuthorID: "11111111-1111-1111-1111-111111111111",
	}

	// beri informasi ke mockrepo kalau func create membawa req maka berikan statusnya sukses
	mockRepo.On("Create", req).Return(nil)

	// dependency injection mockrepo
	serv := NewArticleService(mockRepo)
	// membuat article
	err := serv.CreateArticle(req)

	// memastikan tidak ada pesan eror yang muncul
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}