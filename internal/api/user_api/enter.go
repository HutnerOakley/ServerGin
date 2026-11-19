package user_api

import "ServeGin/internal/service/user_ser"

/*
用户api结构体
*/
type UserApi struct {
	userService *user_ser.UserService
}
