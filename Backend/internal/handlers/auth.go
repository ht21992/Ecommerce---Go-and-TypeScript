// internal/handlers/auth.go
package handlers

import (
	"ecommerce/internal/auth"
	"ecommerce/internal/db"
	"ecommerce/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hash),
	}

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	token, _ := auth.GenerateJWT(user.ID)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(c *gin.Context) {

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	if err := db.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email/password"})
		return

	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email/password"})
		return
	}

	token, _ := auth.GenerateJWT(user.ID)
	c.JSON(http.StatusOK, gin.H{"token": token})

}
