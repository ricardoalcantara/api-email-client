package emailengine

import (
	"bytes"
	"fmt"
	"text/template"

	h "github.com/matcornic/hermes/v2"
	gomail "gopkg.in/gomail.v2"
)

var dialer *gomail.Dialer
var hermes *h.Hermes

func Create() {
	dialer = gomail.NewDialer("localhost", 1025, "from@gmail.com", "<email_password>")
	// dialer = gomail.NewDialer("smtp.gmail.com", 587, "from@gmail.com", "<email_password>")

	hermes = &h.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: h.Product{
			// Appears in header & footer of e-mails
			Name: "Hermes",
			Link: "https://example-hermes.com/",
			// Optional product logo
			Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		},
	}
}

func GetTemplate(templateStr string, context any) string {
	t := template.Must(template.New("tmp").Parse(templateStr))

	var body bytes.Buffer
	err := t.Execute(&body, context)
	if err != nil {
		panic(err)
	}

	return body.String()
}

func GetHermes() (string, string) {
	email := h.Email{
		Body: h.Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Actions: []h.Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: h.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Confirm your account",
						Link:  "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := hermes.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// Generate the plaintext version of the e-mail (for clients that do not support xHTML)
	emailText, err := hermes.GeneratePlainText(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	return emailBody, emailText
}

func SendEmail(from string, to string, subject string, html_body string, text_body string) {
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
