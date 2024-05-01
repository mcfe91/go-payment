package email

import (
	"fmt"
	"net/smtp"
)

func Send(target string, orderID string) error {
	// TODO: remove test credentials
	senderEmail := "email@gmail.com"
	password := "password"

	recipientEmail := target

	message := []byte(fmt.Sprintf("Subject: Payment Processed!\n Process ID: %s\n", orderID))

	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	creds := smtp.PlainAuth("", senderEmail, password, smtpServer)

	smtpAddr := fmt.Sprintf("%s:%d", smtpServer, smtpPort)
	err := smtp.SendMail(smtpAddr, creds, senderEmail, []string{recipientEmail}, message)
	if err != nil {
		return err
	}

	return nil
}