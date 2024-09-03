package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"simple-template/controller"
	"simple-template/db"
)

func Init() error {
	e := echo.New()

	db, err := db.Init()
	if err != nil {
		return err
	}
	defer db.Close()

	e.GET("", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "Application is Running",
		})
	})

	user := e.Group("/user")

	user.GET("", controller.GetUserController(db))
	user.GET("/:id", controller.GetUserControllerById(db))
	user.POST("/create", controller.CreateUserController(db))
	user.PUT("/edit/:id", controller.UpdateUserController(db))
	user.DELETE("/delete/:id", controller.DeleteUserController(db))
	user.DELETE("", controller.BulkDeleteUserController(db))

	return e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
