package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jatraMaya/go-library/config"
	"github.com/jatraMaya/go-library/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// @Summary Create a new user
// @Description Creates a new user with the provided name, email, and password
// @ID create-user
// @Accept json
// @Produce json
// @Param user body SignUpRequest true "User information to be created"
// @Success 201 {object} gin.H{"Message": "Signup success"}
// @Failure 400 {object} gin.H{"Message": "Failed to read request body"}
// @Failure 400 {object} gin.H{"Message": "Failed to hash password"}
// @Failure 400 {object} gin.H{"Message": "Error when signup"}
// @Router /signup [post]
func SignUp(c *gin.Context) {

	var UserBody struct {
		Name     string `json:"Name" binding:"required"`
		Email    string `json:"Email" binding:"required,email"`
		Password string `json:"Password" binding:"required"`
	}

	if c.ShouldBindJSON(&UserBody) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Failed to read request body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(UserBody.Password), bcrypt.DefaultCost)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Failed to hash password",
		})
		return
	}

	user := models.User{Name: UserBody.Name, Email: UserBody.Email, Password: string(hash)}
	result := models.DB.Create(&user)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Error when signup",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Signup success",
	})

}

// @Summary Authenticate user based on email and password
// @Description Authenticates a user based on email and password, returns a JWT token on successful login
// @ID authenticate-user
// @Accept json
// @Produce json
// @Param user body LoginRequest true "User information for login"
// @Success 200 {object} gin.H{"Message": "Login success"}
// @Failure 400 {object} gin.H{"Message": "Required email and password to login"}
// @Failure 400 {object} gin.H{"Message": "Invalid email or password"}
// @Failure 500 {object} gin.H{"Message": "Internal Server Error", "Error": "error details"}
// @Router /login [post]
func Login(c *gin.Context) {
	var requestBody struct {
		Email    string `json:"Email" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Required email and password to login",
			"Error":   err.Error(),
		})
		return
	}

	var user models.User
	if err := models.DB.Where("email = ?", requestBody.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Message": "Invalid email or password",
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Message": "Internal Server Error",
			"Error":   err.Error(),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.AppConfig.JWT.SecretKey))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Message": "Failed to create jwt token",
			"Error":   err.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"Message": "Login success",
	})
}

// @Summary Logout user and delete cookie
// @Description Logs out the user by deleting the authentication cookie
// @ID logout-user
// @Produce json
// @Success 200 {object} gin.H{"Message": "Logout successful"}
// @Router /logout [post]
func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"Message": "Logout successful",
	})
}
