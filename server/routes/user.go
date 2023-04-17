package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	UserRepositroy := repositories.RepositoryUser(mysql.DB)

	h := handlers.HandlerUser(UserRepositroy)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:email", h.GetUserByEmail)
	e.DELETE("/user/:id", h.DeleteUser)
	e.PATCH("/user-profile/:id", middleware.Auth(middleware.UploadFile(h.UpdateUserPicture)))
}
