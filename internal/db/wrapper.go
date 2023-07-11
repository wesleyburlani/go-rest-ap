package db

import (
	"github.com/XSAM/otelsql"
	"github.com/sirupsen/logrus"
	"github.com/wesleyburlani/go-rest-api/internal/config"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type Database struct {
	Queries *Queries
	logger  *logrus.Logger
}

func NewDatabase(c *config.Config, logger *logrus.Logger) *Database {
	db, err := otelsql.Open("postgres", c.DatabaseUrl, otelsql.WithAttributes(
		semconv.DBSystemPostgreSQL,
	))
	if err != nil {
		logger.Fatal(err)
	}

	queries := New(db)
	return &Database{
		Queries: queries,
		logger:  logger,
	}
}
