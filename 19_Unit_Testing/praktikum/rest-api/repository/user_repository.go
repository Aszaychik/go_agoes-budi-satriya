package repository

import (
	"unit_testing/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	// Update(user *domain.User) error
	// Delete(user *domain.User)
	// FindById(id int) error
	FindByEmail(email string) (*domain.User, error)
	// FindAll() []domain.User
	// FindExistsByEmail(email string) bool
}

type UserRepositoryImpl struct{
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (repository *UserRepositoryImpl) Create(user *domain.User) error {
	result := repository.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// func (repository *UserRepositoryImpl) Update(user *domain.User) error {}

// func (repository *UserRepositoryImpl) Delete(user *domain.User) {}

// func (repository *UserRepositoryImpl) FindById(id int) error {}

func (repository *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	user := domain.User{}

	result := repository.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// func (repository *UserRepositoryImpl) FindAll() []domain.User{}