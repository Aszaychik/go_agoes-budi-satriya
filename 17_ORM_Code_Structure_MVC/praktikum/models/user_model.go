package models

import (
	"fmt"
	"orm_code_structure/helpers"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id      string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	Name    string `gorm:"type:varchar(255)" json:"name"`
	Email      string `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Password   string	`gorm:"type:varchar(24);" json:"password"`
	Blogs []Blog `gorm:"foreignKey:user_id;references:id"`
}

type UsersModel struct {
	db *gorm.DB
}

func (um *UsersModel) Init(db *gorm.DB) {
	um.db = db
}

func (u *User) GenerateID() {
	if u.Id == "" {
		u.Id = uuid.NewString()
		return
	}
	fmt.Println("ID already exist")
}

func (um *UsersModel) FindAll() []User {
	users := []User{}
	err := um.db.Find(&users).Error
	helpers.LogIfError("Model : Find all data error, ", err)

	return users
}

func (um *UsersModel) FindById(id string) *User {
	user := User{}
	err := um.db.Table("users").Where("id = ?", id).First(&user).Error
	helpers.LogIfError("Model : Find by id error, ", err)

	return &user
}

func (um *UsersModel) Create(newUser User) *User {
	newUser.GenerateID()

	err := um.db.Create(&newUser).Error
	helpers.LogIfError("Model : Insert data error, ", err)

	return &newUser
}

func (um *UsersModel) Update(updatedData User) *User {
	result := um.db.Where("id = ?", updatedData.Id).Updates(User{Name: updatedData.Name, Email: updatedData.Email, Password: updatedData.Password})
	if result.RowsAffected == 0 {
		logrus.Error("Model : Update data error", "no data affected")
		return nil
	}
	updatedUser := User{}
	err := um.db.Where("id = ?", updatedData.Id).First(&updatedUser).Error
	helpers.LogIfError("Model : Update data error, ", err)

	return &updatedUser
}

func (um *UsersModel) Delete(id string) {
  user := User{}
  user.Id = id
	
  err := um.db.Delete(&user).Error
  helpers.LogIfError("Model : Delete error, ", err)
}