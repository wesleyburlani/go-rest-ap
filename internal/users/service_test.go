package users_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/wesleyburlani/go-rest-api/internal/config"
	"github.com/wesleyburlani/go-rest-api/internal/database"
	"github.com/wesleyburlani/go-rest-api/internal/users"
)

func setupServiceSuite() *users.Service {
	c := config.LoadConfig()
	l := logrus.New()
	l.Out = io.Discard
	db := database.Init(c, l)
	svc := users.NewService(db, l)
	return svc
}

func TestSomething(t *testing.T) {
	svc := setupServiceSuite()
	u, e := svc.Get(1)
	fmt.Printf("%v, %v\n", u, e)
	assert.Nil(t, e)
	assert.Nil(t, u)
}
