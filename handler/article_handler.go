package handler

import (
	"encoding/json"
	"net/http"
	"api-articles/models"
	"api-articles/service"
)

// blueprint articleservice yang digunakan untuk function get n post article
type ArticleHandler struct {
	service service.ArticleService
}

// digunakan untuk pertama kali server dinyalakan, dan pada func ini memasukkan parameter wajib
// yang boleh dilakukan oleh articlehandler -> articleservice.  
func NewArticleHandler(s service.ArticleService) *ArticleHandler {
	if s == nil{
		panic("Parameter articlehandler pada main.go belum di definisikan")
	}
	return &ArticleHandler{service: s}
}

// handler menangani create article (POST)
func (h *ArticleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.ArticleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateArticle(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Artikel berhasil dibuat"})
}

// handler untuk mendapatkan seluruh article (GET) serta untuk mencari artikel berdasarkan body/title && mencari authornya
func (h *ArticleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query().Get("query")
	authorParam := r.URL.Query().Get("author")

	articles, err := h.service.GetArticles(queryParam, authorParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}