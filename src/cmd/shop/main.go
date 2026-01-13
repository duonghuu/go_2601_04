package main

import (
	"fmt"
	"go_2601_04/internal/application/message"
	"go_2601_04/internal/infrastructure/config"
	model "go_2601_04/internal/infrastructure/persistence/mysql"
	"go_2601_04/internal/interface/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()
	// Setup Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Auto Migrate model Infrastructure
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
