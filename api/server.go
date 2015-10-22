// @APIVersion 1.0.0
// @BasePath /
// @APITitle Api Server Example
// @APIDescription Api Server Example
// @Contact <email>
// @TermsOfServiceUrl https://github.com/niilo/api-server-example
// @License MIT
// @LicenseUrl https://github.com/niilo/api-server-example/LICENSE.md
package main

import (
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/rs/cors"
	"github.com/thoas/stats"
)

func startServer() {

	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())

	// https://github.com/rs/cors
	e.Use(cors.Default().Handler)

	// https://github.com/thoas/stats
	s := stats.New()
	e.Use(s.Handler)
	// Route
	e.Get("/stats", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, s.Data())
	})

	createUsersRoutes(e)
	createDocsRoute(e)

	// Group with parent middleware
	a := e.Group("/admin")
	a.Use(func(c *echo.Context) error {
		// Security middleware
		return nil
	})
	a.Get("", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Welcome admin!")
	})

	// Start server
	e.Run(":1323")
}
