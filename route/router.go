package route

import (
	"github.com/gin-gonic/gin"

	"github.com/charlesfan/go-api/route/login"
)

func Init() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	v1 := r.Group("/v1")
	{
		//Login
		login.MakeHandler(v1)
	}

	return r
}
