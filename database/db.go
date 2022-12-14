package db

import (
	"fmt"
	"log"
	"os"

	"github.com/KrishKashiwala/go-crud-arch/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

func Setup() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	DSN := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s sslmode=%s", config.Host, config.User, config.Port, config.Password, config.DBName, config.SSLMode)
	log.Println("Connecting to database...")
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	e := models.MigrateUsers(DB)
	if e != nil {
		log.Println("users didn't migrate")
	}
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)

	}
	log.Println("Connected to ", DB)
	// turned on the loger on info mode
	DB.Logger = logger.Default.LogMode(logger.Info)
}

func GetDB() *gorm.DB {
	return DB
}
