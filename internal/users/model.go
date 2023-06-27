package users

import (
	"time"

	"github.com/wesleyburlani/go-rest-api/pkg/crypto"
	"gorm.io/gorm"
)

type CreateUserProps struct {
	Email    string `validate:"email,required"`
	Password string `validate:"required"`
}

type UpdateUserProps struct {
	Email    string `validate:"email"`
	Password string `validate:"-"`
}

type User struct {
	ID             uint      `json:"id" gorm:"primaryKey;auto_increment:true"`
	Email          string    `json:"email" gorm:"unique,not null,default:null"`
	Password       string    `json:"password,omitempty" gorm:"not null"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	OptionalFields bool      `json:"-" gorm:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, _ = crypto.GenerateHashFromPassword(u.Password)
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password, _ = crypto.GenerateHashFromPassword(u.Password)
	return nil
}
