package rsi

import (
	"github.com/charlesfan/go-api/utils/log"
)

type EmailLoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginService struct {
	// ---Repository---
	//user user.Repository
	// ---Other---
}

func (*loginService) EmailChecking(b *EmailLoginBody) error {
	log.Info("Here in EmailChecking function: ")
	log.Info("Got Email: ", b.Email)
	log.Info("PASSWORD: ", b.Password)

	return nil
}

func NewLoginService() LoginServicer {
	s := &loginService{}

	return s
}
