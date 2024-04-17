package core

// Defines the contract for the Lantern Fly API
type API interface {
	GetAllKills() int
	GetKills(id string) int
}

type DbClient interface {
	GetAllKills() int
	GetKillCount(userId string) int
}
