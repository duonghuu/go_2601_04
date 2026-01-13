package main

import (
	"go_2601_04/internal/application/message"
	model "go_2601_04/internal/infrastructure/persistence/mysql"
	"go_2601_04/internal/interface/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 1. Setup Database
	dsn := "user:user_password@tcp(shop-db)/my_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// 2. Auto Migrate model Infrastructure
	db.AutoMigrate(&model.Message{})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	repo := model.NewMessageRepository(db)

	appSvc := message.NewService(repo)

	handler := http.NewHandler(appSvc)

	r := gin.Default()
	r.GET("/hello", handler.HelloWorld)

	r.Run(":8080")
}
