package users

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	custom_errors "github.com/wesleyburlani/go-rest-api/pkg/errors"
	"gorm.io/gorm"
)

type IService interface {
	WithContext(ctx context.Context) IService
	Create(user CreateUserProps) (User, error)
	Get(id uint) (User, error)
	List(page int, limit int) ([]User, error)
	Update(id uint, user UpdateUserProps) (User, error)
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

func ValidateValuer(field reflect.Value) interface{} {

	if valuer, ok := field.Interface().(driver.Valuer); ok {

		val, err := valuer.Value()
		if err == nil {
			return val
		}
		// handle the error how you want
	}

	return nil
}

func (s *Service) Create(user CreateUserProps) (User, error) {
	validate := validator.New()
	validate.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})
	if err := validate.Struct(&user); err != nil {
		return User{}, custom_errors.NewValidationErrorWithParent("field validation error", err)
	}

	if tx := s.db.WithContext(s.ctx).First(&User{Email: user.Email}); tx.Error == nil {
		return User{}, custom_errors.NewConflictError("user already exist")
	}

	registry := User{
		Email:    user.Email,
		Password: user.Password,
	}
	tx := s.db.WithContext(s.ctx).Create(&registry)

	if tx.Error != nil {
		return User{}, custom_errors.NewUnknownErrorWithParent("error creating user on database", tx.Error)
	}

	s.logger.WithContext(s.ctx).WithFields(logrus.Fields{
		"id": registry.ID,
	}).Debug("new user created")

	registry.Password = sql.NullString{String: ""}
	return registry, nil
}

func (s *Service) Update(id uint, user UpdateUserProps) (User, error) {
	if err := validator.New().Struct(&user); err != nil {
		return User{}, custom_errors.NewValidationErrorWithParent("field validation error", err)
	}

	currentUser := User{
		ID: id,
	}
	if tx := s.db.WithContext(s.ctx).First(&currentUser); tx.Error != nil {
		return User{}, custom_errors.NewNotFoundErrorWithParent("user not found", tx.Error)
	}

	registry := User{
		ID:       id,
		Email:    user.Email,
		Password: user.Password,
	}

	tx := s.db.WithContext(s.ctx).Save(&registry)

	if tx.Error != nil {
		return User{}, custom_errors.NewUnknownErrorWithParent("error creating user on database", tx.Error)
	}

	s.logger.WithContext(s.ctx).WithFields(logrus.Fields{
		"id": registry.ID,
	}).Debug("user updated")

	registry.Password = sql.NullString{String: ""}
	return registry, nil
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
