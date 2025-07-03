package main

import (
	"github.com/PorcoGalliard/rumahweb-interview/cmd/config"
	"github.com/PorcoGalliard/rumahweb-interview/cmd/handler"
	"github.com/PorcoGalliard/rumahweb-interview/cmd/repository"
	"github.com/PorcoGalliard/rumahweb-interview/cmd/routes"
	"github.com/PorcoGalliard/rumahweb-interview/cmd/services"
	pkgConfig "github.com/PorcoGalliard/rumahweb-interview/pkg/config"
	"github.com/PorcoGalliard/rumahweb-interview/resource"
	"github.com/gin-gonic/gin"
)

func main() {
	jwt := "jwtcihuy"
	config := pkgConfig.LoadConfig(&config.UserConfig{},
		pkgConfig.WithConfigPath("files"),
		pkgConfig.WithConfigFile("user_config"),
		pkgConfig.WithConfigType("yaml"),
	)

	postgres := resource.InitPostgres(config.Database)
	router := gin.Default()

	userRepository := repository.NewUserRepository(postgres)
	userService := services.NewUserServices(userRepository, jwt)
	userHandler := handler.NewUserHandler(userService)

	routes.SetupRoutes(router, userHandler, jwt)
	router.Run("")
}