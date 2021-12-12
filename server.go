package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"

	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "redis-server:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

type Person struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email`
}

func main() {

	err := rdb.Set(ctx, "visits", "0", 0).Err()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/visits", getVisits)

	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/users", saveUser)
	e.POST("/save", save)

	e.POST("/employees", savePerson)

	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func saveUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func save(c echo.Context) error {
	name := c.FormValue("name")

	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return nil
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</br>")
}

func savePerson(c echo.Context) error {
	person := new(Person)
	if err := c.Bind(person); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, person)
}

func getVisits(c echo.Context) error {

	_, err := rdb.Incr(ctx, "visits").Result()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "visits").Result()
	if err != nil {
		panic(err)
	}

	return c.String(http.StatusOK, "visit count: "+val)
}
