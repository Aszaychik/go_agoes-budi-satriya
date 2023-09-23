package controllers

import (
	"intro_echo_golang/helper"
	"intro_echo_golang/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var Users []model.User

func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.SetResponse(http.StatusOK, "success get users", Users))
}

func GetUserController(c echo.Context) error {
	userId := c.Param("id")
	for _, user := range Users {
		if strconv.Itoa(user.Id) == userId {
			return c.JSON(http.StatusOK, helper.SetResponse(http.StatusOK, "success get user", user))
		}
	}
	return c.JSON(http.StatusNotFound, helper.SetResponse(http.StatusNotFound, "user not found", nil))
}

func DeleteUserController(c echo.Context) error {
	userId := c.Param("id")
	for _, user := range Users {
		if strconv.Itoa(user.Id) == userId {
			Users = append(Users[:user.Id-1], Users[user.Id:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success delete user",
				"user": user,
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

func CreateUserController(c echo.Context) error {
	user := model.User{}
	err := c.Bind(&user)
	helper.PanicIfError(err)

	if len(Users) == 0 {
		user.Id = 1
	} else {
		newId := Users[len(Users)-1].Id + 1
		user.Id = newId
	}

	Users = append(Users, user)
	return c.JSON(http.StatusOK, helper.SetResponse(http.StatusOK, "success create user", user))
}

func UpdateUserController(c echo.Context) error {
	userId := c.Param("id")
	newUpdateUser := model.User{}
	c.Bind(&newUpdateUser)

	for idx, user := range Users {
		if strconv.Itoa(user.Id) == userId {
			Users[idx] = newUpdateUser
			Users[idx].Id, _ = strconv.Atoi(userId)
			return c.JSON(http.StatusOK, helper.SetResponse(http.StatusOK, "success update user", Users[idx]))
		}
	}
	return c.JSON(http.StatusNotFound, helper.SetResponse(http.StatusNotFound, "user not found", nil))
}