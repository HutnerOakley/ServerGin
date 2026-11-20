package user_api

import (
	"ServeGin/internal/utils/res"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*
退出登录
*/
func (u *UserApi) LogoutView(c *gin.Context) {
	// 获取到请求头中的token
	authHeader := c.Request.Header.Get("Authorization")
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	token := parts[1]

	err := u.userService.Logout(c.Request.Context(), token)
	if err != nil {
		logrus.Errorf("退出登录失败: %v", err)
		res.FailWithMsg("退出登录失败", c)
		return
	}
	res.OKWithMsg("退出登录成功", c)
}
