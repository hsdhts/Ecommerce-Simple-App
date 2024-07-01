package utility

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id string, secret string, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":   id,
		"exp":  jwt.NewNumericDate(time.Now().Add(10 * time.Minute)).Unix(),
		"role": role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string, secret string) (string, string, error) {
	tokens, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", "", err
	}

	if claims, ok := tokens.Claims.(jwt.MapClaims); ok && tokens.Valid {
		id, idOk := claims["id"].(string)
		role, roleOk := claims["role"].(string)
		if !idOk || !roleOk {
			return "", "", fmt.Errorf("invalid token claims")
		}
		return id, role, nil
	}

	return "", "", fmt.Errorf("invalid token")
}
