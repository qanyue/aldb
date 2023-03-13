package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

// Claims JWT结构
type Claims struct {
	UserEmail string `json:"UserEmail"`
	UserPwd   string `json:"UserPwd"`
	jwt.StandardClaims
}

// TokenExpireDuration 过期时间
const TokenExpireDuration = time.Hour * 24

// Secret 密钥
var Secret = []byte("8./Pw52a.p4N%C37")

// GenToken 生成 Token
func GenToken(email string, password string) (string, error) {
	c := Claims{
		email,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "SukiEva",
		},
	}
	// 用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 用指定的secret签名并获得完整的编码后的字符串 token
	return token.SignedString(Secret)
}

// ParseToken 解析 Token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
