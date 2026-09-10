package main

import (
	"fmt"
	"log"
	"net/http"

	"api-articles/config"
	"api-articles/handler"
	"api-articles/repository"
	"api-articles/service"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

    // Memberikan jalan ke repository untuk bisa melakukan query dengan konek db 
	articleRepo := repository.NewArticleRepository(db)
    // Memberikan akses file service untuk memasukkan repository ke dalamnya
	articleService := service.NewArticleService(articleRepo)
    // Action untuk melakukan sesuatu dengan mengambil service yang sudah berisikan logic + query dari repository
	articleHandler := handler.NewArticleHandler(articleService)

	mux := http.NewServeMux()

    mux.HandleFunc("POST /articles", articleHandler.Create)
	mux.HandleFunc("GET /articles", articleHandler.GetAll)

	fmt.Println("Server berjalan di http://localhost:8080 ...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}