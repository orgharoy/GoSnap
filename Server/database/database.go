package database

import (
	"log"

	"github.com/orgharoy/GoSnap/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
	var err error
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", "containers-us-west-185.railway.app", config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	dsn := "postgresql://postgres:oyG6OTJ9UshwXx8ucgsd@containers-us-west-185.railway.app:7642/railway"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	DB.AutoMigrate(&model.User{})

	log.Println("🚀 Connected Successfully to the Database")

	return nil
}
