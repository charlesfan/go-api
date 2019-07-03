package rsi

import (
	"github.com/charlesfan/go-api/repository"
	"github.com/charlesfan/go-api/repository/sqlite"
	"github.com/charlesfan/go-api/repository/user"
)

var (
	// === Repository ===
	userRepo user.Repository
	// === Service ===
	LoginService LoginServicer
)

func Init(db *repository.Database) {
	// === Repository ===
	userRepo = sqlite.NewUserRepository(db.Gdb)
	// === Service ===
	LoginService = NewLoginService(userRepo)
}
