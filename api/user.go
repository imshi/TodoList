package api

import (
	"todolist/service"

	"github.com/gin-gonic/gin"
)

// UserRegister @Tags USER
// @Summary 用户注册
// @Produce json
// @Accept json
// @Param data body service.UserService true "用户名，密码"
func UserRegister(c *gin.Context) {
	// 接收 request 请求数据，执行注册逻辑并返回
	var userRegisterService service.UserService
	if err := c.ShouldBind(&userRegisterService); err != nil {
		c.JSON(400, ErrorResponse(err))
	} else {
		res := userRegisterService.Register()
		c.JSON(200, res)
	}
}

