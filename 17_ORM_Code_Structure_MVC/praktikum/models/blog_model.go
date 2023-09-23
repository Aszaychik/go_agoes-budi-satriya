package models

import (
	"fmt"
	"orm_code_structure/helpers"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Id       string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	UserId      string `gorm:"primaryKey;type:varchar(255)" json:"user_id"`
	Title     string `gorm:"type:varchar(255)" json:"title"`
	Content    string `gorm:"type:varchar(255);" json:"content"`
}

type BlogsModel struct {
	db *gorm.DB
}

func (bm *BlogsModel) Init(db *gorm.DB) {
	bm.db = db
}

func (b *Blog) GenerateBlogID() {
	if b.Id == "" {
		b.Id = uuid.NewString()
		return
	}
	fmt.Println("ID already exist")
}

func (bm *BlogsModel) FindAll() []Blog {
	blogs := []Blog{}
	err := bm.db.Find(&blogs).Error
	helpers.LogIfError("Model : Find all data error, ", err)

	return blogs
}

func (bm *BlogsModel) FindById(id string) *Blog {
	blog := Blog{}
	err := bm.db.Table("blogs").Where("id = ?", id).First(&blog).Error
	helpers.LogIfError("Model : Find by id error, ", err)

	return &blog
}

func (bm *BlogsModel) Create(newBlog Blog) *Blog {
	newBlog.GenerateBlogID()

	err := bm.db.Create(&newBlog).Error
	helpers.LogIfError("Model : Insert data error, ", err)

	return &newBlog
}

func (bm *BlogsModel) Update(updatedData Blog) *Blog {
	result := bm.db.Where("id = ?", updatedData.Id).Updates(Blog{UserId: updatedData.UserId, Title: updatedData.Title, Content: updatedData.Content})
	if result.RowsAffected == 0 {
		logrus.Error("Model : Update data error", "no data affected")
		return nil
	}
	updatedBlog := Blog{}
	err := bm.db.Where("id = ?", updatedData.Id).First(&updatedBlog).Error
	helpers.LogIfError("Model : Update data error, ", err)

	return &updatedBlog
}

func (bm *BlogsModel) Delete(id string) {
  blog := Blog{}
  blog.Id = id
	
  err := bm.db.Delete(&blog).Error
  helpers.LogIfError("Model : Delete error, ", err)
}