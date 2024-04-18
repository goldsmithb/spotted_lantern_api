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
	GetAllKills() ([]int, error)
	GetKills(id string) (int, error)
}

type DbClient interface {
	GetAllKills() ([]int, error)
	GetKillCount(userId string) (int, error)
}
