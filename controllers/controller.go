package controller

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Name    string
	Country string
	Id      int
)

const (
	User     = "postgres"
	Password = "postgres"
	Port     = "5432"
	Host     = "localhost"
	Database = "test"
)

type Users struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Id      int    `json:"id"`
}

var DB *gorm.DB
var DSN string = "user=postgres password=postgres dbname=test port=5432 sslmode=disable "

// CREATE QUERY
func InsertUser(c *gin.Context) {
	reqBody := json.NewDecoder(c.Request.Body)

	//initialize db conn
	DB, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("error occured while DB connection.")
	}
	//read from request body
	var users Users
	if err := reqBody.Decode(&users); err != nil {
		log.Fatal("error occured while reading request body.")
	}

	//captilize all letters
	users.Name = strings.ToUpper(users.Name)
	users.Country = strings.ToUpper(users.Country)

	//insert query in test db
	result := DB.Select("name", "country").Table("person").Create(&users)
	log.Println(result)

}

// READ QUERY
func GetAllUsers(c *gin.Context) []Users {
	// initialize db conn
	DB, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("error occured while DB connection.")
	}

	var users []Users
	// select query from test db
	result := DB.Select("name", "country").Table("person")
	result.Scan(&users)
	return users

}

// UPDATE QUERY
func UpdateUser(c *gin.Context) {
	reqBody := json.NewDecoder(c.Request.Body)
	//initialize db conn
	DB, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("error occured while DB connection.")
	}
	//read from request body
	var users Users
	if err := reqBody.Decode(&users); err != nil {
		log.Fatal("error occured while reading request body.")
	}

	//captilize all letters
	users.Name = strings.ToUpper(users.Name)
	users.Country = strings.ToUpper(users.Country)

	//update query in test db
	result := DB.Select("name", "country").Table("person").Where("id = ?", users.Id).Updates(&users)
	log.Println(result)

}

// DELETE QUERY
func DeleteUser(c *gin.Context) {
	reqBody := json.NewDecoder(c.Request.Body)

	//initialize db conn
	DB, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("error occured while DB connection.")
	}
	//read from request body
	var users Users
	if err := reqBody.Decode(&users); err != nil {
		log.Fatal("error occured while reading request body.")
	}
	result := DB.Select("name", "country").Table("person").Where("name = ?", users.Name).Delete(&users)
	log.Println(result)
}
