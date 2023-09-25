package models

import (
	"errors"
	"fmt"
	"golang_middleware/helpers"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id      string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	Name    string `gorm:"type:varchar(255)" json:"name"`
	Email      string `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Password   string	`gorm:"type:varchar(255);" json:"password"`
	// Blogs []Blog `gorm:"foreignKey:user_id;references:id"`
}

type LoginUser struct{
	Email string `json:"email"`
	Password string `json:"password"`
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

func (um *UsersModel) FindAll() ([]User, error) {
	users := []User{}
	err := um.db.Find(&users).Error
	if err != nil {
		logrus.Error("Model : Find all data error", err)
		return nil, err
	}

	return users, nil
}

func (um *UsersModel) FindById(id string) (*User, error) {
	user := User{}
	err := um.db.Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		logrus.Error("Model : Find by id error", err)
		return nil, err
	}

	return &user, nil
}

func (um *UsersModel) Create(newUser User) (*User, error) {
	newUser.GenerateID()

	err := um.db.Create(&newUser).Error
	if err != nil {
		logrus.Error("Model : Insert data error", err)
		return nil, err
	}

	return &newUser, nil
}

func (um *UsersModel) Update(updatedData User) (*User, error) {
	result := um.db.Where("id = ?", updatedData.Id).Updates(User{Name: updatedData.Name, Email: updatedData.Email, Password: updatedData.Password})
	if result.RowsAffected == 0 {
		logrus.Error("Model : Update data error", "no data affected")
		return nil, errors.New("no data affected")
	}
	updatedUser := User{}
	err := um.db.Where("id = ?", updatedData.Id).First(&updatedUser).Error
	if err != nil {
		logrus.Error("Model : Update data error", err)
		return nil, err
	}

	return &updatedUser, nil
}

func (um *UsersModel) Delete(id string) error {
  user := User{}
  user.Id = id
	
  err := um.db.Delete(&user).Error
	if err != nil {
		logrus.Error("Model : Delete data error", err.Error())
		return err
	}

	return nil
}

func (um *UsersModel) Login(email string, password string) (*User, error) {
	user := User{}
	err := um.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		logrus.Error("Model : Login data error", err.Error())
		return nil, err
	}

	err = helpers.ComparePassword(user.Password, password)
	if err != nil {
		logrus.Error("Model : Login data error", err.Error())
		return nil, err
	}

	return &user, nil
}