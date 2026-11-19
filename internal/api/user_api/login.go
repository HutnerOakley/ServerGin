package user_api

import (
	"ServeGin/internal/utils/jwts"
	"ServeGin/internal/utils/res"
	"ServeGin/internal/utils/validate"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*
登录请求
*/
type LoginRequest struct {
	UserName string `json:"username" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}

/*
登录响应
*/
type LoginResponse struct {
	Token    string `json:"token"`
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
}

/*
登录视图函数
*/
func (u *UserApi) LoginView(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		msg := validate.ValidateError(err)
		res.FailWithMsg(msg, c)
		return
	}

	// 【核心】把 Gin 的标准 Context 剥离出来传进去
	// c.Request.Context() 是标准库的 context.Context
	user, err := u.userService.Login(c.Request.Context(), req.UserName, req.Password)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	// 签发token
	token, err := jwts.GenToken(jwts.Claims{
		UserID: user.ID,
		RoleID: user.RoleID,
	})
	if err != nil {
		logrus.Errorf("生成token失败 %s", err)
		res.FailWithMsg("登录失败", c)
		return
	}

	data := LoginResponse{
		Token:    token,
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
	}
	res.OkWithData(data, c)
}
