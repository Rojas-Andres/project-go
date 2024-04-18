package modelos

import (
	"time"
)
type User struct {
	Id       int
	Name string
	CretedAt time.Time
	Status bool
}


func (user *User) AddUser(id int, name string, createdAt time.Time, status  bool) {
	user.Status = status
	user.Id = id
	user.Name = name
	user.CretedAt = createdAt
}