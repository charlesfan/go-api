package rsi

import (
	"github.com/charlesfan/go-api/repository/user"
	"github.com/charlesfan/go-api/utils/log"
)

type EmailLoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginService struct {
	// ---Repository---
	user user.Repository
	// ---Other---
}

func (s *loginService) EmailChecking(b *EmailLoginBody) error {
	log.Info("Here in EmailChecking function: ")
	log.Info("Got Email: ", b.Email)
	log.Info("PASSWORD: ", b.Password)

	return nil
}

func NewLoginService(u user.Repository) LoginServicer {
	s := &loginService{
		user: u,
	}
	return s
}
