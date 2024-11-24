package emailengine

import (
	"errors"
	"fmt"
	"os"
	"time"

	h "github.com/matcornic/hermes/v2"
	"github.com/ricardoalcantara/api-email-client/internal/models"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
	"github.com/sirupsen/logrus"
	gomail "gopkg.in/gomail.v2"
)

var hermes *h.Hermes

func Create() {
	hermes = &h.Hermes{
		Product: h.Product{
			Name: utils.GetEnv("APP_NAME", "Api Email Client"),
			Link: utils.GetEnv("APP_LINK", "https://github.com/ricardoalcantara/api-email-client"),
			Logo: utils.GetEnv("APP_LOGO", "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png"),
			Copyright: utils.GetEnvOr("APP_COPYRIGHT", func() string {
				year := time.Now().Year()
				return fmt.Sprintf("Copyright © %d Api Email Client. All rights reserved. ", year)
			}),
		},
	}

	if value, ok := os.LookupEnv("HERMES_THEME"); ok {
		if value == "Flat" {
			hermes.Theme = h.Theme(&h.Flat{})
		}
	}

	if value, ok := os.LookupEnv("APP_TROUBLE_TEXT"); ok {
		hermes.Product.TroubleText = value
	}

	if emailChan == nil {
		emailChan = make(chan models.Email)
		go worker(emailChan)
	}
}

var emailChan chan models.Email

func worker(emailChan <-chan models.Email) {
	for email := range emailChan {
		var err error
		var d *gomail.Dialer

		if d, err = email.Smtp.GetDialer(); err != nil {
			utils.PrintError(err)
			continue
		}
		if err = SendEmail(d, email.Smtp.Email, email.Smtp.Name, email.To, email.Subject, email.HtmlBody, email.TextBody); err != nil {
			utils.PrintError(err)
			continue
		}

		if err = models.EmailUpdateSent(email.ID); err != nil {
			utils.PrintError(err)
			continue
		}

		logrus.WithFields(logrus.Fields{
			"id":      email.ID,
			"email":   email.Smtp.Email,
			"to":      email.To,
			"subject": email.Subject,
		}).Info("Email Sent")
	}
}

func SendEmailQueue(email models.Email) error {
	select {
	case emailChan <- email:
		return nil
	default:
		return errors.New("max capacity reached")
	}
}

func SendEmail(dialer *gomail.Dialer, from string, fromName string, to string, subject string, htmlBody string, textBody string) error {
	m := gomail.NewMessage()

	m.SetAddressHeader("From", from, fromName)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	if len(textBody) > 0 {
		m.SetBody("text/plain", textBody)
	}
	if len(htmlBody) > 0 {
		m.AddAlternative("text/html", htmlBody)
	}

	return dialer.DialAndSend(m)
}
