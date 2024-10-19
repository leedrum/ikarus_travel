package internal

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Router *gin.Engine
	Config Config
	DB     *gorm.DB
}
