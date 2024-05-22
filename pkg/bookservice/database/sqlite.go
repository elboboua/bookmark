package database

import (
	"log"

	"github.com/elboboua/bookmark/pkg/bookservice"
)

type Sqlite struct {
	db IDatabaseDriver
}

func (conn *Sqlite) GetAllBooks() ([]bookservice.Book, error) {
	// get all books
	rows, err := conn.db.Query("SELECT id, title, subtitle, description, purchase_link, published_date, language, page_count from book")
	if err != nil {
		log.Println("failed to get all books: "+ err.Error())
		return nil, err
	}
	
	books := []bookservice.Book{}
	for rows.Next() {
		book := bookservice.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Subtitle, &book.Description, &book.PurchaseLink, &book.PublishedDate, &book.Language, &book.PageCount)
		books = append(books, book)
	}

	return books, nil
}