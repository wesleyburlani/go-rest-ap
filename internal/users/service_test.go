package users_test

import (
	"io"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"github.com/wesleyburlani/go-rest-api/internal/config"
	"github.com/wesleyburlani/go-rest-api/internal/database"
	"github.com/wesleyburlani/go-rest-api/internal/users"
	custom_errors "github.com/wesleyburlani/go-rest-api/pkg/errors"
	"gorm.io/gorm"
)

type ServiceTestSuite struct {
	suite.Suite
	cfg    *config.Config
	logger *logrus.Logger
	db     *gorm.DB
	svc    *users.Service
}

func (s *ServiceTestSuite) SetupSuite() {
	s.cfg = config.LoadConfig()
	s.logger = logrus.New()
	s.logger.Out = io.Discard
	s.db = database.Init(s.cfg, s.logger)
	s.svc = users.NewService(s.db, s.logger)
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
	s.db.Delete(&users.User{Email: originalUser.Email})

	createdUser, err := s.svc.Create(originalUser)

	if err != nil {
		s.Fail(err.Error())
	}
	defer s.db.Delete(&createdUser)

	s.Equal(nil, err)
	s.NotEmpty(createdUser.ID)
	s.Equal(originalUser.Email, createdUser.Email)
	s.Equal("", createdUser.Password)
}

func (s *ServiceTestSuite) TestCreateUser_UserAlreadyExists() {
	user := users.CreateUserProps{
		Email:    gofakeit.Email(),
		Password: generateRandomPassword(),
	}

	createdUser, err := s.svc.Create(user)
	s.Nil(err)
	_, err = s.svc.Create(user)
	s.True(custom_errors.IsConflictError(err))
	defer s.db.Delete(&createdUser)
}

func (s *ServiceTestSuite) TestCreateUser_Validation() {
	user := users.CreateUserProps{
		Password: generateRandomPassword(),
	}
	createdUser, err := s.svc.Create(user)

	if err == nil {
		s.db.Delete(&createdUser)
	}

	s.True(custom_errors.IsValidationError(err))

	user.Email = gofakeit.Email()
	createdUser, err = s.svc.Create(user)

	if err == nil {
		s.db.Delete(&createdUser)
	}

	s.True(custom_errors.IsValidationError(err))
}

func (s *ServiceTestSuite) TestUpdateUser() {
	createUser := users.CreateUserProps{
		Email:    gofakeit.Email(),
		Password: generateRandomPassword(),
	}

	updateUser := users.UpdateUserProps{
		Email: gofakeit.Email(),
	}

	// make sure it wasn't created before
	s.db.Delete(&users.User{Email: createUser.Email})
	s.db.Delete(&users.User{Email: updateUser.Email})

	createdUser, err := s.svc.Create(createUser)

	if err != nil {
		s.Fail(err.Error())
	}
	defer s.db.Delete(&createdUser)

	updatedUser, err := s.svc.Update(createdUser.ID, updateUser)
	s.Nil(err)
	s.Equal(updateUser.Email, updatedUser.Email)
	s.Equal("", updatedUser.Password)

	updatePassword := users.UpdateUserProps{
		Email:    "",
		Password: generateRandomPassword(),
	}
	updatedPassword, err := s.svc.Update(createdUser.ID, updatePassword)
	s.Nil(err)
	s.Equal(updateUser.Email, updatedPassword.Email)
	s.Equal("", updatedPassword.Password)
	defer s.db.Delete(&updatedUser)
}
