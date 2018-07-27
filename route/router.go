package route

import (
	"github.com/gin-gonic/gin"

	"github.com/charlesfan/go-api/route/routelogin"
)

func Init() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	v1 := r.Group("/v1")
	{
		//Login
		routelogin.MakeHandler(v1)
	}

	return r
}
