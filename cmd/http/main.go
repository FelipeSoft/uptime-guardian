package main

import (
	"database/sql"
	"log"
	"os"
	"github.com/FelipeSoft/uptime-guardian/internal/application/middleware"
	endpoint_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/endpoint"
	endpoint_handler "github.com/FelipeSoft/uptime-guardian/internal/infrastructure/handler/endpoint"
	host_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/host"
	host_handler "github.com/FelipeSoft/uptime-guardian/internal/infrastructure/handler/host"
	auth_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase"
	auth_handler "github.com/FelipeSoft/uptime-guardian/internal/infrastructure/handler"
	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/repository"
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
	userRepository := repository.NewUserRepositoryMySQL(db)
	endpointRepository := repository.NewEndpointRepositoryMySQL(db)
	hostRepository := repository.NewHostRepositoryMySQL(db)

	// Use Cases
	authUseCase := auth_usecase.NewAuthUseCase(userRepository)

	getAllEndpointUseCase := endpoint_usecase.NewGetAllEndpointUseCase(endpointRepository)
	getByIdEndpointUseCase := endpoint_usecase.NewGetByIdEndpointUseCase(endpointRepository)
	createEndpointUseCase := endpoint_usecase.NewCreateEndpointUseCase(endpointRepository)
	updateEndpointUseCase := endpoint_usecase.NewUpdateEndpointUseCase(endpointRepository)
	deleteEndpointUseCase := endpoint_usecase.NewDeleteEndpointUseCase(endpointRepository)

	getAllHostUseCase := host_usecase.NewGetAllHostUseCase(hostRepository)
	getByIdHostUseCase := host_usecase.NewGetByIdHostUseCase(hostRepository)
	createHostUseCase := host_usecase.NewCreateHostUseCase(hostRepository)
	updateHostUseCase := host_usecase.NewUpdateHostUseCase(hostRepository)
	deleteHostUseCase := host_usecase.NewDeleteHostUseCase(hostRepository)

	// Handlers
	authHandler := auth_handler.NewAuthHandler(authUseCase)

	getAllEndpointHandler := endpoint_handler.NewGetAllEndpointHandler(getAllEndpointUseCase)
	getByIdEndpointHandler := endpoint_handler.NewGetByIdEndpointHandler(getByIdEndpointUseCase)
	createEndpointHandler := endpoint_handler.NewCreateEndpointHandler(createEndpointUseCase)
	updateEndpointHandler := endpoint_handler.NewUpdateEndpointHandler(updateEndpointUseCase)
	deleteEndpointHandler := endpoint_handler.NewDeleteEndpointHandler(deleteEndpointUseCase)

	getAllHostHandler := host_handler.NewGetAllHostHandler(getAllHostUseCase)
	getByIdHostHandler := host_handler.NewGetByIdHostHandler(getByIdHostUseCase)
	createHostHandler := host_handler.NewCreateHostHandler(createHostUseCase)
	updateHostHandler := host_handler.NewUpdateHostHandler(updateHostUseCase)
	deleteHostHandler := host_handler.NewDeleteHostHandler(deleteHostUseCase)

	// Middlewares
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := middleware.ValidateRequestBodyDynamic(c); err != nil {
				return err
			}
			return next(c)
		}
	})

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := middleware.VerifyUserAuthentication(c); err != nil {
				return err
			}
			return next(c)
		}
	})

	// Routes
	e.POST("/auth/login", authHandler.LoginUser)

	e.POST("/endpoint", createEndpointHandler.Execute)
	e.PUT("/endpoint/:id", updateEndpointHandler.Execute)
	e.DELETE("/endpoint/:id", deleteEndpointHandler.Execute)
	e.GET("/endpoint", getAllEndpointHandler.Execute)
	e.GET("/endpoint/:id", getByIdEndpointHandler.Execute)

	e.POST("/host", createHostHandler.Execute)
	e.PUT("/host/:id", updateHostHandler.Execute)
	e.DELETE("/host/:id", deleteHostHandler.Execute)
	e.GET("/host", getAllHostHandler.Execute)
	e.GET("/host/:id", getByIdHostHandler.Execute)

	e.Logger.Fatal(e.Start(os.Getenv("HTTP_SERVER")))
}
