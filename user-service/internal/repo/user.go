package repo

import (
	"fmt"

	"gorm.io/gorm"
)

type UserRepo interface {
	Get() error
}

type user struct {
	db *gorm.DB
}

func newUserRepo(db *gorm.DB) user {
	return user{
		db: db,
	}
}

func (r user) Get() error {
	fmt.Println("get from repo")
	return nil
}
