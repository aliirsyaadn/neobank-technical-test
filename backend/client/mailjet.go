package client

import (
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/mailjet/mailjet-apiv3-go/v4"
	"go.uber.org/zap"
)

type mailjetClient struct {
	log               *zap.SugaredLogger
	client            *mailjet.Client
	fromEmail         string
	fromName          string
	sendOTPTemplateID int
}

func NewMailjetClient(
	log *zap.SugaredLogger,
) MailjetClient {
	environment := os.Getenv("ENV")
	if environment == strings.ToUpper("LOCAL") {
		return &mockMailjetClient{}
	}

	publicKey := os.Getenv("MAILJET_API_PUBLIC_KEY")
	privateKey := os.Getenv("MAILJET_API_PRIVATE_KEY")

	fromEmail := os.Getenv("MAILJET_FROM_EMAIL")
	fromName := os.Getenv("MAILJET_FROM_NAME")

	sendOTPTemplateID, err := strconv.Atoi(os.Getenv("MAILJET_SEND_OTP_TEMPLATE_ID"))
	if err != nil {
		log.Fatal(err)
	}

	client := mailjet.NewMailjetClient(publicKey, privateKey)
	return &mailjetClient{
		log:               log,
		client:            client,
		fromEmail:         fromEmail,
		fromName:          fromName,
		sendOTPTemplateID: sendOTPTemplateID,
	}
}

func (c *mailjetClient) SendOTP(ctx context.Context, email, otp string) error {
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: c.fromEmail,
				Name:  c.fromName,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Name:  "Ali Irsyaad Nursyaban",
					Email: email,
				},
			},
			Subject: "NeoBank Email OTP",
			Variables: map[string]interface{}{
				"otp": otp,
			},
			TemplateID:       c.sendOTPTemplateID,
			TemplateLanguage: true,
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := c.client.SendMailV31(&messages)
	if err != nil {
		c.log.Error("Failed to send email", err.Error())
		return err
	}

	c.log.Info("Email sent", res)
	return nil
}
