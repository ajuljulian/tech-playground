package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"context"

	"github.com/go-redis/redis/v8"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Redis
var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	Password: "", // no password set
	DB:       0,  // use default DB
})

// Postgres
var db *gorm.DB

type User struct {
	gorm.Model
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email`
}

func main() {

	// Set up the database
	initialMigration()

	err := rdb.Set(ctx, "visits", "0", 0).Err()
	if err != nil {
		panic(err)
	}

	// Set up routes etc.
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/visits", getVisits)

	e.GET("/users", allUsers)
	e.GET("/users/:id", getUser)
	e.POST("/users", saveUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
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

	return c.String(http.StatusOK, "current visit count: "+val)
}

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

func initialMigration() {
	host := os.Getenv("PGHOST")
	user := os.Getenv("PGUSER")
	database := os.Getenv("PGDATABASE")
	password := os.Getenv("PGPASSWORD")
	port := os.Getenv("PGPORT")

	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, database, port)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to database")
	}
	fmt.Printf("db: %v", db)

	// Migrate the schema
	db.AutoMigrate(&User{})

}

/*
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
*/
