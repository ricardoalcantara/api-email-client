package emailengine

import (
	"fmt"
	"os"
	"time"

	h "github.com/matcornic/hermes/v2"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
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
}

func SendEmail(dialer *gomail.Dialer, from string, to string, subject string, html_body string, text_body string) {
	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	if len(text_body) > 0 {
		m.SetBody("text/plain", text_body)
	}
	if len(html_body) > 0 {
		m.AddAlternative("text/html", html_body)
	}

	if err := dialer.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
