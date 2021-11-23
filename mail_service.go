package main

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func sendMail(subject string, content string, emails []string) {
	d := gomail.NewDialer("smtp.gmail.com", 587, "gavinsignaturebox@gmail.com", "tethics2019")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", "gavinsignaturebox@gmail.com")
	m.SetHeader("To", emails...)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
