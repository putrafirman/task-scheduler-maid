package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	mysqlclient "putrafirman.com/playground/task-scheduler-maid/client/mysql"
	"putrafirman.com/playground/task-scheduler-maid/handler"
)

func RegisterRouter(e *echo.Echo) *echo.Echo {

	// DB Conn
	db := mysqlclient.ConnectDB()

	th := handler.NewTaskHandler(db)

	// Routes
	e.GET("/ping", ping)
	e.GET("/create-task", th.CreateTask())

	return e
}

// ping pong
func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong!")
}
