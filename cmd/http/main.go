package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
    "github.com/rs/cors"
	middleware "github.com/FelipeSoft/uptime-guardian/internal/application/middleware"
	auth_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase"
	endpoint_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/endpoint"
	host_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/host"
	auth_handler "github.com/FelipeSoft/uptime-guardian/internal/infrastructure/handler"
	endpoint_handler "github.com/FelipeSoft/uptime-guardian/internal/infrastructure/handler/endpoint"
	host_handler "github.com/FelipeSoft/uptime-guardian/internal/infrastructure/handler/host"
	_ "github.com/go-sql-driver/mysql"
	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/repository"
	"github.com/FelipeSoft/uptime-guardian/internal/infrastructure/adapter"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Environment Variable Error: %s", err.Error())
	}
	
	r := http.NewServeMux()
	httpServer := os.Getenv("HTTP_SERVER")

	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatalf("MySQL Connection Error: %s", err.Error())
	}

	// Adapters
	bcryptHashAdapter := adapter.NewBcryptHashAdapter()
	jwtAdapter := adapter.NewJwtAdapter()

	// Repository
	userRepository := repository.NewUserRepositoryMySQL(db)
	endpointRepository := repository.NewEndpointRepositoryMySQL(db)
	hostRepository := repository.NewHostRepositoryMySQL(db)

	// Use Cases
	authUseCase := auth_usecase.NewAuthUseCase(userRepository, bcryptHashAdapter)

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
	authHandler := auth_handler.NewAuthHandler(authUseCase, jwtAdapter)

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
	authMiddleware := middleware.NewAuthMiddleware(jwtAdapter)
	handler := cors.New(cors.Options{
		// AllowedOrigins: []string{"http://localhost:3000"},
		AllowedOrigins: []string{"*"},
	}).Handler(r)

	// Routes
	r.HandleFunc("/auth/login", middleware.Limit(authHandler.LoginUser))

	r.HandleFunc("/endpoint/create", createEndpointHandler.Execute)
	r.HandleFunc("/endpoint/update/{id}", updateEndpointHandler.Execute)
	r.HandleFunc("/endpoint/delete/{id}", deleteEndpointHandler.Execute)

	r.HandleFunc("/endpoint", authMiddleware.RequireAuthentication(getAllEndpointHandler.Execute))
	r.HandleFunc("/endpoint/{id}", getByIdEndpointHandler.Execute)

	r.HandleFunc("/host/create", createHostHandler.Execute)
	r.HandleFunc("/host/update/{id}", updateHostHandler.Execute)
	r.HandleFunc("/host/delete/{id}", deleteHostHandler.Execute)
	r.HandleFunc("/host", getAllHostHandler.Execute)
	r.HandleFunc("/host/{id}", getByIdHostHandler.Execute)

	fmt.Printf("HTTP Server listening on %s", httpServer)
	http.ListenAndServe(httpServer, handler)
}
