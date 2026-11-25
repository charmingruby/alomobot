package app

import (
	"context"
	"math/rand"

	"github.com/charmingruby/alomobot/internal/core/pkmn"
)

func (a *App) FlipGen(_ context.Context) pkmn.Gen {
	return pkmn.Gens[rand.Intn(len(pkmn.Gens))]
}
