package controllers

import (
	"net/http"
	"orm_code_structure/helpers"
	"orm_code_structure/models"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	model models.UsersModel
}

func (uc *UserController) InitUserController(um models.UsersModel) {
	uc.model = um
}

func (uc *UserController) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users  := uc.model.FindAll()

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success get all data", users))
	}
}

func (uc *UserController) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("id")
		user := uc.model.FindById(userId)
		if user.Id == "" {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "Data not found", nil))
		}

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success get data", user))
	}
}

func (uc *UserController) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		inputUser := models.User{}
		err := c.Bind(&inputUser)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid user input", nil))
		}

		response := uc.model.Create(inputUser)
		if response == nil {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Failed create data something happen", nil))
		}

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success create data", response))
	}
}


func (uc *UserController) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("id")
		input := models.User{}

		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid user input", nil))
		}
		input.Id = userId
		
		response := uc.model.Update(input)
		if response == nil {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Failed update data something happen", nil))
		}

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success update data", response))
	}
}

func (uc *UserController) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("id")

		uc.model.Delete(userId)
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success delete data", nil))
	}
}