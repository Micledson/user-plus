package routes

import (
	"github.com/labstack/echo/v4"
	"user-plus/endpoints/handlers"
)

func LoadRoutes() *echo.Echo {
	// Create a new Server
	router := echo.New()

	// Load Routes
	api := router.Group("/api")
	loadUserRoutes(api)
	return router
}

func loadUserRoutes(group *echo.Group) {
	userGroup := group.Group("/user/:email")

	user := handlers.NewUserHandler()
	userGroup.GET("/find", user.FindUserByEmail)
}
