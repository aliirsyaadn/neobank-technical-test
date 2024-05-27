package service

import (
	"context"
	"errors"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/aliirsyaadn/neobank-technical-test/client"
	"github.com/aliirsyaadn/neobank-technical-test/constant"
	"github.com/aliirsyaadn/neobank-technical-test/entity"
	"github.com/aliirsyaadn/neobank-technical-test/model"
	"github.com/aliirsyaadn/neobank-technical-test/repository"
	"github.com/aliirsyaadn/neobank-technical-test/service/util"
	"go.uber.org/zap"
)

type userService struct {
	log                 *zap.SugaredLogger
	mailjetClient       client.MailjetClient
	jwtService          util.JWTService
	passwordHashService util.PasswordHashService
	userRepository      repository.UserRepository
	userOTPRepository   repository.UserOTPRepository
	otpTTL              int
}

func NewUserService(
	log *zap.SugaredLogger,
	mailjetClient client.MailjetClient,
	jwtService util.JWTService,
	passwordHashService util.PasswordHashService,
	userRepository repository.UserRepository,
	userOTPRepository repository.UserOTPRepository,
) UserService {

	otpTTL, err := strconv.Atoi(os.Getenv("OTP_TTL"))
	if err != nil {
		log.Fatal(err)
	}

	return &userService{
		log:                 log,
		mailjetClient:       mailjetClient,
		jwtService:          jwtService,
		passwordHashService: passwordHashService,
		userRepository:      userRepository,
		userOTPRepository:   userOTPRepository,
		otpTTL:              otpTTL,
	}
}

func (s *userService) Login(ctx context.Context, req entity.LoginUserRequest) (res entity.LoginUserResponse, err error) {
	user, err := s.userRepository.Get(ctx, req.ID)
	if err != nil {
		s.log.Error("Failed to get user", err.Error())
		return res, err
	}

	if user.CorporateAccountNo != req.CorporateAccountNo {
		s.log.Error("Invalid corporate account no")
		return res, errors.New("invalid corporate account no")
	}

	err = s.passwordHashService.ComparePassword(ctx, req.Password, user.Password)
	if err != nil {
		s.log.Error("Failed to compare password", err.Error())
		return res, errors.New("invalid password")
	}

	token, err := s.jwtService.GenerateToken(ctx, user.ID, user.Role)
	if err != nil {
		s.log.Error("Failed to generate token", err.Error())
		return res, err
	}

	res = entity.LoginUserResponse{
		Token: token,
		User:  user,
	}

	return res, nil
}

func (s *userService) Register(ctx context.Context, req entity.RegisterUserRequest) (res entity.RegisterUserResponse, err error) {
	valid, ok := constant.USER_ROLES[req.Role]
	if !ok || !valid {
		s.log.Error("Invalid role")
		return res, errors.New("invalid role")
	}

	// TODO: validate unique email, corporate account no, and id by logic service to mapping error

	userOTP, err := s.userOTPRepository.Get(ctx, req.Email, req.OTP)
	if err != nil {
		s.log.Error("Failed to get user otp", err.Error())
		return res, errors.New("invalid otp")
	}

	if userOTP.ExpiredAt.Before(time.Now()) {
		s.log.Error("OTP is expired")
		return res, errors.New("otp is expired")
	}

	hashedPassword, err := s.passwordHashService.HashPassword(ctx, req.PasswordStr)
	if err != nil {
		s.log.Error("Failed to hash password", err.Error())
		return res, err
	}

	user := model.User{
		ID:                 req.ID,
		CorporateAccountNo: req.CorporateAccountNo,
		CorporateName:      req.CorporateName,
		Name:               req.Name,
		Role:               req.Role,
		Phone:              req.Phone,
		Email:              req.Email,
		Password:           hashedPassword,
		Verified:           true,
	}

	err = s.userRepository.Insert(ctx, &user)
	if err != nil {
		s.log.Error("Failed to insert user", err.Error())
		return res, err
	}

	res = entity.RegisterUserResponse{
		User: user,
	}

	return res, nil
}

func (s *userService) SendOTP(ctx context.Context, req entity.SendOTPRequest) error {
	_, err := s.userRepository.GetByEmail(ctx, req.Email)
	if err == nil {
		s.log.Error("Email already registered")
		return errors.New("email already registered")
	}

	otp := s.generateOTP()
	userOTP := model.UserOTP{
		Email:     req.Email,
		OTP:       otp,
		ExpiredAt: time.Now().Add(time.Duration(s.otpTTL) * time.Second),
	}

	err = s.userOTPRepository.Insert(ctx, &userOTP)
	if err != nil {
		s.log.Error("Failed to insert user otp", err.Error())
		return err
	}

	err = s.mailjetClient.SendOTP(ctx, req.Email, otp)
	if err != nil {
		s.log.Error("Failed to send otp", err.Error())
		return err
	}

	return nil
}

func (s *userService) generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	return strconv.Itoa(rand.Intn(max-min+1) + min)
}
