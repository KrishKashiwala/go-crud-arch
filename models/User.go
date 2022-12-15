package models

import "gorm.io/gorm"

type User struct {
	Id       int    `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
