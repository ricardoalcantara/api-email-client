package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	hermes "github.com/matcornic/hermes/v2"
	"github.com/ricardoalcantara/api-email-client/pkg/client"
	"github.com/ricardoalcantara/api-email-client/pkg/types"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
	log.Logger = log.
		With().
		Caller().
		Logger().
		Output(zerolog.ConsoleWriter{Out: os.Stderr})

	err := godotenv.Load()
	if err != nil {
		log.Debug().Msg("Fail loading .env file")
	}
}

func main() {

	apiClient := client.New(client.WithBaseURL("http://localhost:5555"))

	getTokenResponse, err := apiClient.GetToken(os.Getenv("ADMIN_EMAIL"), os.Getenv("ADMIN_PASSWORD"))
	if err != nil {
		panic(err)
	}
	if getTokenResponse.Error != nil {
		panic(getTokenResponse.Error)
	}

	apiClient = client.New(client.WithBaseURL("http://localhost:5555"), client.WithToken(getTokenResponse.Result.AccessToken))
	createApikeyResponse, err := apiClient.CreateAPIKey(&types.CreateApiKeyDto{
		Name:        "Teste",
		IpWhitelist: "127.0.0.1",
	})
	if err != nil {
		panic(err)
	}
	if createApikeyResponse.Error != nil {
		panic(createApikeyResponse.Error)
	}

	apiClient = client.New(client.WithBaseURL("http://localhost:5555"), client.WithApiKey(createApikeyResponse.Result.Key))
	listTemplatesResponse, err := apiClient.ListTemplates(1)
	if err != nil {
		panic(err)
	}
	if listTemplatesResponse.Error != nil {
		panic(listTemplatesResponse.Error)
	}
	fmt.Println(listTemplatesResponse.Result)
}

func gen() {

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
