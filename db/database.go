package db

import (
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
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

	url, err := url.Parse(config.DatabaseUrl)

	if err != nil {
		logger.Fatalln(err)
	}

	dbName := strings.Trim(url.Path, "/")

	db, err := gorm.Open(postgres.Open(config.DatabaseUrl), &gormConfig)

	if err != nil {
		logger.Fatalln(err)
	}

	err = db.Use(otelgorm.NewPlugin(otelgorm.WithDBName(dbName)))

	if err != nil {
		logger.Fatalln(err)
	}

	err = db.AutoMigrate(&models.Album{})

	if err != nil {
		logger.Fatalln(err)
	}

	return db
}
