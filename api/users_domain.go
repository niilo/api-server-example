package main

type User struct {
	ID   string `json:"id,required" validate:"required" description:"user ID"`
	Name string `json:"name,required" validate:"required" description:"Firstname and lastname"`
}

type Users struct {
	Users map[string]User `json:"users"`
}
