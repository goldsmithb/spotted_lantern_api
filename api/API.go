package api

import (
	"github.com/goldsmithb/spotted_lantern_api/config"
	"github.com/goldsmithb/spotted_lantern_api/core"
)

// Implements core.API
type api struct {
	config   *config.Config
	dbClient core.DbClient
}

func NewAPI(conf *config.Config, db core.DbClient) *api {
	return &api{
		config:   conf,
		dbClient: db,
	}
}

func (a *api) GetKills(id string) int {
	return a.dbClient.GetKillCount(id)
}

func (a *api) GetAllKills() int {
	return a.dbClient.GetAllKills()
}
