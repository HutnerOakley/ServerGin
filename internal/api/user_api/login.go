package user_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
登录视图函数
*/
func (u *UserApi) LoginView(c *gin.Context) {
	c.String(http.StatusOK, "登录成功")
}
