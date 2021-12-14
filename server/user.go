package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func allUsers(c echo.Context) error {
	var users []User
	db.Find(&users)
	fmt.Printf("All Users endpoint hit: %v", users)
	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	fmt.Printf("Retrieve user with id %s", id)
	var user User
	db.First(&user, id)
	return c.JSON(http.StatusOK, user)
}

func saveUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	user := User{Name: name, Email: email}
	db.Create(&user)

	fmt.Printf("Add new user with name: %s and email: %s", name, email)
	return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")
	email := c.FormValue("email")

	var user User
	db.First(&user, id)
	user.Name = name
	user.Email = email

	fmt.Printf("Update user with id: %s, set name to %s and email to %s", id, name, email)
	db.Save(&user)
	return c.JSON(http.StatusOK, user)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")

	db.Delete(&User{}, id)
	fmt.Printf("Deleted user with id: %s", id)
	return c.String(http.StatusOK, "ok")
}
