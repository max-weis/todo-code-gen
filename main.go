package main

import (
	"fmt"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"strings"
	custommiddleware "todo-code-gen/internal/server/middleware"
	"todo-code-gen/internal/todo/boundary"
)

func main() {
	appContext := Initialize()

	e := echo.New()
	appContext.TodoRouter.RegisterHandlersWithBaseURL(e, "api")
	e.Logger = custommiddleware.ProvideZerologAdapter(zerolog.New(os.Stdout))

	configureSwaggerUI(e)

	e.GET("/health", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{"status": "UP"})
	})

	e.GET("/todos/spec", func(ctx echo.Context) error {
		swagger, _ := boundary.GetSwagger()
		return ctx.JSON(http.StatusOK, swagger)
	})

	shouldSkipMiddleware := func(context echo.Context) bool {
		return !strings.HasPrefix(context.Path(), fmt.Sprintf("/%s", "api"))
	}
	prometheus.NewPrometheus("echo", shouldSkipMiddleware).Use(e)
	e.Use(custommiddleware.AccessLogger(shouldSkipMiddleware))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", 8080)))
}

func configureSwaggerUI(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType},
	}))

	e.GET("/todos/swagger-ui", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/swagger-ui/index.html?spec=/todos/spec")
	})

	e.Static("/swagger-ui/", "./swagger-ui")
}
