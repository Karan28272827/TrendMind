package utils

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendVerificationEmail(toEmail string, token string) error {
	from := mail.NewEmail("TrendMind", os.Getenv("SENDER_EMAIL"))
	subject := "Verify your email"
	to := mail.NewEmail("User", toEmail)

	fmt.Printf("\nEntering send verification email: %s not available", toEmail)

	verificationLink := fmt.Sprintf("%s/verify?token=%s", os.Getenv("FRONTEND_URL"), token)

	plainTextContent := "Click the following link to verify your email: \n" + verificationLink
	htmlContent := fmt.Sprintf("<p>Click the link to verify your email:</p><a href='%s'>Verify Email</a>", verificationLink)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(message)
	if err != nil {
		return err
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("failed to send email %s", response.Body)
	}
	return nil

}
