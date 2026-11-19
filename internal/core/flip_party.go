package core

import (
	"context"
	"math/rand"

	"github.com/charmingruby/alomobot/internal/pkmn"
	"github.com/charmingruby/alomobot/internal/pkmn/pokeapi"
)

const partySize = 6

func FlipParty(ctx context.Context, api *pokeapi.PokeAPI) ([]pkmn.Pokemon, error) {
	party := make([]pkmn.Pokemon, 0, partySize)
	unique := make(map[int]struct{}, partySize)

	for len(party) < 6 {
		p, err := api.GetPokemonByDexID(ctx, rand.Intn(pkmn.NationalDexCount))
		if err != nil {
			return nil, err
		}

		if _, ok := unique[p.ID]; ok {
			continue
		}

		party = append(party, p)
		unique[p.ID] = struct{}{}
	}

	return party, nil
}
