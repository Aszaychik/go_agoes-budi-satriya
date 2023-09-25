package models

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id       string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	Title     string `gorm:"type:varchar(255)" json:"title"`
	Author    string `gorm:"type:varchar(255);" json:"author"`
	Publisher string `gorm:"type:varchar(255);" json:"publisher"`
}

type BooksModel struct {
	db *gorm.DB
}

func (bm *BooksModel) Init(db *gorm.DB) {
	bm.db = db
}


func (b *Book) GenerateBookID() {
	if b.Id == "" {
		b.Id = uuid.NewString()
		return
	}
	fmt.Println("ID already exist")
}

func (bm *BooksModel) FindAllBooks() ([]Book, error) {
	books := []Book{}
	err := bm.db.Find(&books).Error
	if err != nil {
		logrus.Error("Model : Find all data error", err)
		return nil, err
	}

	return books, nil
}

func (bm *BooksModel) FindById(id string) (*Book, error) {
	book := Book{}
	err := bm.db.Table("books").Where("id = ?", id).First(&book).Error
	if err != nil {
		logrus.Error("Model : Find by id error", err)
		return nil, err
	}

	return &book, nil
}

func (bm *BooksModel) Create(newBook Book) (*Book, error) {
	newBook.GenerateBookID()

	err := bm.db.Create(&newBook).Error
	if err != nil {
		logrus.Error("Model : Insert data error", err)
		return nil, err
	}

	return &newBook, nil
}

func (bm *BooksModel) Update(updatedData Book) (*Book, error) {
	result := bm.db.Where("id = ?", updatedData.Id).Updates(Book{Title: updatedData.Title, Author: updatedData.Author, Publisher: updatedData.Publisher})
	if result.RowsAffected == 0 {
		logrus.Error("Model : Update data error", "no data affected")
		return nil, errors.New("no data affected")
	}
	updatedBook := Book{}
	err := bm.db.Where("id = ?", updatedData.Id).First(&updatedBook).Error
	if err != nil {
		logrus.Error("Model : Update data error", err)
		return nil, err
	}

	return &updatedBook, nil
}

func (bm *BooksModel) Delete(id string) error {
  book := Book{}
  book.Id = id
	
  err := bm.db.Delete(&book).Error
	if err != nil {
		logrus.Error("Model : Delete data error", err.Error())
		return err
	}

	return nil
}