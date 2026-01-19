package main

import (
	"fmt"
	"go_2601_04/internal/di"
	"go_2601_04/internal/infrastructure/config"
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

	r, err := di.InitializeApp(dsn)
	if err != nil {
		panic(err)
	}

	r.Run(":8080")
}
