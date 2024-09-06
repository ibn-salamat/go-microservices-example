package postgres

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(pgDsn string) (*gorm.DB, error) {
	log.Println("Connecting to Postgres")
	db, err := gorm.Open(postgres.Open(pgDsn), &gorm.Config{})

	return db, err
}
