package users

import (
	"time"

	"github.com/wesleyburlani/go-rest-api/pkg/crypto"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;auto_increment:true" binding:"-"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, _ = crypto.GenerateHashFromPassword(u.Password)
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password, _ = crypto.GenerateHashFromPassword(u.Password)
	return nil
}
