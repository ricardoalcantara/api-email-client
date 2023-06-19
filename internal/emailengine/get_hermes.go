package emailengine

import (
	h "github.com/matcornic/hermes/v2"
	"github.com/ricardoalcantara/api-email-client/internal/utils"
	"github.com/sirupsen/logrus"
)

func GetHermes(body any) (string, string) {
	data, err := utils.TypeConverter[h.Body](&body)
	if err != nil {
		logrus.Println(err.Error())
	}
	email := h.Email{
		Body: *data,
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

// utils.TypeConverter
