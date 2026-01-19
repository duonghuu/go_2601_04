package mysql

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB, migrationsPath string) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	source, err := (&file.File{}).Open("file://" + migrationsPath)
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance(
		"file",
		source,
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(fmt.Errorf("migration failed: %w", err))
	}

	log.Println("✅ Database migrated successfully")
}
