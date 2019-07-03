package sqlite

import (
	"github.com/jinzhu/gorm"

	"github.com/charlesfan/go-api/repository/user"
	"github.com/charlesfan/go-api/utils/log"
)

type userRepository struct {
	copy *gorm.DB
	db   *gorm.DB
}

func (r *userRepository) Get(id user.UUID) (*user.User, error) {
	var d user.User
	if err := r.db.Where("UUID = ?", id).Find(&d).Error; err != nil {
		log.Error("userRepository Get fail => ", err)
		return nil, err
	}

	return &d, nil
}

func (r *userRepository) Create(v *user.User) (*user.User, error) {
	d := r.db.Create(v)
	if err := d.Error; err != nil {
		log.Error("userRepository Create fail => ", err)
		return nil, err
	}
	x := d.Value.(*user.User)
	return x, nil
}

func (r *userRepository) Delete(id user.UUID) error {
	if err := r.db.Where("UUID =?", id).Delete(&user.User{}).Error; err != nil {
		log.Error("userRepository Delete fail => ", err)
		return err
	}
	return nil
}

func (r *userRepository) FindAll() (*[]user.User, error) {
	var p []user.User
	if err := r.db.Find(&p).Error; err != nil {
		log.Error("userRepository FindAll fail => ", err)
		return nil, err
	}
	return &p, nil
}

func (r *userRepository) Query(query interface{}, args ...interface{}) *gorm.DB {
	return r.db.Where(query, args...)
}

func (r *userRepository) NewTransactions() {
	r.db = r.db.Begin()
}

func (r *userRepository) TransactionsRollback() {
	r.db.Rollback()
	r.db = r.copy
}

func (r *userRepository) TransactionsCommit() {
	r.db.Commit()
	r.db = r.copy
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{
		copy: db,
		db:   db,
	}
}
