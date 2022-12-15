package routes

import (
	controller "github.com/KrishKashiwala/go-crud-arch/controllers"
	"github.com/KrishKashiwala/go-crud-arch/middlewares"
	"github.com/gin-gonic/gin"
)

func RouteServer(router *gin.Engine) {

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
		publicRoutes.POST("/login", func(c *gin.Context) {
			controller.Login(c)
		})

		//update routes
		publicRoutes.PUT("/update", func(c *gin.Context) {
			controller.UpdateUser(c)
		})

		//delete routes
		publicRoutes.DELETE("/delete", func(c *gin.Context) {
			result := controller.DeleteUser(c)
			c.JSON(200, gin.H{
				"msg": result,
			})
		})
	}

	//private routes
	privateRoutes := router.Group("/api/private", middlewares.AuthValidate)
	{
		privateRoutes.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"secret_key_for_nuclear_code": "Hello World",
			})
		})

	}
}
