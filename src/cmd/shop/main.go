package main

import (
	app "go_2601_04/internal/application/user"
	mysqlRepo "go_2601_04/internal/infrastructure/persistence/mysql"
	handler "go_2601_04/internal/interfaces/http"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "user:user_password@tcp(shop-db:3306)/my_database?charset=utf8mb4&parseTime=True&loc=Local"
	db := mysqlRepo.NewDatabase(dsn)

	mysqlRepo.RunMigrations(db, "migrations")

	repo := mysqlRepo.NewUserRepository(db)
	service := app.NewService(repo)
	userHandler := handler.NewUserHandler(service)

	r := gin.Default()
	userHandler.Register(r)

	r.Run(":8080")
}
