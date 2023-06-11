package crypto

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JwtAuth struct {
	secretKey []byte
}

type JwtProps struct {
	Username string
}

type JwtClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewJwtAuth(secretKey []byte) *JwtAuth {
	return &JwtAuth{
		secretKey: secretKey,
	}
}

func (instance *JwtAuth) GenerateToken(props JwtProps) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &JwtClaims{
		Username: props.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(instance.secretKey)
}

func (instance *JwtAuth) DecodeToken(token string) (*JwtProps, error) {
	var claims JwtClaims
	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return instance.secretKey, nil
	})

	if err != nil {
		return &JwtProps{}, err
	}

	if !tkn.Valid {
		return &JwtProps{}, errors.New("invalid token")
	}

	props := JwtProps{
		Username: claims.Username,
	}
	return &props, nil
}
