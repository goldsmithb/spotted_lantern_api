package api

// Implements core.API

type API struct {
}

func NewAPI() *API {
	return &API{}
}

func (a *API) GetAllKills() int {
	return 3000
}

func (a *API) GetKills(_ string) int {
	return 1
}
