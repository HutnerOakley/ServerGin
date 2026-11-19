package user_ser

import (
	"ServeGin/internal/global"
	"ServeGin/internal/models"
	"ServeGin/internal/utils/pwd"
	"context"
	"errors"
)

type UserService struct {
}

/*
创建UserService实例
*/
func NewUserService() *UserService {
	return &UserService{}
}

/*
用户登录Service
*/
func (s *UserService) Login(ctx context.Context, username, password string) (*models.UserModel, error) {
	// 判断Context是否有
	if ctx == nil {
		ctx = context.Background()
	}
	var user models.UserModel
	// 使用WithContext，把上下文传给数据库
	err := global.DB.WithContext(ctx).Take(&user, "user_name=?", username).Error
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}
	if !pwd.CompareHashAndPassword(user.Password, password) {
		return nil, errors.New("用户名或密码错误")
	}
	return &user, nil
}
