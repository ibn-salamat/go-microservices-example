package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(pgDsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(pgDsn), &gorm.Config{})

	return db, err
}
