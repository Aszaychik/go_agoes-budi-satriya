package controllers

import (
	"net/http"
	"orm_code_structure/helpers"
	"orm_code_structure/models"

	"github.com/labstack/echo/v4"
)

type BlogController struct {
	model models.BlogsModel
}

func (bc *BlogController) InitBlogController(bm models.BlogsModel) {
	bc.model = bm
}

func (bc *BlogController) GetBlogs() echo.HandlerFunc {
	return func(c echo.Context) error {
		blogs  := bc.model.FindAll()

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success get all data", blogs))
	}
}

func (bc *BlogController) GetBlog() echo.HandlerFunc {
	return func(c echo.Context) error {
		blogId := c.Param("id")
		blog := bc.model.FindById(blogId)
		if blog.Id == "" {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "Data not found", nil))
		}

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success get data", blog))
	}
}

func (bc *BlogController) CreateBlog() echo.HandlerFunc {
	return func(c echo.Context) error {
		inputBlog := models.Blog{}
		err := c.Bind(&inputBlog)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid user input", nil))
		}

		response := bc.model.Create(inputBlog)
		if response == nil {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Failed create data something happen", nil))
		}

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success create data", response))
	}
}


func (bc *BlogController) UpdateBlog() echo.HandlerFunc {
	return func(c echo.Context) error {
		blogId := c.Param("id")
		input := models.Blog{}

		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid user input", nil))
		}
		input.Id = blogId
		
		response := bc.model.Update(input)
		if response == nil {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Failed update data something happen", nil))
		}

		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success update data", response))
	}
}

func (bc *BlogController) DeleteBlog() echo.HandlerFunc {
	return func(c echo.Context) error {
		blogId := c.Param("id")

		bc.model.Delete(blogId)
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success delete data", nil))
	}
}