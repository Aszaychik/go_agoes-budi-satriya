package models

import (
	"fmt"
	"orm_code_structure/helpers"

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

func (u *Book) GenerateBookID() {
	if u.Id == "" {
		u.Id = uuid.NewString()
		return
	}
	fmt.Println("ID already exist")
}

func (bm *BooksModel) FindAll() []Book {
	books := []Book{}
	err := bm.db.Find(&books).Error
	helpers.LogIfError("Model : Find all data error, ", err)

	return books
}

func (bm *BooksModel) FindById(id string) *Book {
	book := Book{}
	err := bm.db.Table("books").Where("id = ?", id).First(&book).Error
	helpers.LogIfError("Model : Find by id error, ", err)

	return &book
}

func (bm *BooksModel) Create(newBook Book) *Book {
	newBook.GenerateBookID()

	err := bm.db.Create(&newBook).Error
	helpers.LogIfError("Model : Insert data error, ", err)

	return &newBook
}

func (bm *BooksModel) Update(updatedData Book) *Book {
	result := bm.db.Where("id = ?", updatedData.Id).Updates(Book{Title: updatedData.Title, Author: updatedData.Author, Publisher: updatedData.Publisher})
	if result.RowsAffected == 0 {
		logrus.Error("Model : Update data error", "no data affected")
		return nil
	}
	updatedBook := Book{}
	err := bm.db.Where("id = ?", updatedData.Id).First(&updatedBook).Error
	helpers.LogIfError("Model : Update data error, ", err)

	return &updatedBook
}

func (bm *BooksModel) Delete(id string) {
	deletedBook := Book{}
	deletedBook.Id = id
	err := bm.db.Delete(&deletedBook).Error
	helpers.LogIfError("Model : Delete error, ", err)
}