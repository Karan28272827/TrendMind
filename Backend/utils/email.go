
package utils

import (
	"fmt"
	"net/smtp"
	"os"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func SendResetEmail(toEmail, token string) error {
    from := os.Getenv("EMAIL_SENDER_ADDRESS")
    password := os.Getenv("EMAIL_SENDER_PASSWORD")
    host := os.Getenv("EMAIL_SMTP_HOST")
    port := os.Getenv("EMAIL_SMTP_PORT")

    // Create plain text email body (not clickable link)
    body := fmt.Sprintf(
        "To reset your password, visit the following URL and enter this token:\n\n"+
        "URL: %s/ResetPasswordPage\n"+
        "Token: %s\n\n"+
        "This token will expire in 1 hour.",
        os.Getenv("FRONTEND_URL"),
        token,
    )

    msg := []byte(
        "Subject: Password Reset Request\n" +
        "\n" +
        body,
    )

    auth := smtp.PlainAuth("", from, password, host)
    err := smtp.SendMail(host+":"+port, auth, from, []string{toEmail}, msg)
    if err != nil {
        return fmt.Errorf("failed to send email: %w", err)
    }
    return nil
}

func CreateResetPasswordToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // 1 hour expiration

	secretKey := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func VerifyResetToken(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return "", fmt.Errorf("invalid token: %v", err)
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email, ok := claims["email"].(string)
		if !ok {
			return "", fmt.Errorf("email not found in token")
		}
		return email, nil
	}

	return "", fmt.Errorf("invalid token")
}
