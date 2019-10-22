package login

import "github.com/gin-gonic/gin"

func MakeHandler(r *gin.RouterGroup) {
	g := r.Group("/login")
	{
		g.POST("/email", CheckEmail(), EmailLogin)
	}
}
