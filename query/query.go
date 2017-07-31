package query

import (
	servers "github.com/alanfran/SteamCondenserGo"
)

// Querier gets server information.
type Querier interface {
	Get(address string) servers.Response
}
