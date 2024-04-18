package core

type User struct {
	Id       string
	Username string
	Email    string
	Passkey  string
	Score    int
}

// Defines the contract for the Lantern Fly API
type API interface {
	GetAllKills() int
	GetKills(id string) int
}

type DbClient interface {
	GetAllKills() int
	GetKillCount(userId string) int
}
