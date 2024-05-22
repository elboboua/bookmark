package main

import (
	"context"
	"fmt"
	"os"

	"github.com/elboboua/bookmark/pkg/bookservice"
	"github.com/elboboua/bookmark/pkg/bookservice/database"
)

func main() {
	db := database.CreateNewDatabase()
	books, err := db.GetAllBooks()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, book := range books {
		fmt.Printf("%+v\n", book)
	}

	bookService := bookservice.NewBookService()
	_, input := shiftInput(os.Args)
	if len(input) == 0 {
		fmt.Println("USAGE: bookmark COMMAND")
	}
	command, input := shiftInput(input)

	switch command {
	case "search":
		query, _ := shiftInput(input)
		fmt.Println("INFO: Searching for: "+query+"...")
		books, err :=bookService.SearchBooks(context.Background(), query)
		if err != nil {
			fmt.Println("ERROR: Search failed")
			return
		}
		for _, book := range books {
			fmt.Printf("%s: %s\n", book.Title, book.Authors[0])
		}
	default:
		fmt.Println("Error: \""+command+"\" is not a valid command")
	}
	
		
	
	

}

func shiftInput(input []string) (string, []string) {
	currentInput := input[0]
	input = input[1:]
	return currentInput, input
}