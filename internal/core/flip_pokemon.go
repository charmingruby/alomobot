package core

import (
	"context"
	"math/rand"

	"github.com/charmingruby/alomobot/internal/pkmn"
	"github.com/charmingruby/alomobot/internal/pkmn/pokeapi"
)

func FlipPokemon(ctx context.Context, api *pokeapi.PokeAPI) (pkmn.Pokemon, error) {
	return api.GetPokemonByDexID(ctx, rand.Intn(pkmn.NationalDexCount))
}
