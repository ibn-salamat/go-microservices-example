package repo

import (
	"users/internal/models"
	"users/internal/transport/http/handler/payload"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetByID(id int) (*models.User, error)
	Create(payload.CreateUserPayload) (*models.User, error)
}

type ur struct {
	db *gorm.DB
}

func newUserRepo(db *gorm.DB) ur {
	return ur{
		db: db,
	}
}

func (r ur) GetByID(id int) (*models.User, error) {
	var user models.User

	result := r.db.Model(&user).Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r ur) Create(payload payload.CreateUserPayload) (*models.User, error) {
	user := models.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	result := r.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
