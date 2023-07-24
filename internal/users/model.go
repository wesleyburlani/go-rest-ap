package users

import (
	"database/sql"
	"time"

	"github.com/wesleyburlani/go-rest-api/internal/db"
	"gopkg.in/guregu/null.v4"
)

type UserRole string

const (
	AdminRole   UserRole = "admin"
	DefaultRole UserRole = "default"
)

type NullUserRole struct {
	Role  UserRole
	Valid bool
}

type CreateUserProps struct {
	Username string   `validate:"required"`
	Email    string   `validate:"required,email"`
	Password string   `validate:"required"`
	Role     UserRole `validate:"required,oneof=admin default"`
}

func (p *CreateUserProps) ToDB() db.CreateUserParams {
	return db.CreateUserParams{
		Username: p.Username,
		Email:    p.Email,
		Password: p.Password,
		Role:     string(p.Role),
	}
}

type UpdateUserProps struct {
	Username null.String  `validate:""`
	Email    null.String  `validate:"omitempty,email"`
	Password null.String  `validate:""`
	Role     NullUserRole `validate:"omitempty,oneof=admin default"`
}

func (p *UpdateUserProps) ToDB() db.UpdateUserParams {
	return db.UpdateUserParams{
		Username: p.Username.NullString,
		Email:    p.Email.NullString,
		Password: p.Password.NullString,
		Role:     sql.NullString{String: string(p.Role.Role), Valid: p.Role.Valid},
	}
}

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// parse user from database model
func NewUserFromDB(user db.User) User {
	return User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Role:      UserRole(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
