package main

import (
	"fmt"
	"os"

	hermes "github.com/matcornic/hermes/v2"
)

func main() {
	name := "forgot-password"
	h := hermes.Hermes{
		Product: hermes.Product{
			Name:        "OnePanel",
			Link:        "https://www.onepanel.com.br",
			Logo:        "https://www.pscanary.com/logo.svg",
			Copyright:   "Copyright © 2023 One Panel. Todos os Direitos Reservados",
			TroubleText: "Se tiver problemas com o botão '{ACTION}', copie o código e cole na página de recuperação de senha ou no link a seguir.",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Greeting: "Olá, ",
			Name:     "{{ .Name }}",
			Intros: []string{
				"Você recebeu esse email porque foi solicitado a recuperação de senha dessa conta.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Para recuperar clica no botão ou copie o código e cole na página de recuperação de senha.",
					InviteCode:   "{{ .Code }}",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Recuperar agora",
						Link:  "{{ .Link }}",
					},
				},
			},
			Outros: []string{
				"Se não foi você quem solicitou, pode desconsiderar esse email.",
			},
			Signature: "Atenciosamente",
		},
	}

	emailHtml, err := h.GenerateHTML(email)
	if err != nil {
		panic(err)
	}

	emailText, err := h.GeneratePlainText(email)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fmt.Sprintf("tmp/%s.html", name), []byte(emailHtml), 0644)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fmt.Sprintf("tmp/%s.txt", name), []byte(emailText), 0644)
	if err != nil {
		panic(err)
	}
}
