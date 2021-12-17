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

	// Test Kafka. Use a goroutine to not block the main thread.
	go testKafka()

	err := rdb.Set(ctx, "visits", "0", 0).Err()
	if err != nil {
		panic(err)
	}

	// Set up routes etc.
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

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

	// Test Kafka. Use a goroutine to not block the main thread.
	go testKafka()

	val, err := rdb.Get(ctx, "visits").Result()
	if err != nil {
		panic(err)
	}

	return c.String(http.StatusOK, "current visit count: "+val)
}

func initialMigration() {
	host := os.Getenv("PGHOST")
	user := os.Getenv("PGUSER")
	database := os.Getenv("PGDATABASE")
	password := os.Getenv("PGPASSWORD")
	port := os.Getenv("PGPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, database, port)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

}
