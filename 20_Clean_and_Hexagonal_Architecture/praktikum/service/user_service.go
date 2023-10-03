package service

import (
	"fmt"
	"simple_clean_code/helper"
	"simple_clean_code/model/domain"
	"simple_clean_code/model/web"
	"simple_clean_code/repository"
)

type UserService interface {
	CreateUser(request web.UserCreateRequest) (*domain.User, error)
	GetAllUser() ([]domain.User, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (service *UserServiceImpl) CreateUser(request web.UserCreateRequest) (*domain.User, error) {

	// Check if the email already exists
	existingUser, _ := service.UserRepository.FindByEmail(request.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("Email already exists")
	}

	// Convert request to domain
	user := helper.UserCreateRequestToUserDomain(request)

	// Convert password to hash
	user.Password = helper.HashPassword(user.Password)

	result, err := service.UserRepository.Save(user)
	if err != nil {
		return nil, fmt.Errorf("Error when creating user: %s", err.Error())
	}

	return result, nil
}

func (service *UserServiceImpl) GetAllUser() ([]domain.User, error) {
	users, err := service.UserRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	return users, nil
}