package client

import "context"

type MailjetClient interface {
	SendOTP(ctx context.Context, email, otp string) error
}
