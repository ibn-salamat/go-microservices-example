package repo

import "fmt"

type UserRepo interface {
	Get() error
}

type user struct{}

func (r user) Get() error {
	fmt.Println("get from repo")
	return nil
}
