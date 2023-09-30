package service

import (
	"fmt"
	"unit_testing/helper"
	"unit_testing/model/domain"
	"unit_testing/model/web"
	"unit_testing/repository"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	CreateUser(ctx echo.Context, request web.UserCreateRequest) (*domain.User, error)
	LoginUser(ctx echo.Context, request web.UserLoginRequest) (*domain.User, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate  *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate: validate,
	}
}

func (service *UserServiceImpl) CreateUser(ctx echo.Context, request web.UserCreateRequest) (*domain.User, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	// Check if the email already exists
	existingUser, _ := service.UserRepository.FindByEmail(request.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("Email already exists")
	}

	// Convert request to domain
	user := helper.UserCreateRequestToUserDomain(request)
	// Convert password to hash
	user.Password = helper.HashPassword(user.Password)

	err = service.UserRepository.Create(user)
	if err != nil {
		return nil, fmt.Errorf("Error when creating user: %s", err.Error())
	}

	return user, nil
}

func (service *UserServiceImpl) LoginUser(ctx echo.Context, request web.UserLoginRequest) (*domain.User, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	// Check if the email already exists
	existingUser, _ := service.UserRepository.FindByEmail(request.Email)
	if existingUser == nil {
		return nil, fmt.Errorf("Invalid email or password")
	}
	
	// Convert request to domain
	user := helper.UserLoginRequestToUserDomain(request)

	// Convert request to domain
	err = helper.ComparePassword(existingUser.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("Invalid email or password")
	}

	return existingUser, nil
}