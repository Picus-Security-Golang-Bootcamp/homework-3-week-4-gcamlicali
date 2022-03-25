package postgres

import (
	AuthorRepo "DBHW/domain/author"
	BookRepo "DBHW/domain/book"
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
)

func ExecuteQueries(db *gorm.DB) {

	// Repositories
	authorRepo := AuthorRepo.NewAuthorRepository(db)
	bookRepo := BookRepo.NewBookRepository(db)

	err := authorRepo.Migrations()
	if err != nil {
		log.Fatalln("Program Closed by Migration Error:", err)
		os.Exit(1)
	}

	err = bookRepo.Migrations()
	if err != nil {
		log.Fatalln("Program Closed by Migration Error:", err)
		os.Exit(1)
	}

	authorRepo.FillAuthorData()
	bookRepo.FillBookData()

	//-----------------------------------------------------------//
	fmt.Println("\n******Listing All Books in DB ******\n")
	books := bookRepo.FindAll()
	for _, book := range books {
		fmt.Println(book.ToString())
	}

	//-----------------------------------------------------------//
	fmt.Println("\n******Buying books by id from DB ******\n")

	book, err := bookRepo.BuyBookByID(1, 5)
	fmt.Println("After buy 5 books")
	fmt.Println(book.ToString())

	//-----------------------------------------------------------//
	fmt.Println("\n******Getting book ID:2 from DB******\n")
	book, _ = bookRepo.GetBookByID(2)
	fmt.Println(book.ToString())

	//-----------------------------------------------------------//
	fmt.Println("\n******Searching books include 'Great' in name from DB******\n")
	books, err = bookRepo.SearchBooksByName("Great")
	for _, book := range books {
		fmt.Println(book.ToString())
	}

	//-----------------------------------------------------------//
	fmt.Println("\n******Deleting book ID:5 from DB******\n")
	bookRepo.DeleteBookById(5)
	fmt.Println("Remaining Books : ")
	books = bookRepo.FindAll()
	for _, book := range books {
		fmt.Println(book.ToString())
	}

	//bookRepo.GetBooksByAuthor("tolstoy")
}
