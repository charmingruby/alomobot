package core

import (
	"context"

	"github.com/charmingruby/alomobot/internal/pkmn"
	"github.com/charmingruby/alomobot/internal/pkmn/pokeapi"
)

func FlipPokemon(ctx context.Context) (pkmn.Pokemon, error) {
	api := pokeapi.PokeAPIFromContext(ctx)

	return api.GetPokemonByDexID(ctx, pkmn.RandomNationalDexID())
}
