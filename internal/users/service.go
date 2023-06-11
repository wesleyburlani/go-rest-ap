package users

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/wesleyburlani/go-rest-api/pkg/errors"
	"gorm.io/gorm"
)

type IService interface {
	WithContext(ctx context.Context) IService
	Create(user User) (User, error)
	Get(id uint) (User, error)
	List(page int, limit int) ([]User, error)
	Update(id uint, user User) (User, error)
}

type Service struct {
	db     *gorm.DB
	logger *logrus.Logger
	ctx    context.Context
}

func NewService(db *gorm.DB, logger *logrus.Logger) *Service {
	return &Service{
		db:     db,
		logger: logger,
		ctx:    context.Background(),
	}
}

func (s *Service) WithContext(ctx context.Context) *Service {
	s.ctx = ctx
	return s
}

func (s *Service) Create(user User) (User, error) {
	tx := s.db.WithContext(s.ctx).Create(&user)

	if tx.Error != nil {
		return User{}, errors.NewUnknownErrorWithParent("error creating user on database", tx.Error)
	}

	s.logger.WithContext(s.ctx).WithFields(logrus.Fields{
		"id": user.ID,
	}).Debug("new user created")

	user.Password = ""
	return user, nil
}
