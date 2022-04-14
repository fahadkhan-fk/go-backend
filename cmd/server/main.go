package main

import (
	"fmt"

	"github.com/fahadkhan-fk/go-backend/handler/user"
	"github.com/fahadkhan-fk/go-backend/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello Fahad !!")

	// Setup Database connection
	postgres.EstablishDBConnection()

	// Gin framework routes
	route := gin.Default()

	// route path eg:
	route.POST("/user", user.Create)
	route.GET("/users/", user.GetAll)
	route.GET("/user/:id", user.GetByID)   // path variable `id`
	route.PUT("/user/:id", user.Update)    // path variable `id`
	route.DELETE("/user/:id", user.Delete) // path variable `id`

	// listening at :8080
	route.Run(":8080")

}
