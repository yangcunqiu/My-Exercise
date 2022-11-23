package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	url      string `yaml:"database.mysql.url"`
	username string `yaml:"database.mysql.username"`
	password string `yaml:"database.mysql.password"`
}

var db *gorm.DB

func initDatabase() {
	var err error
	db, err = gorm.Open(mysql.Open("1"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
