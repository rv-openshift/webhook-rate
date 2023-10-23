package mail

import (
	"fmt"
	"net/smtp"
)

func SendmaiL(email string, password string) error {	
	to := []string{
		email,
	}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	message := []byte("Subject: Test-email!\r\n"+"This is a test email message.\r\n")
	auth := smtp.PlainAuth("", email, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, "no-reply@gmail.com", to, message)
	if err != nil {
		return err
	}
	fmt.Println("Email Sent Successfully!")
	return err
}