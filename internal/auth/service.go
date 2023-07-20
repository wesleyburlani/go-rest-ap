package auth

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/wesleyburlani/go-rest-api/internal/db"
	"github.com/wesleyburlani/go-rest-api/internal/users"
	"github.com/wesleyburlani/go-rest-api/pkg/crypto"
	custom_errors "github.com/wesleyburlani/go-rest-api/pkg/errors"
)

type IService interface {
	WithContext(ctx context.Context) IService
	Login(email string, password string) (*crypto.JwtToken, error)
	ValidateJwtToken(token string) (*crypto.JwtProps, error)
	GetUserFromJwtToken(token string) (*users.User, error)
}

type Service struct {
	db       *db.Database
	logger   *logrus.Logger
	ctx      context.Context
	jwtAuth  *crypto.JwtAuth
	usersSvc users.IService
}

func (s *Service) WithContext(ctx context.Context) IService {
	s.ctx = ctx
	return s
}

func NewService(
	db *db.Database,
	logger *logrus.Logger,
	jwtAuth *crypto.JwtAuth,
	usersSvc users.IService,
) *Service {
	return &Service{
		db:       db,
		logger:   logger,
		ctx:      context.Background(),
		jwtAuth:  jwtAuth,
		usersSvc: usersSvc,
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

	t, err := s.jwtAuth.GenerateToken(crypto.JwtProps{
		Username: user.Email,
	})

	if err != nil {
		return &crypto.JwtToken{}, custom_errors.NewUnknownError(err.Error())
	}
	return t, nil
}

func (s *Service) ValidateJwtToken(token string) (*crypto.JwtProps, error) {
	props, err := s.jwtAuth.DecodeToken(token)

	if err != nil {
		return &crypto.JwtProps{}, custom_errors.NewUnauthorizedError(err.Error())
	}

	return props, nil
}

func (s *Service) GetUserFromJwtToken(token string) (*users.User, error) {
	props, err := s.ValidateJwtToken(token)

	if err != nil {
		return &users.User{}, custom_errors.NewUnauthorizedError(err.Error())
	}

	user, err := s.usersSvc.GetByEmail(props.Username)

	if err != nil {
		return &users.User{}, custom_errors.NewUnauthorizedError(err.Error())
	}

	return &user, nil
}
