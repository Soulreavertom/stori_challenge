package emailservice

import (
	"log"
	"net/smtp"
	"os"
)

func SendEmail(subject string, body string, to string) {

	from := os.Getenv("EMAIL")
	pass := os.Getenv("KEYM")

	//subject := config.GetEmailSubject()

	headers := []string{
		"MIME-Version: 1.0",
		"Content-Type: text/html; charset=UTF-8",
		"Subject: " + subject,
	}

	message := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		headers[0] + "\r\n" +
		headers[1] + "\r\n" +
		headers[2] + "\r\n\r\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(message))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Println("Successfully sended to " + to)
}
