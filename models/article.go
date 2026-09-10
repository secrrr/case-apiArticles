package models

import(
	"time"

	"github.com/google/uuid"
)

type Article struct{
	ID 			uuid.UUID 	`json:"id"`
	Title 		string 		`json:"title"`
	Body 		string 		`json:"body"`
	CreatedAt 	time.Time 	`json:"created_at"`
	AuthorID  	uuid.UUID  	`json:"-"`
	Author    	Author    	`json:"author"`
}

type ArticleRequest struct{
	Title    string `json:"title"`
	Body     string `json:"body"`
	AuthorID string `json:"author_id"`
}