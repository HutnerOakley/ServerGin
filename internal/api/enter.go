package api

import "ServeGin/internal/api/user_api"

/*
api结构体
*/
type Api struct {
	UserApi user_api.UserApi // 用户api
}

var App = new(Api)
