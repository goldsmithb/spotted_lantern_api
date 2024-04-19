package core

import "github.com/google/uuid"

type User struct {
	UserId   uuid.UUID
	Username string
	Email    string
	Hash     string
	Score    int
}

// Defines the contract for the Lantern Fly API
type API interface {
	CheckUserExists(email string) bool

	GetAllKills() ([]int, error)
	GetKills(id string) (int, error)
}

type DbClient interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user User) error
	GetAllKills() ([]int, error)
	GetKillCount(userId string) (int, error)
	GetAllUsers() ([]User, error)
	GetHashForEmail(email string) (string, error)
}
