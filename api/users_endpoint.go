package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func createUsersRoutes(e *echo.Echo) {
	e.Post("/users", createUser)
	e.Get("/users", getUsers)
	e.Get("/users/:id", getUser)
}

// @Title createUser
// @Description create user
// @Accept  json
// @Param   user      body   User  true        "user object"
// @Success 201 {object}	User
// @Failure 400 {object} error "failed to create user"
// @Router /users [post]
func createUser(c *echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := validate.Struct(u); err != nil {
		return err
	}
	users.Users[u.ID] = *u
	return c.JSON(http.StatusCreated, u)
}

// @Title getUsers
// @Description get all users
// @Accept  json
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
