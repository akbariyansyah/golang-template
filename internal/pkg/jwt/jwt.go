package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWT interface {
	GenerateToken(email string) (*TokenDetails, error)
	VerifyToken(tokenString string) (string, error)
}

type JWTManager struct {
	SecretKey       string
	TokenExpiration time.Duration
}

func NewJWTManager(secretKey string, tokenExpiration string) *JWTManager {
	expiration := time.Duration(0)
	if tokenExpiration != "" {
		expiration, _ = time.ParseDuration(tokenExpiration)
	}

	return &JWTManager{
		SecretKey:       secretKey,
		TokenExpiration: expiration,
	}
}

func (j *JWTManager) GenerateToken(email string) (*TokenDetails, error) {
	expirationTime := time.Now().Add(j.TokenExpiration)
	claims := jwt.MapClaims{
		"email": email,
		"exp":   expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return nil, err
	}

	return &TokenDetails{
		AccessToken: signedToken,
		ExpiredAt:   expirationTime,
	}, nil
}

func (j *JWTManager) VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("email not found in token")
	}

	return email, nil
}
