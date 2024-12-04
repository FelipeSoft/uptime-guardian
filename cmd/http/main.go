package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/FelipeSoft/uptime-guardian/internal/http/application/middleware"
	"github.com/FelipeSoft/uptime-guardian/internal/http/application/usecase"
	"github.com/FelipeSoft/uptime-guardian/internal/http/infrastructure/handler"
	"github.com/FelipeSoft/uptime-guardian/internal/http/infrastructure/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Environment Variable Error: %s", err.Error())
	}

	e := echo.New()
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatalf("MySQL Connection Error: %s", err.Error())
	}

	// Repository
	endpointRepository := repository.NewEndpointRepositoryMySQL(db)

	// Use Cases
	getAllEndpointUseCase := usecase.NewGetAllEndpointUseCase(endpointRepository)
	getByIdEndpointUseCase := usecase.NewGetByIdEndpointUseCase(endpointRepository)
	createEndpointUseCase := usecase.NewCreateEndpointUseCase(endpointRepository)
	updateEndpointUseCase := usecase.NewUpdateEndpointUseCase(endpointRepository)
	deleteEndpointUseCase := usecase.NewDeleteEndpointUseCase(endpointRepository)

	// Handlers
	getAllEndpointHandler := handler.NewGetAllEndpointHandler(getAllEndpointUseCase)
	getByIdEndpointHandler := handler.NewGetByIdEndpointHandler(getByIdEndpointUseCase)
	createEndpointHandler := handler.NewCreateEndpointHandler(createEndpointUseCase)
	updateEndpointHandler := handler.NewUpdateEndpointHandler(updateEndpointUseCase)
	deleteEndpointHandler := handler.NewDeleteEndpointHandler(deleteEndpointUseCase)

	// Middlewares
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := middleware.ValidateRequestBodyDynamic(c); err != nil {
				return err
			}

			return next(c)
		}
	})

	// Routes
	e.POST("/endpoint", createEndpointHandler.Execute)
	e.PUT("/endpoint/:id", updateEndpointHandler.Execute)
	e.DELETE("/endpoint/:id", deleteEndpointHandler.Execute)
	e.GET("/endpoint", getAllEndpointHandler.Execute)
	e.GET("/endpoint/:id", getByIdEndpointHandler.Execute)

	e.Logger.Fatal(e.Start(os.Getenv("HTTP_SERVER")))
}
