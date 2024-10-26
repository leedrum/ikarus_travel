package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/leedrum/ikarus_travel/model"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LoadCurrentUser(ctx *gin.Context) model.User {
	user := ctx.MustGet("user").(model.User)

	return user
}
