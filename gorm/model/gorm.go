package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var Conn *gorm.DB

func init() {
	host := "localhost"
	port := 4432
	user := "postgres"
	password := "Oushu6@China"
	database := "test_gorm"

	// data source name
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, database, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("database connected failed ", err)
	}

	Conn = db

	// 创建表
	err = Conn.AutoMigrate(&User{})
	if err != nil {
		log.Fatalln("database auto migrate failed ", err)
	}
}
