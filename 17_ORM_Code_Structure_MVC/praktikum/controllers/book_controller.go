package controllers

import (
	"net/http"
	"orm_code_structure/helpers"
	"orm_code_structure/models"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	model models.BooksModel
}

func (bc *BookController) InitBookController(bm models.BooksModel) {
	bc.model = bm
}

func (bc *BookController) GetBooks() echo.HandlerFunc {
	return func(c echo.Context) error {
		books  := bc.model.FindAll()

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success get all data", books))
	}
}

func (bc *BookController) GetBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		bookId := c.Param("id")
		book := bc.model.FindById(bookId)
		if book == nil {
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
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid user input", nil))
		}

		response := bc.model.Create(inputBook)
		if response == nil {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Failed create data something happen", nil))
		}

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success create data", response))
	}
}


func (bc *BookController) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		bookId := c.Param("id")
		input := models.Book{}

		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid user input", nil))
		}
		input.Id = bookId
		
		response := bc.model.Update(input)
		if response == nil {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Failed update data something happen", nil))
		}

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success update data", response))
	}
}

func (bc *BookController) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		bookId := c.Param("id")

		bc.model.Delete(bookId)
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success delete data", nil))
	}
}