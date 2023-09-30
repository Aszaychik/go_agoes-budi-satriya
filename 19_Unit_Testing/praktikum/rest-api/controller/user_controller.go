package controller

import (
	"strconv"
	"strings"
	"unit_testing/helper"
	"unit_testing/model/web"
	"unit_testing/service"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	RegisterUserController(ctx echo.Context ) error
	LoginUserController(ctx echo.Context ) error
	GetUserController(ctx echo.Context) error
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func(c *UserControllerImpl) RegisterUserController(ctx echo.Context) error {
	userCreateRequest := web.UserCreateRequest{}
	err := ctx.Bind(&userCreateRequest) 
	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}

	response, err := c.UserService.CreateUser(ctx, userCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "Email already exists") {
			return helper.StatusEmailAlreadyExist(ctx, err)
		}

		return helper.StatusInternalServerError(ctx, err)
	}

	return helper.StatusCreated(ctx, "Success to create user", response)
}

func (c *UserControllerImpl) LoginUserController(ctx echo.Context) error {
	userLoginRequest := web.UserLoginRequest{}
	err := ctx.Bind(&userLoginRequest)
	if err != nil {
		return helper.StatusBadRequest(ctx, err)
	}

	response, err := c.UserService.LoginUser(ctx, userLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return helper.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "Invalid email or password") {
			return helper.StatusBadRequest(ctx, err)
		}

		return helper.StatusInternalServerError(ctx, err)
	}

	UserLoginResponse := helper.UserDomainToUserLoginResponse(response)

	return helper.StatusOK(ctx, "Success to login user", UserLoginResponse)
}

func (c *UserControllerImpl) GetUserController(ctx echo.Context) error {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return helper.StatusInternalServerError(ctx, err)
	}
	
	response, err := c.UserService.FindById(ctx, userIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "User not found") {
			return helper.StatusNotFound(ctx, err)
		}

		return helper.StatusInternalServerError(ctx, err)
	}

	return helper.StatusOK(ctx, "Success to get user", response)
}