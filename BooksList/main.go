package main

import (
	"BooksList/src/Handler"
	"BooksList/src/Service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	//Logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("Logger initialized successfully")

	//Criando uma nova instância do Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Criando  serviço e o handler
	bookService := service.NewBookService()
	bookHandler := handler.NewBookHandler(bookService)

	//Rotas
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})
	e.GET("/health", handler.HealthCheck)

	//Rotas relacionadas aos livros
	books := e.Group("/books")
	books.POST("", bookHandler.AddBook)
	books.GET("/:id", bookHandler.GetBook)
	books.PUT("/", bookHandler.UpdateBook)
	books.DELETE("/:id", bookHandler.DeleteBook)

	// Iniciando o servidor na porta 8080
	e.Logger.Fatal(e.Start(":8080"))
}

