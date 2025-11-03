package controllers

import (
	"net/http"
	"time"
	"web-news-gin-gorm-sqlite-jwt/database"
	"web-news-gin-gorm-sqlite-jwt/middleware"
	"web-news-gin-gorm-sqlite-jwt/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var db, _ = database.SetupDB()

var jwtkey = []byte("secret-W0rd")

func Register(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var input models.User

	if err := db.Where("name = ?", user.Name).First(&input).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(hashedPassword)

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	if err := db.Where("name = ?", input.Name).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "неправильный логин или пароль"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "неправильный логин или пароль"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtkey)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to sign token"})
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}


