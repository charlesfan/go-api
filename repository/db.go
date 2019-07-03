package repository

import (
	"github.com/jinzhu/gorm"
)

type Database struct {
	Gdb *gorm.DB
}

func (w *Database) IsConnected() bool {
	if w.Gdb == nil {
		return false
	}
	return true
}

func (w *Database) Close() bool {
	w.Gdb.Close()
	return true
}

func NewDatabase(db *gorm.DB) *Database {
	return &Database{
		Gdb: db,
	}
}
