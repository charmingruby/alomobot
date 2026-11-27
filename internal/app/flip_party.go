package app

import (
	"context"
	"strconv"

	"github.com/charmingruby/alomobot/internal/core/pkmn"
)

const partySize = 6

type FlipPartyFilter struct {
	Bench int
}

func (a *App) FlipParty(ctx context.Context, filter FlipPartyFilter) ([]pkmn.Pokemon, []pkmn.Pokemon, error) {
	party := make([]pkmn.Pokemon, 0, partySize)
	bench := make([]pkmn.Pokemon, 0, max(0, filter.Bench))
	unique := make(map[int]struct{}, partySize)

	for len(party) < partySize || len(bench) < filter.Bench {
		p, err := a.api.GetPokemon(ctx, strconv.Itoa(pkmn.RandomNationalDexID()))
		if err != nil {
			return nil, nil, err
		}

		if _, ok := unique[p.ID]; ok {
			continue
		}

		if len(party) < partySize {
			party = append(party, p)
		} else {
			bench = append(bench, p)
		}

		unique[p.ID] = struct{}{}
	}

	return party, bench, nil
}
