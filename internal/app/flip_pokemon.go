package app

import (
	"context"

	"github.com/charmingruby/alomobot/internal/core/pkmn"
)

func (a *App) FlipPokemon(ctx context.Context) (pkmn.Pokemon, error) {
	return a.api.GetPokemonByDexID(ctx, pkmn.RandomNationalDexID())
}
