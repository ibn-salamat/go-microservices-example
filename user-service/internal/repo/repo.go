package repo

type Repo interface {
	User() UserRepo
}

type repo struct {
}

func New() Repo {
	return repo{}
}

func (r repo) User() UserRepo {
	u := user{}
	return u
}
