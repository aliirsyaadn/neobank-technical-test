package client

import "context"

type mockMailjetClient struct {
}

func (c *mockMailjetClient) SendOTP(ctx context.Context, email, otp string) error {
	return nil
}
