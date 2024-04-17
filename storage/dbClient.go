package storage

// implements core.dbClient
type dbClient struct {
}

func NewDbClient() *dbClient {
	return &dbClient{}
}

func (db *dbClient) GetAllKills() int {
	return 100
}

func (db *dbClient) GetKillCount(userId string) int {
	return 10
}
