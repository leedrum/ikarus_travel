package internal

import (
	"github.com/leedrum/ikarus_travel/model"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(config Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DBSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to database")
	}
	db.AutoMigrate(&model.User{}, &model.Tour{}, &model.Hotel{}, &model.Reservation{})

	return db
}
