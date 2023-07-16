package users

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/wesleyburlani/go-rest-api/internal/db"
	"github.com/wesleyburlani/go-rest-api/pkg/crypto"
	custom_errors "github.com/wesleyburlani/go-rest-api/pkg/errors"
)

type IService interface {
	WithContext(ctx context.Context) IService
	Create(user CreateUserProps) (User, error)
	GetById(id uint) (User, error)
	GetByEmail(email string) (User, error)
	List(page int32, limit int32) ([]User, error)
	Update(id uint, user UpdateUserProps) (User, error)
}

type Service struct {
	db     *db.Database
	logger *logrus.Logger
	ctx    context.Context
}

func NewService(db *db.Database, logger *logrus.Logger) *Service {
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

func (s *Service) Create(user CreateUserProps) (User, error) {
	encryptPassword, err := crypto.GenerateHashFromPassword(user.Password)

	if err != nil {
		return User{}, custom_errors.NewUnknownError(err.Error())
	}

	if _, err := s.db.Queries.GetUserByEmail(s.ctx, user.Email); err == nil {
		return User{}, custom_errors.NewConflictError("user already exists")
	}

	createdUser, err := s.db.Queries.CreateUser(s.ctx, db.CreateUserParams{
		Email:    user.Email,
		Password: encryptPassword,
	})

	if err != nil {
		return User{}, custom_errors.NewUnknownError(err.Error())
	}

	return User{
		ID:        createdUser.ID,
		Email:     createdUser.Email,
		Password:  "",
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}, nil
}

func (s *Service) GetById(id int64) (User, error) {
	user, err := s.db.Queries.GetUserById(s.ctx, id)

	if err != nil {
		return User{}, custom_errors.NewNotFoundError("user not found")
	}

	return User{
		ID:        user.ID,
		Email:     user.Email,
		Password:  "",
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *Service) GetByEmail(email string) (User, error) {
	user, err := s.db.Queries.GetUserByEmail(s.ctx, email)

	if err != nil {
		return User{}, custom_errors.NewNotFoundError("user not found")
	}

	return User{
		ID:        user.ID,
		Email:     user.Email,
		Password:  "",
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
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
		result = append(result, User{
			ID:        user.ID,
			Email:     user.Email,
			Password:  "",
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return result, nil
}

func (s *Service) Update(id int64, user UpdateUserProps) (User, error) {
	if _, err := s.db.Queries.GetUserById(s.ctx, id); err != nil {
		return User{}, custom_errors.NewNotFoundError("user not found")
	}

	updatedUser, err := s.db.Queries.UpdateUser(s.ctx, db.UpdateUserParams{
		ID:       id,
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		return User{}, custom_errors.NewUnknownError(err.Error())
	}

	return User{
		ID:        updatedUser.ID,
		Email:     updatedUser.Email,
		Password:  "",
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}, nil
}
