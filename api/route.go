package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitAPIRoutes(_ *gin.Engine) {
	fmt.Println("API routes initialized")
}
