package user_ser

import (
	"ServeGin/internal/global"
	"ServeGin/internal/models"
	"ServeGin/internal/service/redis_service"
	"ServeGin/internal/utils/jwts"
	"ServeGin/internal/utils/pwd"
	"context"
	"errors"
	"time"
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

/*
用户退出登录Service
*/
func (s *UserService) Logout(ctx context.Context, token string) error {

	if ctx == nil {
		ctx = context.Background()
	}

	// 1.验证token
	claims, err := jwts.ParseToken(token)
	if err != nil {
		return err
	}

	// 2.放入redis的黑名单
	sub := claims.ExpiresAt.Sub(time.Now())
	err = redis_service.AddTokenBlacklist(ctx, token, sub)
	if err != nil {
		return err
	}
	return nil
}
