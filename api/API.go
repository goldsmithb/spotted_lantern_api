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

func (a *api) GetAllKills() ([]int, error) {
	return a.dbClient.GetAllKills()
}

func (a *api) GetKills(id string) (int, error) {
	return a.dbClient.GetKillCount(id)
}
