package mysql

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	return db
}
