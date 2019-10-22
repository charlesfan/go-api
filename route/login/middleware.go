package login

import (
	"github.com/gin-gonic/gin"

	"github.com/charlesfan/go-api/service/rsi"
	"github.com/charlesfan/go-api/utils/log"
)

func CheckEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Implement
		b := rsi.EmailLoginBody{}

		if err := c.Bind(&b); err != nil {
			log.Error(err)
			log.Error("Email or Password does not exist")

			resp := map[string]string{"error": "Email or Password does not exist"}
			code := 401

			c.JSON(code, resp)
			c.Abort()
			return
		}
		log.Info("Here is CheckEmail Middleware Function: PASS")
		c.Set("info", b)
		c.Next()
	}
}
