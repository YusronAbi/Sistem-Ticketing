package config

import (
	"SISTEM-TICKETING/entity"
	"log"
)

func MigrateDB() {
	err := DB.AutoMigrate(
		&entity.Event{},
		&entity.Ticket{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed!")
}
