package repo

import "gorm.io/gorm"

type Repo interface {
	User() UserRepo
}

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repo {
	return repo{
		db: db,
	}
}

func (r repo) User() UserRepo {
	u := newUserRepo(r.db)
	return u
}
