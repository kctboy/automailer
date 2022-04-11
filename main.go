package main

import (
	"flag"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"

	"github.com/kctboy/automailer/lib"
	"github.com/scorredoira/email"
)

var file string
var adressTo string
var mailTexts string

func init() {
	flag.StringVar(&file, "f", "none", "the file")
	flag.StringVar(&mailTexts, "t", "Hello, this is the default message", "the text massage")
	flag.StringVar(&adressTo, "e", "bosboomhut@gmail.com", "The email adress where it wil be send to")
	flag.Parse()
}

func SendMail(config lib.Config) error {
	// compose the message
	m := email.NewMessage("", mailTexts)
	m.From = mail.Address{Name: "Golang Application", Address: adressTo}
	m.To = []string{config.Email.To}

	// add attachments
	if file != "none" {
		if err := m.Attach(file); err != nil {
			return err
		}
	}

	// add headers
	m.AddHeader("Subject", "output from automailer")

	// send it
	auth := smtp.PlainAuth("", config.Email.From, config.Email.SmtpPassword, config.Email.Mailserver)
	if err := email.Send(config.Email.Mailserver+":"+config.Email.Mailport, auth, m); err != nil {
		return err
	}

	return nil
}

func main() {
	config, err := lib.LoadConfiguration("config.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config loaded")

	fmt.Println(SendMail(config))
}
