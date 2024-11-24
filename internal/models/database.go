package models

import (
	"errors"
	"os"
	"strings"

	"github.com/ricardoalcantara/api-email-client/internal/utils"
	"github.com/rs/zerolog/log"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDataBase() {
	dbUrl := os.Getenv("DB_URL")
	envDialector := os.Getenv("DB_DIALECTOR")
	var err error
	var dialector gorm.Dialector
	switch strings.ToLower(envDialector) {
	case "sqlite":
		dialector = sqlite.Open(dbUrl)
	case "mysql":
		dialector = mysql.Open(dbUrl)
	case "postgres":
		dialector = postgres.Open(dbUrl)
	default:
		log.Fatal().Err(err).Msg("connection error:")
	}
	db, err = gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("connection error:")
	} else {
		log.Debug().Msg("Db Connected")
	}

	migrate()
	createAdmin()
}

func migrate() {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Smtp{})
	db.AutoMigrate(&Template{})
	db.AutoMigrate(&Email{})
	db.AutoMigrate(&ApiKey{})
}

func createAdmin() {

	if err := db.Take(&User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		email := utils.GetEnvOr("ADMIN_EMAIL", func() string {
			return utils.GenString(50)
		})
		password := utils.GetEnvOr("ADMIN_PASSWORD", func() string {
			return utils.GenString(100)
		})

		log.Debug().Msg("Admin Created")

		user := User{
			Name:  "Admin",
			Email: email,
		}
		err := user.SetPassword(password)
		if err != nil {
			log.Fatal().Err(err).Msg("Admin Password Error")
		}

		user.Save()
	}
}
