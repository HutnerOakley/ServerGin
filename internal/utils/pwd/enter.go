package pwd

import (
	"golang.org/x/crypto/bcrypt"
)

/*
使用bcrypt进行加密
*/
func GenerateHashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

/*
验证密码是否匹配
*/
func CompareHashAndPassword(hashPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
