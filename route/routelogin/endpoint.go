package routelogin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/charlesfan/go-api/service/rsi"
	"github.com/charlesfan/go-api/utils/log"
)

func EmailLogin(c *gin.Context) {
	//Implement: Use backend service
	s := rsi.LoginService
	b := c.MustGet("info").(rsi.EmailLoginBody)

	if err := s.EmailChecking(&b); err != nil {
		log.Error(err)

		resp := map[string]string{"error": "EmailChecking return false"}
		code := 401

		c.JSON(code, resp)
		c.Abort()
		return
	}

	responseData := map[string]interface{}{
		"msg": "success",
	}
	c.JSON(http.StatusOK, responseData)
}
