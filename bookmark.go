package main

import (
	"context"
	"fmt"

	"github.com/elboboua/bookmark/pkg/bookservice"
)

func main() {
	fmt.Println("Hello, world")
	bookService := bookservice.NewBookService()
	books, err :=bookService.SearchBooks(context.Background(),"ninjas")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, book := range books {
		fmt.Printf("%s: %s\n", book.Title, book.Authors[0])
	}

}