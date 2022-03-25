package BookRepo

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) Migrations() error {
	err := b.db.AutoMigrate(&Book{})

	if err != nil {
		fmt.Println("Migration Error")
		return err
	}

	return nil
}

func (b *BookRepository) FillBookData() {
	books := GetAllBooksFromJson()

	for _, book := range books {
		b.db.Where(Book{BookID: book.BookID}).
			Attrs(Book{BookID: book.BookID, Name: book.Name, Pages: book.Pages, Price: book.Price,
				Stock: book.Stock, StockID: book.StockID, ISBN: book.ISBN, AuthorsID: book.AuthorsID}).
			FirstOrCreate(&book)
	}
}

func (b *BookRepository) FindAll() []Book {
	var books []Book
	result := b.db.Preload("Author").Where("deleted_at IS NULL").Find(&books)

	if result.Error != nil {
		log.Fatal("Cekerken hata oldu")
	}

	return books
}

func (b *BookRepository) GetBookByID(id int) (Book, error) {
	var book Book
	err := b.db.Preload("Author").Where(Book{BookID: id}).Find(&book).Error

	if err != nil {
		log.Fatal("Unknown ID")
		return book, err
	}

	return book, nil
}

func (b *BookRepository) BuyBookByID(id int, quantity uint) (Book, error) {
	var book Book
	err := b.db.Preload("Author").Where(Book{BookID: id}).Find(&book).Error // add where delete_at != null query
	// add id not found

	if err != nil {
		return book, err
	}

	if book.Stock < quantity {
		return book, errors.New("Not enough stock")
	}

	book.Stock -= quantity
	err = b.db.Save(&book).Error

	if err != nil {
		log.Fatal("Error while DB Updating: ", err)
		return book, err
	}

	return book, nil
}

func (b *BookRepository) DeleteBookById(id int) error {
	var book Book

	result := b.db.Where(Book{BookID: id}).Find(&book).Model(&book).Update("deleted_at", time.Now())
	if result.Error != nil {
		log.Fatal("Delete book by id is not completed: ", result.Error)
		return result.Error
	}

	return nil
}

func (b *BookRepository) SearchBooksByName(bookName string) ([]Book, error) {
	var books []Book

	result := b.db.Preload("Author").Where("name ILIKE ? ", "%"+bookName+"%").Find(&books)
	if result.Error != nil {

		fmt.Println("Sorguda hata var: ", result.Error)
		return books, result.Error
	}

	return books, nil
}

func (b *BookRepository) GetBooksByAuthor(authName string) ([]Book, error) {
	var books []Book
	nameStr := fmt.Sprintf("%%%s%%", authName)
	err := b.db.Preload("Author").Where("author_authorname ILIKE ? ", nameStr).Find(&books).Error

	if err != nil {
		log.Fatal("Unknown ID")
		//return book, err
	}

	return books, nil
}
