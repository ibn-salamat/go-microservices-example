package models

type User struct {
	Id     string `bson:"_id,omitempty"`
	UserId int    `bson:"user_id"`
}
