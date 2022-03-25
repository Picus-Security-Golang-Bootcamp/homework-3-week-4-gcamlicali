package AuthorRepo

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (a *AuthorRepository) Migrations() error {
	err := a.db.AutoMigrate(&Author{})
	if err != nil {
		fmt.Println("Migration Error")
		return err
	}

	return nil
}

func (a *AuthorRepository) FillAuthorData() {
	authors := GetAllAuthorsFromJson()

	for _, author := range authors {
		a.db.Where(Author{AuthorID: author.AuthorID}).
			Attrs(Author{AuthorID: author.AuthorID, AuthorName: author.AuthorName}).
			FirstOrCreate(&author)
	}
}

func (a *AuthorRepository) GetAllAuthors() []Author {
	var author []Author
	result := a.db.Where("deleted_at IS NULL").Find(&author)

	if result.Error != nil {
		log.Fatal("Cekerken hata oldu")
	}

	return author
}
