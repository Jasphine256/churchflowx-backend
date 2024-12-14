package database

import (
	"churchflowx/config"
	// "churchflowx/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitialiseDB() *gorm.DB {

	// PostgreSQL DSN
	config.LoadEnv()
	dsn := config.GetEnv("DATABASE_URL", "")

	// Initialize GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Database connected!")

	// db.AutoMigrate(
	// 	&models.Admin{},
	// 	&models.Member{},
	// 	&models.Minister{},
	// 	&models.Pastor{},
	// 	&models.Visitor{},
	// 	&models.Task{},
	// 	&models.Plan{},
	// 	&models.Project{},
	// 	&models.Fund{},
	// 	&models.Payment{},
	// )

	return db
}
