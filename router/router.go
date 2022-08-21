package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// NewRouter 入口路由
func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))

	v1 := r.Group("api/v1")
	{
		// ping pong
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "ping success")
		})

		// 用户操作
		v1.POST("/user/register", api.UserRegister)
	}
}
