package service

import (
	"fmt"
	"simple_clean_code/model/domain"
	"simple_clean_code/model/web"
	"simple_clean_code/repository"
)

type UserService interface {
	CreateUser(request web.UserCreateRequest) (*domain.User, error)
	GetAllUser() ([]domain.User, error)
}

type UserServiceImpl struct {
	Repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &UserServiceImpl{Repository: repository}
}

func (service UserServiceImpl) CreateUser(request web.UserCreateRequest) (*domain.User, error) {
	existingUser, err := service.Repository.FindByEmail(request.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("Email already exist")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	user := domain.User{
		Email: request.Email,
		Password: request.Password,
	}

	userResult, err := service.Repository.Save(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return userResult, nil
}

func (service UserServiceImpl) GetAllUser() ([]domain.User, error) {
	users, err := service.Repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	return users, nil
}