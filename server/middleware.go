package server

import (
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type CustomValidator struct {
	validator *validator.Validate
	e         *echo.Echo
}

func RegisterMiddleware(e *echo.Echo) *echo.Echo {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set logger level
	e.Logger.SetLevel(log.DEBUG)

	// Custom validator
	e.Validator = &CustomValidator{validator: validator.New(), e: e}

	// Single API Key
	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "query:api-key",
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == os.Getenv("X_API_KEY"), nil
		},
	}))

	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.New().String()
		},
	}))

	return e
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Get Logger from echo context and release after done
		c := cv.e.AcquireContext()
		c.Logger().Error(err)
		defer cv.e.ReleaseContext(c)

		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
