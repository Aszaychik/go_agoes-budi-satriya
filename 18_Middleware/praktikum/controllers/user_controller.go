package controllers

import (
	"golang_middleware/helpers"
	"golang_middleware/middleware"
	"golang_middleware/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	model models.UsersModel
}

func (uc *UserController) InitUserController(um models.UsersModel) {
	uc.model = um
}

func (uc *UserController) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uc.model.FindAll()
		if err != nil {
				logrus.Error(err)
				return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success get all data", users))
	}
}


func (uc *UserController) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("id")
		user, err := uc.model.FindById(userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
		}

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

			inputUser.Password = helpers.HashPassword(inputUser.Password)

			response, err := uc.model.Create(inputUser)
			if err != nil {
					logrus.Error(err)
					return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
			}

			token, err := middleware.GenerateJWTCookie(response.Name, c)
			if err != nil {
					return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
			}

			if token == "" {
					return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Token not found", nil))
			}

			responseMap := map[string]interface{}{
					"message": "Success create data",
					"user":    response,
					"token":   token,
			}

			return c.JSON(http.StatusOK, responseMap)
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
		
		response, err := uc.model.Update(input)
		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success update data", response))
	}
}

func (uc *UserController) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("id")

		err := uc.model.Delete(userId)
		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success delete data", nil))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		inputUser := models.LoginUser{}
		err := c.Bind(&inputUser)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid user input", nil))
		}

		response, err := uc.model.Login(inputUser.Email, inputUser.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Email or Password Invalid", nil))
		}

		token, err := middleware.GenerateJWTCookie(response.Name, c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
		}

		if token == "" {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Token not found", nil))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login Success",
		"name":    response.Name,
		"cookiesToken":   token,
	})
	}
}

func (uc *UserController) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.SetCookie(&http.Cookie{
				Name:     "token",
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
		})
		
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Logout Success", nil))
	}
}