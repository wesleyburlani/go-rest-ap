package users_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/wesleyburlani/go-rest-api/internal/config"
	"github.com/wesleyburlani/go-rest-api/internal/db"
	"github.com/wesleyburlani/go-rest-api/internal/users"
	"github.com/wesleyburlani/go-rest-api/pkg/crypto"
	custom_errors "github.com/wesleyburlani/go-rest-api/pkg/errors"
)

type ServiceTestSuite struct {
	suite.Suite
	ctx    context.Context
	cfg    *config.Config
	logger *logrus.Logger
	db     *db.Database
	auth   *crypto.JwtAuth
	svc    *users.Service
}

func (s *ServiceTestSuite) SetupSuite() {
	s.ctx = context.Background()
	s.cfg = config.LoadConfig()
	s.logger = logrus.New()
	s.db = db.NewDatabase(s.cfg, s.logger)
	s.auth = crypto.NewJwtAuth([]byte(s.cfg.JwtSecretKey))
	s.svc = users.NewService(s.db, s.logger, s.auth)
}

func generateRandomPassword() string {
	return gofakeit.Password(true, true, true, true, false, 10)
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (s *ServiceTestSuite) TestCreateUser() {
	originalUser := users.CreateUserProps{
		Email:    gofakeit.Email(),
		Password: generateRandomPassword(),
	}

	// make sure it wasn't created before
	s.db.Queries.DeleteUserByEmail(s.ctx, originalUser.Email)

	createdUser, err := s.svc.Create(originalUser)

	if err != nil {
		s.Fail(err.Error())
	}
	defer s.db.Queries.DeleteUserByEmail(s.ctx, originalUser.Email)

	s.Equal(nil, err)
	s.NotZero(createdUser.ID)
	s.Equal(originalUser.Email, createdUser.Email)
	s.Equal("", createdUser.Password)
}

func (s *ServiceTestSuite) TestCreateUser_UserAlreadyExists() {
	user := users.CreateUserProps{
		Email:    gofakeit.Email(),
		Password: generateRandomPassword(),
	}

	_, err := s.svc.Create(user)
	s.Nil(err)
	_, err = s.svc.Create(user)
	s.logger.Info(err)
	s.True(custom_errors.IsConflictError(err))
	defer s.db.Queries.DeleteUserByEmail(s.ctx, user.Email)
}

func (s *ServiceTestSuite) TestGet() {
	user := users.CreateUserProps{
		Email:    gofakeit.Email(),
		Password: generateRandomPassword(),
	}

	createdUser, err := s.svc.Create(user)
	if err != nil {
		s.Fail(err.Error())
	}
	defer s.db.Queries.DeleteUserByEmail(s.ctx, user.Email)

	userById, err := s.svc.Get(createdUser.ID)
	if err != nil {
		s.Fail(err.Error())
	}

	s.Equal(createdUser.Email, userById.Email)
	s.Equal("", userById.Password)
}

func (s *ServiceTestSuite) TestGet_UserDoesNotExist() {
	_, err := s.svc.Get(999999999)
	s.True(custom_errors.IsNotFoundError(err))
}

func (s *ServiceTestSuite) TestGetByEmail() {
	user := users.CreateUserProps{
		Email:    gofakeit.Email(),
		Password: generateRandomPassword(),
	}

	createdUser, err := s.svc.Create(user)
	if err != nil {
		s.Fail(err.Error())
	}
	defer s.db.Queries.DeleteUserByEmail(s.ctx, user.Email)

	userByEmail, err := s.svc.GetByEmail(createdUser.Email)
	if err != nil {
		s.Fail(err.Error())
	}

	s.Equal(createdUser.Email, userByEmail.Email)
	s.Equal("", userByEmail.Password)
}

func (s *ServiceTestSuite) TestGetByEmail_UserDoesNotExist() {
	_, err := s.svc.GetByEmail(gofakeit.Email())
	s.True(custom_errors.IsNotFoundError(err))
}

func (s *ServiceTestSuite) TestListUsers() {
	user1 := users.CreateUserProps{
		Email:    gofakeit.Email(),
		Password: generateRandomPassword(),
	}
	user2 := users.CreateUserProps{
		Email:    gofakeit.Email(),
		Password: generateRandomPassword(),
	}

	createdUser1, err := s.svc.Create(user1)
	s.logger.Info(createdUser1)
	if err != nil {
		s.Fail(err.Error())
	}
	defer s.db.Queries.DeleteUserByEmail(s.ctx, user1.Email)
	createdUser2, err := s.svc.Create(user2)
	s.logger.Info(user2)
	if err != nil {
		s.Fail(err.Error())
	}
	defer s.db.Queries.DeleteUserByEmail(s.ctx, user2.Email)

	users, err := s.svc.List(0, 100)
	if err != nil {
		s.Fail(err.Error())
	}

	s.Equal(2, len(users))
	s.Equal(createdUser1.Email, users[0].Email)
	s.Equal(createdUser2.Email, users[1].Email)
}

func (s *ServiceTestSuite) TestListUsers_Empty() {
	users, err := s.svc.List(0, 100)
	if err != nil {
		s.Fail(err.Error())
	}

	s.Equal(0, len(users))
}

func (s *ServiceTestSuite) TestUpdateUser() {
	user := users.CreateUserProps{
		Email:    gofakeit.Email(),
		Password: generateRandomPassword(),
	}

	createdUser, err := s.svc.Create(user)
	if err != nil {
		s.Fail(err.Error())
	}
	defer s.db.Queries.DeleteUserByEmail(s.ctx, user.Email)

	updateUser := users.UpdateUserProps{
		Email:    sql.NullString{String: gofakeit.Email(), Valid: true},
		Password: sql.NullString{String: generateRandomPassword(), Valid: true},
	}

	updatedUser, err := s.svc.Update(createdUser.ID, updateUser)
	if err != nil {
		s.Fail(err.Error())
	}

	s.Equal(updateUser.Email.String, updatedUser.Email)
	s.Equal("", updatedUser.Password)
}

func (s *ServiceTestSuite) TestUpdateUser_UserDoesNotExist() {
	updateUser := users.UpdateUserProps{
		Email:    sql.NullString{String: gofakeit.Email(), Valid: true},
		Password: sql.NullString{String: generateRandomPassword(), Valid: true},
	}

	_, err := s.svc.Update(999999999, updateUser)
	s.True(custom_errors.IsNotFoundError(err))
}

func (s *ServiceTestSuite) TestUpdateUser_UpdateEmail() {
	user := users.CreateUserProps{
		Email:    gofakeit.Email(),
		Password: generateRandomPassword(),
	}

	createdUser, err := s.svc.Create(user)
	if err != nil {
		s.Fail(err.Error())
	}

	createdUserWithPwd, err := s.db.Queries.GetUserById(s.ctx, createdUser.ID)
	if err != nil {
		s.Fail(err.Error())
	}

	defer s.db.Queries.DeleteUserByEmail(s.ctx, user.Email)

	updateUser := users.UpdateUserProps{
		Email:    sql.NullString{String: gofakeit.Email(), Valid: true},
		Password: sql.NullString{String: "", Valid: false},
	}

	updatedUser, err := s.svc.Update(createdUser.ID, updateUser)
	if err != nil {
		s.Fail(err.Error())
	}

	updatedUserWithPwd, err := s.db.Queries.GetUserById(s.ctx, createdUser.ID)
	if err != nil {
		s.Fail(err.Error())
	}

	defer s.db.Queries.DeleteUserByEmail(s.ctx, updateUser.Email.String)

	s.Equal(updateUser.Email.String, updatedUser.Email)
	s.Equal("", updatedUser.Password)
	s.Equal(updatedUserWithPwd.Password, createdUserWithPwd.Password)
}

func (s *ServiceTestSuite) TestDelete() {
	user := users.CreateUserProps{
		Email:    gofakeit.Email(),
		Password: generateRandomPassword(),
	}

	createdUser, err := s.svc.Create(user)
	if err != nil {
		s.Fail(err.Error())
	}

	err = s.svc.Delete(createdUser.ID)
	if err != nil {
		s.Fail(err.Error())
	}

	_, err = s.svc.Get(createdUser.ID)
	s.True(custom_errors.IsNotFoundError(err))
}

func (s *ServiceTestSuite) TestDelete_UserDoesNotExist() {
	err := s.svc.Delete(999999999)
	s.True(custom_errors.IsNotFoundError(err))
}
