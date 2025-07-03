package routes

import (
	"github.com/PorcoGalliard/rumahweb-interview/cmd/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler *handler.UserHandler, JWTSecret string) {
	router.POST("/v1/login", userHandler.Login)
	router.POST("/v1/register", userHandler.Register)

	// private := router.Group("/auth")
	// router.POST("/v1/users", userHandler.Register)
	// private.Use(middleware.AuthMiddleware(JWTSecret))
	// private.GET("/v1/users/:id", userHandler.GetUserByID)
	// private.GET("/v1/users", userHandler.GetAllUser)
	// private.PUT("/v1/users/:id", userHandler.UpdateUser)
	// private.DELETE("/v1/users/:id", userHandler.DeleteUser)

	router.GET("/v1/users/:id", userHandler.GetUserByID)
	router.GET("/v1/users", userHandler.GetAllUser)
	router.PUT("/v1/users/:id", userHandler.UpdateUser)
	router.DELETE("/v1/users/:id", userHandler.DeleteUser)
} 