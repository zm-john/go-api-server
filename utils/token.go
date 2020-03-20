package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// CustomClaim 自定义 token 声明
type CustomClaim struct {
	Sub string `json:"sub"`
	jwt.StandardClaims
}

// JWTConfig jwt config
type JWTConfig struct {
	Key string
	TTL int64
}

// Token 自定义 token
type Token struct {
	// token 值
	Value string `json:"token"`
	// 过期时间
	ExpiresAt int64 `json:"expires_at"`
}

// TokenExpiredError token 过期了
type TokenExpiredError struct{}

func (t TokenExpiredError) Error() string {
	return "token 过期了"
}

// NewToken new a token for key
func NewToken(sub string, conf JWTConfig) (*Token, error) {
	signingKey := []byte(conf.Key)

	// Create the Claims
	ttl := time.Duration(conf.TTL)
	exp := int64(time.Now().Add(time.Second * ttl).Unix())
	claims := &CustomClaim{
		Sub: sub,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	val, err := token.SignedString(signingKey)
	return &Token{Value: val, ExpiresAt: exp}, err
}

// ParseToken parse token that give
func ParseToken(tokenString string, conf JWTConfig) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.Key), nil
	})

	// if err != nil {
	// 	return "", err
	// }

	claims := token.Claims.(*CustomClaim)

	if int64(time.Now().Unix()) > claims.ExpiresAt {
		return "", TokenExpiredError{}
	}

	return claims.Sub, err
}
