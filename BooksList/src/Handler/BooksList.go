package handler

import (
	models "BooksList/src/Model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Read struct {
	books []models.Book
}

//HealthCheck
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "UP",
	})
}

func AddBook(c echo.Context) error {
	var books models.Book
	if err := c.Bind(&books); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, books)
}

func GetBook(c echo.Context) error { 
	var books models.Book
	return c.JSON(http.StatusOK, books)
}

func AttBook(c echo.Context) error {
	var books models.Book
	return c.JSON(http.StatusOK, books)
}

func DeleteBook(c echo.Context) error{
	var books models.Book
	return c.JSON(http.StatusOK, books)
}