package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/term"

	"github.com/matcornic/hermes"
)

var (
	errEmptyServerConfig   = errors.New("SMTP server config is empty")
	errEmptyPort           = errors.New("SMTP port config is empty")
	errEmptyUser           = errors.New("SMTP user is empty")
	errEmptySenderIdentity = errors.New("SMTP sender identity is empty")
	errEmptySenderEmail    = errors.New("SMTP sender email is empty")
	errEmptyReceiverEmails = errors.New("no receiver emails configured")
)

type example interface {
	Email() hermes.Email
	Name() string
}

func main() {

	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "Hermes",
			Link: "https://example-hermes.com/",
			Logo: "https://github.com/matcornic/hermes/blob/master/examples/gopher.png?raw=true",
		},
	}
	sendEmails := os.Getenv("HERMES_SEND_EMAILS") == "true"

	examples := []example{
		new(welcome),
		new(reset),
		new(receipt),
		new(maintenance),
		new(inviteCode),
	}

	themes := []hermes.Theme{
		new(hermes.Default),
		new(hermes.Flat),
	}

	// Generate emails
	for _, theme := range themes {
		h.Theme = theme
		for _, e := range examples {
			generateEmails(h, e.Email(), e.Name())
		}
	}

	// Send emails only when requested
	if sendEmails {
		port, _ := strconv.Atoi(os.Getenv("HERMES_SMTP_PORT"))
		password := os.Getenv("HERMES_SMTP_PASSWORD")
		SMTPUser := os.Getenv("HERMES_SMTP_USER")
		if password == "" {
			log.Printf("Enter SMTP password of '%s' account: ", SMTPUser)
			bytePassword, _ := term.ReadPassword(0)
			password = string(bytePassword)
		}
		smtpConfig := smtpAuthentication{
			Server:         os.Getenv("HERMES_SMTP_SERVER"),
			Port:           port,
			SenderEmail:    os.Getenv("HERMES_SENDER_EMAIL"),
			SenderIdentity: os.Getenv("HERMES_SENDER_IDENTITY"),
			SMTPPassword:   password,
			SMTPUser:       SMTPUser,
		}
		options := sendOptions{
			To: os.Getenv("HERMES_TO"),
		}
		for _, theme := range themes {
			h.Theme = theme
			for _, e := range examples {
				options.Subject = "Hermes | " + h.Theme.Name() + " | " + e.Name()
				log.Printf("Sending email '%s'...\n", options.Subject)
				htmlBytes, err := os.ReadFile(fmt.Sprintf("%v/%v.%v.html", h.Theme.Name(), h.Theme.Name(), e.Name()))
				if err != nil {
					panic(err)
				}
				txtBytes, err := os.ReadFile(fmt.Sprintf("%v/%v.%v.txt", h.Theme.Name(), h.Theme.Name(), e.Name()))
				if err != nil {
					panic(err)
				}
				err = send(smtpConfig, options, string(htmlBytes), string(txtBytes))
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func generateEmails(h hermes.Hermes, email hermes.Email, example string) {
	_ = "STUB: not implemented"
	// Generate the HTML template and save it
	return
}

// Generate the plaintext template and save it

type smtpAuthentication struct {
	Server         string
	Port           int
	SenderEmail    string
	SenderIdentity string
	SMTPUser       string
	SMTPPassword   string
}

// sendOptions are options for sending an email
type sendOptions struct {
	To      string
	Subject string
}

// send sends the email
func send(smtpConfig smtpAuthentication, options sendOptions, htmlBody string, txtBody string) error {
	_ = "STUB: not implemented"
	return nil
}
