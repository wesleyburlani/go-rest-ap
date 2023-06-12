package users

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	custom_errors "github.com/wesleyburlani/go-rest-api/pkg/errors"
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
	if tx := s.db.WithContext(s.ctx).First(&user); tx.Error == nil {
		return User{}, custom_errors.NewConflictError("user already exist")
	}

	tx := s.db.WithContext(s.ctx).Create(&user)

	if tx.Error != nil {
		return User{}, custom_errors.NewUnknownErrorWithParent("error creating user on database", tx.Error)
	}

	s.logger.WithContext(s.ctx).WithFields(logrus.Fields{
		"id": user.ID,
	}).Debug("new user created")

	user.Password = ""
	return user, nil
}

func (s *Service) Get(id uint) (User, error) {
	user := User{}
	tx := s.db.WithContext(s.ctx).First(&user)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return User{}, custom_errors.NewNotFoundErrorWithParent("user not found", tx.Error)
		}
		return User{}, custom_errors.NewUnknownErrorWithParent("error finding user on database", tx.Error)
	}
	return user, nil
}
