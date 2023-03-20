package main

import (
	"log"
	"sync"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"tinderclone_back/src/pkg/database"
	"tinderclone_back/src/pkg/dto"
	"tinderclone_back/src/pkg/handlers"
	"tinderclone_back/src/pkg/middlewares"
	"tinderclone_back/src/pkg/services"
	"tinderclone_back/src/pkg/stores"
)

var (
	server     = "localhost"
	port       = "8000"
	apiVersion = "v1"
	dbUsername = "root"
	dbPassword = ""
	dbPort     = "5432"
	dbHost     = "localhost"
	dbName     = "tinder"
)

func launchServer(wg *sync.WaitGroup, ch chan string) {
	if serverInstance != nil {
		ch <- "Server instance already running"
		wg.Done()
	}
	serverInstance = echo.New()
	db := database.InitializeDb(&database.DbConfig{
		DbUsername: dbUsername,
		DbPassword: dbPassword,
		DbPort:     dbPort,
		DbHost:     dbHost,
		DbName:     dbName,
	})
	stores.InitializeStores(db)
	services.InitializeServices()
	initializeHandlers(serverInstance)
	provideInitData()
	serverInstance.Logger.Fatal(serverInstance.Start(server + ":" + port))
	wg.Done()
}

func initializeHandlers(si *echo.Echo) {
	healthCheckHandler := handlers.NewHealthCheckHandler()
	registerHandler := handlers.NewRegisterHandler(services.AccountMaker())
	loginHandler := handlers.NewLoginHandler(services.Authenticator())
	accountHandler := handlers.NewAccounterHandler(services.Accounter())
	countrierHandler := handlers.NewCountrierHandler(services.Countrier())
	citierHandler := handlers.NewCityHandler(services.Citier())
	permitterHandler := handlers.NewPermitterHandler(services.Permitter())

	adminGroup := serverInstance.Group("api/" + apiVersion + "/health")
	adminGroup.Use(echojwt.JWT([]byte("secret")))
	si.GET("api/"+apiVersion+"/health", healthCheckHandler.HandleHealthCheck, middlewares.AdminMiddleware)

	userGroup := serverInstance.Group("/api/" + apiVersion + "/user")
	userGroup.Use(echojwt.JWT([]byte("secret")))
	si.GET("api/"+apiVersion+"/user", accountHandler.GetAccountInformations, middlewares.LoggedUserMiddleware)

	si.POST("api/"+apiVersion+"/user", registerHandler.HandleRegister)
	si.POST("api/"+apiVersion+"/auth", loginHandler.HandleLogin)

	si.GET("api/"+apiVersion+"/country/all", countrierHandler.HandleGetAllCountriesNames)
	si.POST("api/"+apiVersion+"/country", countrierHandler.HandleSaveCountry, middlewares.AdminMiddleware)

	si.POST("api/"+apiVersion+"/city", citierHandler.HandleSaveNewCity, middlewares.AdminMiddleware)

	si.PATCH("api/"+apiVersion+"/permission", permitterHandler.HandleGiveUserAdminPermission, middlewares.AdminMiddleware)
}

func provideInitData() {
	services.Countrier().SaveNewCountry("ADMIN")
	log.Print("Country executed")

	services.Citier().SaveNewCity("ADMIN", "ADMIN")
	log.Print("City executed")

	services.AccountMaker().RegisterUser(dto.RegisterUser{
		Username: "ADMIN",
		Password: "ADMIN",
		Name:     "ADMIN",
		Surname:  "ADMIN",
		Countries: []string{
			"ADMIN",
		},
		Cities: []string{
			"ADMIN",
		},
	})
	log.Print("User executed")

	services.Permitter().GiveUserAdminPermission("ADMIN")
	log.Print("Permissions executed")
}
