package main

     // Packages
import (
	"fmt"
	"strings"
	"bufio"
	"os"

)

      // Defining the Book struct with fields for title (String), author (String), and publication year (Integer).
type Book struct {
	Title          string
	Author         string
	PublicationYear int
}

func main() {
	// List of Books and it's details
	// Title of the Book, Author, Year of the Publication
	books := []Book{
		{
			Title:          "Harry Potter and the Philosopher's Stone",
			Author:         "J.K. Rowling",
			PublicationYear: 1997,
		},
		{
			Title:          "Harry Potter and the Chamber of Secrets",
			Author:         "J.K. Rowling",
			PublicationYear: 1998,
		},
		{
			Title:          "Harry Potter and the Prisoner of Azkaban",
			Author:         "J.K. Rowling",
			PublicationYear: 1999,
		},
		{
			Title:          "Harry Potter and the Goblet of Fire",
			Author:         "J.K. Rowling",
			PublicationYear: 2000,
		},
		{
			Title:          "Harry Potter and the Order of the Phoenix",
			Author:         "J.K. Rowling",
			PublicationYear: 2003,
		},
		{
			Title:          "Harry Potter and the Half-Blood Prince",
			Author:         "J.K. Rowling",
			PublicationYear: 2005,
		},
		{
			Title:          "Harry Potter and the Deathly Hallows",
			Author:         "J.K. Rowling",
			PublicationYear: 2007,
		},
}
    // for Loop is being used for determining which of the details the user put
    // bufio.NewScanner is used for the spaces in the author's name and title to make the program easier to read and search
    // os.Stdin is used to access the standard input stream to read user input.
    // Boolean is used for true and false to determine if the user input is in the List or not
    for {
		var userInput string
		fmt.Println("Search by author, book title, or publication year:")

    // Handle user input with spaces
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			userInput = scanner.Text()
		}

		found := false
		for _, book := range books {
			if book.Author == userInput || book.Title == userInput || fmt.Sprintf("%d", book.PublicationYear) == userInput {
				fmt.Println("\n---------- Book Details ----------\n")
				fmt.Println("Title:", book.Title)
				fmt.Println("Author:", book.Author)
				fmt.Println("Publication Year:", book.PublicationYear)
				fmt.Println("\n*********************************************")
				found = true
			}
		}

		if !found {
			fmt.Println("No results found.")
		}

	// Asking If the user want to search again
	// If yes the program will continue to ask about the title, author or publication year
	// If No the program will end and print Thank You!
		fmt.Println("Do you want to search again? (yes/no)")
		var decision string
		if scanner.Scan() {
			decision = scanner.Text()
		}
		if strings.ToLower(strings.TrimSpace(decision)) != "yes" {
			fmt.Println("Thank you!")
			break
		}
	}
}
