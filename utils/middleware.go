package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jatraMaya/go-library/config"
	"github.com/jatraMaya/go-library/models"
)

func RequiredAuth(c *gin.Context) {
	tokenString, _ := c.Cookie("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.AppConfig.JWT.SecretKey), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Message": "User unauthorized",
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Message": "User Unauthorized",
			})
			return
		}

		var user models.User
		models.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Message": "User Unauthorized",
			})
			return
		}

		c.Set("user", user)
		c.Next()
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Message": "User Unauthorized",
		})
		return
	}
}
