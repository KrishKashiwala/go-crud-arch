package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	db "github.com/KrishKashiwala/go-crud-arch/database"
	"github.com/KrishKashiwala/go-crud-arch/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthValidate(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cookie with auth header not found",
		})
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(403)
		}

		// find user
		var user models.User
		db := db.GetDB()
		db.First(&user, claims["sub"])

		if user.Id == 0 {
			c.AbortWithStatus(403)
		}

		//set user
		c.Set("user", user)
	} else {
		c.JSON(403, gin.H{"error": "unauthorized"})
	}
	c.Next()
}
