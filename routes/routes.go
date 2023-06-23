package routes

import (
	"crud-database-postgresql/db"
	"crud-database-postgresql/handlers"
	"crud-database-postgresql/repository"
	"crud-database-postgresql/services"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo{
	e := echo.New()

	userRepository := repository.NewRepository(db.Init())
	userService := services.NewService(userRepository)
	userHandler := handlers.NewHandler(userService)

	e.GET("/", handlers.Root)
	e.GET("/user", userHandler.GetAll)
	e.GET("/user/:id", userHandler.FindById)
	e.POST("/createUser", userHandler.Create)
	e.PUT("/updateUser/:id", userHandler.Update)
	e.DELETE("/deleteUser/:id", userHandler.Delete)
	return e
}