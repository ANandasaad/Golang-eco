package controller

import (
	"fmt"
	"golang-eco/helper"
	"golang-eco/models"
	"golang-eco/repositiories"
	"golang-eco/utilis"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf(" hashing password failed: %v", err)
	}

	return string(hashPassword), nil
}

// func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

// }
func SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()
		var user models.User

		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			utilis.RespondError(ctx, http.StatusBadRequest, err.Error())
			return
		}
		Validate := validator.New()
		err = Validate.Struct(user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		count, err := helper.GetUserByEmail(context, *user.Email)

		if err != nil {
			utilis.RespondError(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		if count > 0 {
			utilis.RespondError(ctx, http.StatusBadGateway, "User is already registered")
			return
		}

		if user.Password == nil {
			utilis.RespondError(ctx, http.StatusBadRequest, "Password is required")
			return
		}
		hashPassword, err := HashPassword(*user.Password)
		if err != nil {
			utilis.RespondError(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		user.Password = &hashPassword
		err = repositiories.CreateUser(&user)

		if err != nil {
			utilis.RespondError(ctx, http.StatusBadRequest, err.Error())
			return
		}
		utilis.RespondJSON(ctx, http.StatusCreated, user)

	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
