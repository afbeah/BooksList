package main

import (
	handlers "BooksList/src/Handler"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {

	//Logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("Logger initialized successfully", zap.String("ex", "logger simple ex"))

	// Criando uma nova instância do Echo
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Definindo uma rota básica
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	//Rota para o Healthcheck
	e.GET("/health", handlers.HealthCheck)

	// Iniciando o servidor na porta 8080
	e.Logger.Fatal(e.Start(":8080"))

	//Rotas
	e.GET("books/getbook", handlers.GetBook)
	e.POST("books/addbook", handlers.AddBook)
	e.PUT("books/attbook", handlers.AttBook)
	e.DELETE("books/deletebook", handlers.DeleteBook)



}
