package hybrid_serverless

import (
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	"github.com/labstack/echo/v4"

	"putrafirman.com/playground/task-scheduler-maid/server"
)

func init() {
	// init for GCF
	functions.HTTP("dev-putra-logging", echoRouter)
}

func echoRouter(w http.ResponseWriter, r *http.Request) {
	// GCF Overrider
	InitRouter().ServeHTTP(w, r)
}

func InitRouter() *echo.Echo {
	// Echo instance
	e := echo.New()
	e = server.RegisterMiddleware(e)
	e = server.RegisterRouter(e)

	return e
}
