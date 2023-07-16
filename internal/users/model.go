package users

import (
	"database/sql"
	"time"
)

type CreateUserProps struct {
	Email    string `validate:"email,required"`
	Password string `validate:"required"`
}

type UpdateUserProps struct {
	Email    sql.NullString `validate:"email"`
	Password sql.NullString `validate:"-"`
}

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
