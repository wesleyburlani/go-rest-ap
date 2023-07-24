package users

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/wesleyburlani/go-rest-api/internal/db"
	"github.com/wesleyburlani/go-rest-api/pkg/crypto"
	custom_errors "github.com/wesleyburlani/go-rest-api/pkg/errors"
	"github.com/wesleyburlani/go-rest-api/pkg/validation"
	"gopkg.in/guregu/null.v4"
)

type IService interface {
	WithContext(ctx context.Context) IService
	Create(user CreateUserProps) (User, error)
	Get(id int64) (User, error)
	GetByEmail(email string) (User, error)
	List(page int32, limit int32) ([]User, error)
	Update(id int64, user UpdateUserProps) (User, error)
	Delete(id int64) error
}

type Service struct {
	db        *db.Database
	logger    *logrus.Logger
	ctx       context.Context
	auth      *crypto.JwtAuth
	validator *validation.Validator
}

func NewService(
	db *db.Database,
	logger *logrus.Logger,
	auth *crypto.JwtAuth,
	validator *validation.Validator,
) *Service {
	return &Service{
		db:        db,
		logger:    logger,
		ctx:       context.Background(),
		auth:      auth,
		validator: validator,
	}
}

func (s *Service) WithContext(ctx context.Context) IService {
	s.ctx = ctx
	return s
}

func (s *Service) Create(user CreateUserProps) (User, error) {
	if err := s.validator.Validate(user); err != nil {
		return User{}, custom_errors.NewValidationError(err.Error())
	}

	if encryptPassword, err := crypto.GenerateHashFromPassword(user.Password); err == nil {
		user.Password = encryptPassword
	} else {
		return User{}, custom_errors.NewUnknownError(err.Error())
	}

	if _, err := s.db.Queries.GetUserByEmail(s.ctx, user.Email); err == nil {
		return User{}, custom_errors.NewConflictError("user already exists")
	}

	createdUser, err := s.db.Queries.CreateUser(s.ctx, user.ToDB())

	if err != nil {
		return User{}, custom_errors.NewUnknownError(err.Error())
	}

	return NewUserFromDB(createdUser), nil
}

func (s *Service) Get(id int64) (User, error) {
	user, err := s.db.Queries.GetUserById(s.ctx, id)

	if err != nil {
		return User{}, custom_errors.NewNotFoundError("user not found")
	}

	return NewUserFromDB(user), nil
}

func (s *Service) GetByEmail(email string) (User, error) {
	user, err := s.db.Queries.GetUserByEmail(s.ctx, email)

	if err != nil {
		return User{}, custom_errors.NewNotFoundError("user not found")
	}

	return NewUserFromDB(user), nil
}

func (s *Service) List(page int32, limit int32) ([]User, error) {
	users, err := s.db.Queries.ListUsers(s.ctx, db.ListUsersParams{
		Offset: page * limit,
		Limit:  limit,
	})

	if err != nil {
		return []User{}, custom_errors.NewUnknownError(err.Error())
	}

	var result []User

	for _, user := range users {
		result = append(result, NewUserFromDB(user))
	}

	return result, nil
}

func (s *Service) Update(id int64, user UpdateUserProps) (User, error) {
	if err := s.validator.Validate(user); err != nil {
		return User{}, custom_errors.NewValidationError(err.Error())
	}

	if _, err := s.db.Queries.GetUserById(s.ctx, id); err != nil {
		return User{}, custom_errors.NewNotFoundError("user not found")
	}

	if user.Password.Valid {
		p, err := crypto.GenerateHashFromPassword(user.Password.String)
		if err != nil {
			return User{}, custom_errors.NewUnknownError(err.Error())
		}
		user.Password = null.NewString(p, true)
	}

	updatedUser, err := s.db.Queries.UpdateUser(s.ctx, user.ToDB())

	if err != nil {
		return User{}, custom_errors.NewUnknownError(err.Error())
	}

	return NewUserFromDB(updatedUser), nil
}

func (s *Service) Delete(id int64) error {
	if _, err := s.db.Queries.GetUserById(s.ctx, id); err != nil {
		return custom_errors.NewNotFoundError("user not found")
	}

	if _, err := s.db.Queries.DeleteUserById(s.ctx, id); err != nil {
		return custom_errors.NewUnknownError(err.Error())
	}

	return nil
}
