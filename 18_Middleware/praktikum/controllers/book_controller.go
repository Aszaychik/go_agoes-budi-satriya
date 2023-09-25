package controllers

import (
	"golang_middleware/helpers"
	"golang_middleware/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type BookController struct {
	model models.BooksModel
}

func (bc *BookController) InitBookController(bm models.BooksModel) {
	bc.model = bm
}

func (bc *BookController) GetBooks() echo.HandlerFunc {
	return func(c echo.Context) error {
		books, err := bc.model.FindAllBooks()
		if err != nil {
				logrus.Error(err)
				return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success get all data", books))
	}
}


func (bc *BookController) GetBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		bookId := c.Param("id")
		book, err := bc.model.FindById(bookId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
		}

		if book.Id == "" {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "Data not found", nil))
		}

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success get data", book))
	}
}

func (bc *BookController) CreateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
			inputBook := models.Book{}
			err := c.Bind(&inputBook)
			if err != nil {
					return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid book input", nil))
			}

			response, err := bc.model.Create(inputBook)
			if err != nil {
					logrus.Error(err)
					return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
			}

			return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success create data", response))
	}
}


func (uc *BookController) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		bookId := c.Param("id")
		input := models.Book{}

		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid book input", nil))
		}
		input.Id = bookId
		
		response, err := uc.model.Update(input)
		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success update data", response))
	}
}

func (uc *BookController) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		bookId := c.Param("id")

		err := uc.model.Delete(bookId)
		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success delete data", nil))
	}
}