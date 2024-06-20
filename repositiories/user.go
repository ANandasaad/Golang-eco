package repositiories

import (
	"errors"
	"golang-eco/database"
	"golang-eco/helper"
	"golang-eco/models"
	"math/rand"
	"time"
)

var (
	ErrUserIsAlreadyRegistered = errors.New("user is already registered")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func CreateUser(user *models.User) error {
	user.UserID = uint(rand.Uint32())
	token, refreshToken, err := helper.GenerateToken(*user.Email, *user.FirstName, *user.LastName, user.UserID)
	if err != nil {
		return err
	}
	user.Token = &token
	user.RefreshToken = &refreshToken
	user.UserCart = make([]models.ProductUser, 0)
	user.AddressDetails = make([]models.Address, 0)
	user.OrderStatus = make([]models.Order, 0)

	result := database.DB.Create(user)
	return result.Error
}

func AutoMigrate() {
	database.DB.AutoMigrate(&models.User{})
}
