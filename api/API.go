package api

import (
	"github.com/goldsmithb/spotted_lantern_api/core"
	"github.com/goldsmithb/spotted_lantern_api/storage"
)

// Implements core.API
type api struct {
	dbClient core.DbClient
}

func NewAPI() *api {
	return &api{
		dbClient: storage.NewDbClient(),
	}
}

func (a *api) GetKills(id string) int {
	return a.dbClient.GetKillCount(id)
}

func (a *api) GetAllKills() int {
	return a.dbClient.GetAllKills()
}
