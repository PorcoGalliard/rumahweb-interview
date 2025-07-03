package resource

import (
	"fmt"
	"log"

	"github.com/PorcoGalliard/rumahweb-interview/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgres(cfg config.PostgreConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Password,cfg.Name, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("❌ Failed connect to Postgres: %v", err)
	}

	log.Printf("✅ Successfully connect to Postgres with user: %s and port: %s\n", cfg.User, cfg.Port)
	return db
}