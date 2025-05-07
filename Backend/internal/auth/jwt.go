// internal/auth/jwt.go

package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userID uint, isAdmin bool) (string, error) {

	secret := []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"user_id":  userID,
		"is_admin": isAdmin,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)

}
