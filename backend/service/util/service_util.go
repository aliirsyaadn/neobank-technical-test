package util

import (
	"context"
)

type PasswordHashService interface {
	HashPassword(ctx context.Context, password string) (string, error)
	ComparePassword(ctx context.Context, password, hashedPassword string) error
}

type JWTService interface {
	GenerateToken(ctx context.Context, id, role string) (string, error)
	ValidateToken(ctx context.Context, token string) (id, role string, err error)
}
