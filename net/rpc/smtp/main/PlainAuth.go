package main

import (
	"net/smtp"
	"log"
)

func main() {
	// hostname is used by PlainAuth to validate the TLS certificate.
	hostname := "mail.foxmail.com"
	auth := smtp.PlainAuth("", "bjzhang03@foxmail.com", "password", hostname)

	msg := []byte("test");
	from := "bjzhang03";
	recipients := []string{"bjzhang03", "123"};
	err := smtp.SendMail(hostname+":25", auth, from, recipients, msg)
	if err != nil {
		log.Fatal(err)
	}

}
