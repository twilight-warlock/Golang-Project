package main

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func sendMail() {
	d := gomail.NewDialer("smtp.gmail.com", 587, "gavinsignaturebox@gmail.com", "tethics2019")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", "gavinsignaturebox@gmail.com")
	m.SetHeader("To", "neelansh.m@somaiya.edu", "gavinsignaturebox@gmail.com")
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Richard</b> and <i>Jared</i>!")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
