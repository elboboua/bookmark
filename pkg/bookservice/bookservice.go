package bookservice

import (
	"context"
	"net/http"
	"os"
)

type Book struct {
	ID string
	Title string
	Subtitle string
	Description string
	Authors []string
	PurchaseLink string
	PublishedDate string
	Language string
	PageCount int
	Categories []string
}

type BookService interface {
	SearchBooks(ctx context.Context, query string) ([]Book, error)
}

func NewBookService() BookService {
	apiKey := os.Getenv("GOOGLE_BOOK_API_KEY")
	if apiKey == "" {
		panic("no api key provided for google book api")
	}
	return &GoogleBookService{
		apiKey: apiKey,
		client: http.Client{},
	}
}