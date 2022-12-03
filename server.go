package main

import (
	"log"

	controller "github.com/KrishKashiwala/go-crud-arch/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading in .env file")
	}

}

func main() {

	//print hello server
	println("Hello Server")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	group := router.Group("/api")

	publicRoutes := group.Group("/public")
	{

		//get routes
		publicRoutes.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"data": controller.GetAllUsers(c),
			})

		})

		//post routes
		publicRoutes.POST("/create", func(c *gin.Context) {
			controller.InsertUser(c)
		})

		//update routes
		publicRoutes.PUT("/update", func(c *gin.Context) {
			controller.UpdateUser(c)
		})

		//delete routes
		publicRoutes.DELETE("/delete", func(c *gin.Context) {
			controller.DeleteUser(c)
		})
	}
	router.Run()

}
