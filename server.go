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
	"strings"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/rs/cors"
	"github.com/thoas/stats"
)

type (
	User struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	Users struct {
		Users map[string]User `json:"users"`
	}
)

var (
	users Users
)

//----------
// Handlers
//----------

func welcome(c *echo.Context) error {
	return c.Render(http.StatusOK, "welcome", "Joe")
}

func createUser(c *echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	users.Users[u.ID] = *u
	return c.JSON(http.StatusCreated, u)
}

// @Title getUsers
// @Description get all users
// @Accept  json
// @Param   id     path    string     true        "User Id"
// @Success 200 {object} Users
// @Failure 400 {object} error "users not found"
// @Router /users [get]
func getUsers(c *echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

// @Title getUser
// @Description get user by Id
// @Accept  json
// @Param   id     path    string     true        "User Id"
// @Success 200 {object} User
// @Failure 400 {object} error "id required"
// @Failure 400 {object} error "user not found for id"
// @Router /users/{id} [get]
func getUser(c *echo.Context) error {
	return c.JSON(http.StatusOK, users.Users[c.P(0)])
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())

	//------------------------
	// Third-party middleware
	//------------------------

	// https://github.com/rs/cors
	e.Use(cors.Default().Handler)

	// https://github.com/thoas/stats
	s := stats.New()
	e.Use(s.Handler)
	// Route
	e.Get("/stats", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, s.Data())
	})

	//--------
	// Routes
	//--------

	e.Post("/users", createUser)
	e.Get("/users", getUsers)
	e.Get("/users/:id", getUser)

	//-------
	// Group
	//-------

	// Group with parent middleware
	a := e.Group("/admin")
	a.Use(func(c *echo.Context) error {
		// Security middleware
		return nil
	})
	a.Get("", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Welcome admin!")
	})

	for apiKey, _ := range apiDescriptionsJson {
		e.Get("/docs/"+apiKey, ApiDescriptionHandler)
	}

	// Start server
	e.Run(":1323")
}

func init() {
	u := Users{}
	u.Users = map[string]User{
		"1": User{
			ID:   "1",
			Name: "Wreck-It Ralph",
		},
	}
	users = u
}

func ApiDescriptionHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := strings.TrimLeft(r.RequestURI, "docs/")

	if json, ok := apiDescriptionsJson[apiKey]; ok {
		w.Write([]byte(json))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
