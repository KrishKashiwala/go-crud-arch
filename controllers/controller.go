package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	db "github.com/KrishKashiwala/go-crud-arch/database"
	"github.com/KrishKashiwala/go-crud-arch/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Id       int    `json:"id"`
}

// CREATE QUERY
func InsertUser(c *gin.Context) {
	db := db.GetDB()
	reqBody := json.NewDecoder(c.Request.Body)
	//read from request body
	var users models.User
	if err := reqBody.Decode(&users); err != nil {
		log.Println("error occured while reading request body.")
	}

	// check if user already exists
	var checkUser models.User
	checkResult := db.Select("name", "password").Table("users").Where("name = ?", users.Name).Scan(&checkUser)
	if checkResult.RowsAffected != 0 {
		c.JSON(403, gin.H{
			"msg": "user exists",
		})
		return
	}

	// hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(users.Password), 10)
	if err != nil {
		log.Fatal("password hash failed.")
	}

	newUser := models.User{
		Name:     users.Name,
		Password: string(hashedPassword),
	}

	//create new user
	result := db.Create(&newUser)

	if result.Error != nil {
		log.Fatal("user not created.!")
		return
	}

	//response
	c.JSON(200, gin.H{
		"msg": "user created successfully.",
	})

}

// READ QUERY
func GetAllUsers(c *gin.Context) []Users {
	var users []Users
	// select query from test db
	db := db.GetDB()
	result := db.Select("name", "password", "id").Table("users")
	log.Println(result)
	result.Scan(&users)
	return users

}

// UPDATE QUERY
func UpdateUser(c *gin.Context) *gorm.DB {
	db := db.GetDB()
	reqBody := json.NewDecoder(c.Request.Body)

	//read from request body
	var users Users
	if err := reqBody.Decode(&users); err != nil {
		log.Fatal("error occured while reading request body.")
	}

	// update only selected fields
	result := db.Table("users").Where("id = ?", users.Id).Updates(&users)
	return result

}

// DELETE QUERY
func DeleteUser(c *gin.Context) string {

	db := db.GetDB()
	reqBody := json.NewDecoder(c.Request.Body)

	//read from request body
	var users Users
	if err := reqBody.Decode(&users); err != nil {
		log.Fatal("error occured while reading request body.")
	}
	result := db.Table("users").Where("id = ?", users.Id).Delete(&users)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"msg": "can't deleted the user",
		})
	}
	return "successfully deleted " + strconv.Itoa(users.Id)
}

// ADVANCED QUERIES

func Login(c *gin.Context) {
	reqBody := json.NewDecoder(c.Request.Body)
	//read from request body
	var users Users
	if err := reqBody.Decode(&users); err != nil {
		c.JSON(404, gin.H{
			"msg": "error occured while reading request body.",
		})
		return
	}

	user, err := FindUser(users)
	if err != nil {
		return
	}
	log.Println("user received : ", user)
	// compare hash with password
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(users.Password))
	if passErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid password or email",
		})
		return
	}

	//Generate a JWT TOKEN
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//sign and get the complete encoded token as a string using secret key
	tokenString, tokenErr := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if tokenErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed token generation",
		})
	}

	// set cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// SELECT based on name
func FindUser(user Users) (Users, error) {
	db := db.GetDB()
	var checkUser Users
	// where query in users table
	result := db.First(&checkUser, "name = ?", user.Name)

	result.Scan(&checkUser)
	return checkUser, nil

}
