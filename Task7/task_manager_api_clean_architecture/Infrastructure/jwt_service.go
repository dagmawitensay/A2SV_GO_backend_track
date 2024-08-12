package infrastructure

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//  Functions to generate and validate JWT tokens.

type JWTService interface {
	GenerateToken(userID string, email string, role string, duration time.Duration) (string, error)
	ValidateToken(token string) (map[string]string, error)
}


type jwtService struct {
	secretKey []byte
}

func NewJWTService(secretKey []byte) JWTService {
	return &jwtService{
		secretKey: secretKey,
	}
}

func (j *jwtService) GenerateToken(userID string, email string, role string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"email": email,
		"role": role,
	})

	jwtToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (j *jwtService) ValidateToken(jwtToken string) (map[string]string, error) {
    token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        return j.secretKey, nil
    })

    if err != nil || !token.Valid {
        return nil, fmt.Errorf("invalid jwt")
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userID, userIDExists := claims["userID"].(string)
        userRole, userRoleExists := claims["role"].(string)
        if !userIDExists || !userRoleExists {
            return nil, fmt.Errorf("invalid jwt claims")
        }

        return map[string]string{
            "userID": userID,
            "role":   userRole,
        }, nil
    }

    return nil, fmt.Errorf("invalid jwt claims")
}