package core

import (
	"context"

	"github.com/charmingruby/alomobot/internal/pkmn"
	"github.com/charmingruby/alomobot/internal/pkmn/pokeapi"
)

const partySize = 6

type FlipPartyFilter struct {
	Bench int
}

func FlipParty(ctx context.Context, filter FlipPartyFilter) ([]pkmn.Pokemon, []pkmn.Pokemon, error) {
	api := pokeapi.PokeAPIFromContext(ctx)
	party := make([]pkmn.Pokemon, 0, partySize)
	bench := make([]pkmn.Pokemon, 0, max(0, filter.Bench))
	unique := make(map[int]struct{}, partySize)

	for len(party) < partySize || len(bench) < filter.Bench {
		p, err := api.GetPokemonByDexID(ctx, pkmn.RandomNationalDexID())
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
