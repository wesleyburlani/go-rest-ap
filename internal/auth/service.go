package auth

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/wesleyburlani/go-rest-api/internal/db"
	"github.com/wesleyburlani/go-rest-api/pkg/crypto"
	custom_errors "github.com/wesleyburlani/go-rest-api/pkg/errors"
)

type IService interface {
	WithContext(ctx context.Context) IService
	Login(email string, password string) (*crypto.JwtToken, error)
}

type Service struct {
	db     *db.Database
	logger *logrus.Logger
	ctx    context.Context
	auth   *crypto.JwtAuth
}

func (s *Service) WithContext(ctx context.Context) *Service {
	s.ctx = ctx
	return s
}

func NewService(db *db.Database, logger *logrus.Logger, auth *crypto.JwtAuth) *Service {
	return &Service{
		db:     db,
		logger: logger,
		ctx:    context.Background(),
		auth:   auth,
	}
}

func (s *Service) Login(email string, password string) (*crypto.JwtToken, error) {
	user, err := s.db.Queries.GetUserByEmail(s.ctx, email)

	if err != nil {
		return &crypto.JwtToken{}, custom_errors.NewNotFoundError("user not found")
	}

	if !crypto.CompareHashAndPassword(user.Password, password) {
		return &crypto.JwtToken{}, custom_errors.NewUnauthorizedError("invalid password")
	}

	t, err := s.auth.GenerateToken(crypto.JwtProps{
		Username: user.Email,
	})

	if err != nil {
		return &crypto.JwtToken{}, custom_errors.NewUnknownError(err.Error())
	}
	return t, nil
}
