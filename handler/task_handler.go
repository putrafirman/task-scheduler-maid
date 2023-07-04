package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type taskHandler struct {
	dbi *gorm.DB
}

type TaskHandler interface {
	CreateTask() echo.HandlerFunc
}

func NewTaskHandler(dbi *gorm.DB) TaskHandler {
	return &taskHandler{
		dbi: dbi,
	}
}

func (p *taskHandler) CreateTask() echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.JSON(http.StatusOK, "PONG")
	}
}
