package main

import (
	"log"
	"time"

	db "github.com/KrishKashiwala/go-crud-arch/database"
	"github.com/KrishKashiwala/go-crud-arch/routes"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading in .env file")
	}

}

type Repository struct {
	DB *gorm.DB
}

func main() {
	// db connection
	db.Setup()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// Set up CORS middleware options
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
	routes.RouteServer(router)

	router.Run()

}
