package mailer

import (
	"bytes"
	"fmt"
	"github.com/go-mail/mail"
	"mailer-service/token"
	"text/template"
	"log"
	"os"
	"strconv"
)

func SendMail(record *token.Record) {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	senderEmail := os.Getenv("SMTP_SENDER_EMAIL")
	senderName := os.Getenv("SMTP_SENDER_NAME")

	// Convert smtpPortStr to an integer
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Printf("Failed to convert SMTP port to integer: %v\n", err)
		return
	}

	mailer := mail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	m := mail.NewMessage()

	m.SetHeader("From", senderName+" <"+senderEmail+">")
	m.SetHeader("To", record.Email)

	var tmpl *template.Template

	if record.LoginTime != "" {
		tmpl, err = template.ParseFiles("templates/login_template.html")
		if err != nil {
			log.Printf("Failed to load login template: %v\n", err)
			return
		}
		m.SetHeader("Subject", "Notification: Login Information")
	} else if record.PaymentStatus != "" {
		tmpl, err = template.ParseFiles("templates/payment_template.html")
		if err != nil {
			log.Printf("Failed to load payment template: %v\n", err)
			return
		}
		m.SetHeader("Subject", "Notification: Payment Information")
	} else {
		fmt.Printf("Invalid record: %+v\n", record)
		return
	}

	body := &bytes.Buffer{}
	if err := tmpl.Execute(body, record); err != nil {
		log.Printf("Failed to execute template: %v\n", err)
		return
	}

	m.SetBody("text/html", body.String())

	if err := mailer.DialAndSend(m); err != nil {
		log.Printf("Failed to send email: %v\n", err)
	}
}
