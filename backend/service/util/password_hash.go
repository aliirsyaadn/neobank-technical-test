package util

import (
	"context"
	"os"
	"strconv"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type passwordHashService struct {
	log        *zap.SugaredLogger
	bcryptCost int
}

func NewPasswordHashService(
	log *zap.SugaredLogger,
) PasswordHashService {
	bcryptCost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if err != nil {
		log.Fatal(err)
	}

	return &passwordHashService{
		log:        log,
		bcryptCost: bcryptCost,
	}
}

func (s *passwordHashService) HashPassword(ctx context.Context, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), s.bcryptCost)
	if err != nil {
		s.log.Error("Failed to hash password", err.Error())
		return "", err
	}

	return string(hashedPassword), nil
}

func (s *passwordHashService) ComparePassword(ctx context.Context, password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		s.log.Error("Failed to compare password", err.Error())
		return err
	}
	return nil
}
