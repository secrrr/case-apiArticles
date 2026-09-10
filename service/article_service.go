package service

import (
	"errors"
	"api-articles/models"
	"api-articles/repository"
)

type ArticleService interface {
	CreateArticle(req models.ArticleRequest) error
	GetArticles(searchQuery, authorName string, limit, offset int) ([]models.Article, error)
}

type articleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

func (s *articleService) CreateArticle(req models.ArticleRequest) error {
	if req.Title == "" || req.Body == "" {
		return errors.New("title dan body wajib diisi")
	}
	if len(req.AuthorID) == 0 {
		return errors.New("author_id wajib diisi")
	}
	return s.repo.Create(req)
}

func (s *articleService) GetArticles(searchQuery, authorName string, limit, offset int) ([]models.Article, error) {
	return s.repo.FindAll(searchQuery, authorName, limit, offset)
}

