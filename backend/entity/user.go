package entity

import "github.com/aliirsyaadn/neobank-technical-test/model"

type RegisterUserRequest struct {
	ID                 string `json:"id" binding:"required"`
	CorporateAccountNo string `json:"corporate_account_no" binding:"required"`
	CorporateName      string `json:"corporate_name" binding:"required"`
	Name               string `json:"name" binding:"required"`
	Role               string `json:"role" binding:"required"`
	Phone              string `json:"phone" binding:"required,e164"`
	PasswordStr        string `json:"password" binding:"required"`
	Email              string `json:"email" binding:"required,email"`
	OTP                string `json:"otp" binding:"required"`
}

type RegisterUserResponse struct {
	model.User
}

type LoginUserRequest struct {
	CorporateAccountNo string `json:"corporate_account_no" binding:"required"`
	ID                 string `json:"id" binding:"required"`
	Password           string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
	model.User
}

type SendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type GetListUserRequest struct {
	CorporateAccountNos []string
}
