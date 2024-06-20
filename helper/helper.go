package helper

import (
	"golang-eco/database"
	"golang-eco/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/net/context"
)

func GetUserByEmail(ctx context.Context, email string) (int64, error) {
	var count int64

	err := database.DB.WithContext(ctx).Model(&models.User{}).Where("email = ?", email).Count(&count).Error

	return count, err

}

func GenerateToken(email string, firstName string, lastName string, userId uint) (string, string, error) {
	var secretKey = []byte("secret-key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email, "firstName": firstName, "lastName": lastName, "userId": userId, "expires": time.Now().Add(time.Hour * 24).Unix()})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email, "firstName": firstName, "lastName": lastName, "userId": userId, "expires": time.Now().Add(time.Hour * 24).Unix()})
	tokenString, err := token.SignedString(secretKey)
	refreshTokenString, err := refreshToken.SignedString(secretKey)

	if err != nil {
		return "", "", err
	}

	// Sign Refresh token

	if err != nil {
		return "", "", err
	}
	return tokenString, refreshTokenString, nil
}
