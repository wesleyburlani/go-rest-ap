package database

import (
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"github.com/wesleyburlani/go-rest-api/internal/config"
	"github.com/wesleyburlani/go-rest-api/internal/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	gormLogger "gorm.io/gorm/logger"
)

func Init(c *config.Config, logger *logrus.Logger) *gorm.DB {
	gormConfig := gorm.Config{}

	if c.Mode == config.ReleaseMode {
		gormConfig.Logger = gormLogger.Default.LogMode(gormLogger.Silent)
	}

	url, err := url.Parse(c.DatabaseUrl)

	if err != nil {
		logger.Fatalln(err)
	}

	dbName := strings.Trim(url.Path, "/")

	db, err := gorm.Open(postgres.Open(c.DatabaseUrl), &gormConfig)

	if err != nil {
		logger.Fatalln(err)
	}

	err = db.Use(otelgorm.NewPlugin(otelgorm.WithDBName(dbName)))

	if err != nil {
		logger.Fatalln(err)
	}

	// @TODO: add migrations
	err = db.AutoMigrate(&users.User{})

	if err != nil {
		logger.Fatalln(err)
	}

	return db
}
