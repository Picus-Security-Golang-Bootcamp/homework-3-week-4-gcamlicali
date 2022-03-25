package BookRepo

import (
	AuthorRepo "DBHW/domain/author"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
)

type Book struct {
	gorm.Model
	BookID    int               `json:"ID"`
	Name      string            `json:"Name"`
	Pages     int               `json:"Pages"`
	Stock     uint              `json:"Stock"`
	Price     int               `json:"Price"`
	StockID   string            `json:"StockID"`
	ISBN      int               `json:"ISBN"`
	AuthorsID int               `json:"AuthorID"`
	Author    AuthorRepo.Author `gorm:"foreignKey:AuthorID;references:AuthorsID"`
}

func (b *Book) ToString() string {
	//return fmt.Sprintf("BookID : %d, Name : %s, Pages : %d, Price : %d,Stock : %d,StockID : %s,ISBN : %d,AuthorID : %d", book.BookID, book.Name, book.Pages, book.Price, book.Stock, book.StockID, book.ISBN, book.AuthorID)

	return fmt.Sprintf("BookID : %d, Name : %s, Pages : %d, Price : %d, Stock : %d, StockID : %s, ISBN : %d, Author : %s", b.BookID, b.Name, b.Pages, b.Price, b.Stock, b.StockID, b.ISBN, b.Author.AuthorName)
}

func GetAllBooksFromJson() []Book {

	//var list Books
	list := []Book{}

	jsonFile, err := os.Open("books.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Patates while opening json: ", err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//Parse to List struct
	err = json.Unmarshal(byteValue, &list)

	//fmt.Println(byteValue)
	if err != nil {
		log.Fatal("Patates while unmarshal json: ", err)
	}
	//b.bookList = bookList
	//fmt.Println(list)

	return list
}
