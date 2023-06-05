package users

import (
	"context"

	"github.com/wesleyburlani/go-rest-api/models"
)

type IUsersService interface {
	WithContext(ctx context.Context) IUsersService
	Create(props models.UserProps) models.User
	Get(id uint) (models.User, error)
	List(page int, limit int) []models.User
	Update(id uint, props models.UserProps) (models.User, error)
}
