package users

import (
	"database/sql"
	"time"

	"github.com/wesleyburlani/go-rest-api/pkg/crypto"
	"gorm.io/gorm"
)

type CreateUserProps struct {
	Email    sql.NullString `validate:"email,required"`
	Password sql.NullString `validate:"required"`
}

type UpdateUserProps struct {
	Email    sql.NullString `validate:"email"`
	Password sql.NullString `validate:"-"`
}

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;auto_increment:true"`
	Email     sql.NullString `json:"email" gorm:"unique,not null,default:null"`
	Password  sql.NullString `json:"password,omitempty" gorm:"not null"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password.String, _ = crypto.GenerateHashFromPassword(u.Password.String)
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password.String, _ = crypto.GenerateHashFromPassword(u.Password.String)
	return nil
}
