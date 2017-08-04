package main

import (
	"github.com/alanfran/steampipe/protocol"
	"golang.org/x/net/context"
)

type QueryService struct {
	a *app
}

func (s *QueryService) Query(ctx context.Context, in *protocol.Address) (*protocol.Response, error) {
	r, err := s.a.QueryCache.Get(in.GetAddr())

	if err != nil {
		return &protocol.Response{}, err
	}

	return &protocol.Response{
		Name:       r.Name,
		Map:        r.Map,
		Game:       r.Game,
		Players:    int32(r.Players),
		MaxPlayers: int32(r.MaxPlayers),
		Bots:       int32(r.Bots),
	}, err
}
