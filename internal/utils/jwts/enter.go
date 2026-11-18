package jwts

import (
	"ServeGin/internal/global"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenInvalid = errors.New("无效的token")
	ErrTokenExpired = errors.New("token 已过期")
)

/*
自定义载荷
*/
type Claims struct {
	UserID uint `json:"userID"`
	RoleID int8 `json:"roleID"`
}

/*
jwt载荷
*/
type MyClaims struct {
	Claims
	jwt.RegisteredClaims // jwt内嵌的标准说明
}

/*
生成jwt
*/
func GenToken(data Claims) (string, error) {
	//1 构造claims
	claims := MyClaims{
		Claims: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.JWT.Expires) * time.Hour)), // 过期时间
			Issuer:    global.Config.JWT.Issuer,                                                                 // 签发人
		},
	}
	// 2 指定签名算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用秘钥进行签名
	return token.SignedString([]byte(global.Config.JWT.Secret))
}

/*
验证jwt
*/
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析、验证并返回token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 关键安全步骤：校验签名算法，防止算法混淆攻击
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("非预期的签名算法: %v", token.Header["alg"])
		}
		return []byte(global.Config.JWT.Secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrTokenInvalid
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}
