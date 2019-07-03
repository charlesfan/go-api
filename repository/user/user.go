package user

import (
	"database/sql/driver"

	"github.com/jinzhu/gorm"

	"github.com/charlesfan/go-api/repository"
)

type UUID string

func (u UUID) Value() (driver.Value, error) { return string(u), nil }

type User struct {
	UUID     UUID   `gorm:"column:uuid;unique;type:uuid;primary_key"`
	Email    string `gorm:"column:email;unique;not null"`
	Password string `gorm:"column:password;not null"`
}

func (User) TableName() string {
	return "user"
}

type Repository interface {
	Get(id UUID) (*User, error)
	Create(p *User) (*User, error)
	Delete(id UUID) error
	FindAll() (*[]User, error)
	Query(query interface{}, args ...interface{}) *gorm.DB
	repository.Transactionser
}
