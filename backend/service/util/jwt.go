package util

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

type jwtService struct {
	log       *zap.SugaredLogger
	secretKey string
	ttl       int
}

func NewJWTService(
	log *zap.SugaredLogger,
) JWTService {
	ttl, err := strconv.Atoi(os.Getenv("JWT_TTL"))
	if err != nil {
		log.Fatal(err)
	}
	return &jwtService{
		log:       log,
		secretKey: os.Getenv("JWT_SECRET"),
		ttl:       ttl,
	}
}

func (s *jwtService) GenerateToken(ctx context.Context, id, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  time.Now().Add(time.Hour * time.Duration(s.ttl)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		s.log.Error("Failed to generate token", err.Error())
		return "", err
	}

	return tokenString, nil
}

func (s *jwtService) ValidateToken(ctx context.Context, token string) (id, role string, err error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
	if err != nil {
		s.log.Error("Failed to validate token", err.Error())
		return "", "", err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		s.log.Error("Failed to validate token", err.Error())
		return "", "", err
	}

	id = claims["id"].(string)
	role = claims["role"].(string)
	return id, role, nil
}
