package models

import (
	"errors"
	"fmt"
	"os"

	"github.com/ricardoalcantara/api-email-client/internal/utils"
	log "github.com/sirupsen/logrus"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDataBase() {
	db_url := os.Getenv("DB_URL")
	var err error
	db, err = gorm.Open(sqlite.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatal("connection error:", err)
	} else {
		log.Debug("Db Connected")
	}

	migrate()
	createAdmin()
	createTemplate()
}

func migrate() {
	db.AutoMigrate(&Client{})
	db.AutoMigrate(&Smtp{})
	db.AutoMigrate(&Template{})
}

func createAdmin() {
	var clientId string
	var clientSecret string

	if value, ok := os.LookupEnv("CLIENT_ID"); ok {
		clientId = value
	} else {
		clientId = utils.GenString(50)
	}

	if value, ok := os.LookupEnv("CLIENT_SECRET"); ok {
		clientSecret = value
	} else {
		clientSecret = utils.GenString(50)
	}

	if err := db.First(&Client{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		log.Debug("Admin Created")

		client := Client{
			Name:         "Admin",
			ClientId:     clientId,
			ClientSecret: clientSecret,
		}

		client.Save()

		fmt.Printf("ClientId: %s\n", client.ClientId)
		fmt.Printf("ClientSecret: %s\n", client.ClientSecret)
	}
}

func createTemplate() {
	if err := db.First(&Template{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		t := Template{
			Name:         "Default",
			TemplateHtml: "<h1>{{.Name}}</h1>",
			TemplateText: "{{.Name}}",
			JsonSchema:   "{ Name: string }",
		}

		t.Save()
	}
}
