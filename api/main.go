package main

import "gopkg.in/go-playground/validator.v8"

var (
	users    Users
	validate *validator.Validate
)

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

func main() {
	config := &validator.Config{TagName: "validate"}
	validate = validator.New(config)
	startServer()
}
