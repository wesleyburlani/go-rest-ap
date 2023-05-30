package db

import (
	"github.com/sirupsen/logrus"
	"github.com/wesleyburlani/go-rest-api/models"
	"github.com/wesleyburlani/go-rest-api/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	gormLogger "gorm.io/gorm/logger"
)

func Init(config *utils.Config, logger *logrus.Logger) *gorm.DB {
	gormConfig := gorm.Config{}

	if config.Mode == utils.ReleaseMode {
		gormConfig.Logger = gormLogger.Default.LogMode(gormLogger.Silent)
	}

	db, err := gorm.Open(postgres.Open(config.DatabaseUrl), &gormConfig)

	if err != nil {
		logger.Fatalln(err)
	}

	err = db.AutoMigrate(&models.Album{})

	if err != nil {
		logger.Fatalln(err)
	}

	return db
}
