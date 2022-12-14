package controller

import (
	"encoding/json"
	"log"
	"strings"

	db "github.com/KrishKashiwala/go-crud-arch/database"
	"github.com/KrishKashiwala/go-crud-arch/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

// create model
type Customer struct {
	gorm.Model
	Name     string
	Password string
}

var (
	Name    string
	Country string
	Id      int
)

type Users struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Id      int    `json:"id"`
}

var DB *gorm.DB

// CREATE QUERY
func InsertUser(c *gin.Context) {
	db := db.GetDB()
	reqBody := json.NewDecoder(c.Request.Body)
	//read from request body
	var users models.User
	if err := reqBody.Decode(&users); err != nil {
		log.Println(users)
		log.Println("error occured while reading request body.")
	}

	//captilize all letters
	users.Name = strings.ToUpper(users.Name)
	users.Country = strings.ToUpper(users.Country)

	// check if user already exists
	var checkUser models.User
	checkResult := DB.Select("name", "password").Table("user").Where("name = ?", users.Name).Scan(&checkUser)
	if checkResult.RowsAffected != 0 {
		log.Println("User already exists")
	}

	//insert query in test db
	result := db.Select("name", "password").Table("user").Create(&users)
	log.Println(result)

}

// READ QUERY
func GetAllUsers(c *gin.Context) []Users {
	var users []Users
	// select query from test db
	db := db.GetDB()
	result := db.Select("name", "country").Table("person")
	log.Println(result)
	result.Scan(&users)
	return users

}

// UPDATE QUERY
func UpdateUser(c *gin.Context) {
	reqBody := json.NewDecoder(c.Request.Body)

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

	db := db.GetDB()
	reqBody := json.NewDecoder(c.Request.Body)

	//read from request body
	var users Users
	if err := reqBody.Decode(&users); err != nil {
		log.Fatal("error occured while reading request body.")
	}
	result := db.Select("name", "country").Table("person").Where("name = ?", users.Name).Delete(&users)
	log.Println(result)
}

// ADVANCED QUERIES

func Login(c *gin.Context) Users {
	user := SelectUser(c)
	if len(user) == 0 {
		log.Println("User not found")
	}
	return user[0]
}

// SELECT based on name
func SelectUser(c *gin.Context) []Users {
	reqBody := json.NewDecoder(c.Request.Body)
	db := db.GetDB()
	//read from request body
	var users Users
	if err := reqBody.Decode(&users); err != nil {
		log.Fatal("error occured while reading request body.")
	}

	var user []Users
	// select query from test db
	result := db.Select("name", "country").Table("person").Where("name = ?", users.Name)
	result.Scan(&user)
	return user

}
