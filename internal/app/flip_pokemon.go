package app

import (
	"context"
	"strconv"

	"github.com/charmingruby/alomobot/internal/core/pkmn"
)

func (a *App) FlipPokemon(ctx context.Context) (pkmn.Pokemon, error) {
	return a.api.GetPokemon(ctx, strconv.Itoa(pkmn.RandomNationalDexID()))
}
