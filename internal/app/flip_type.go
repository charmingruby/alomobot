package app

import (
	"context"
	"math/rand"

	"github.com/charmingruby/alomobot/internal/core/pkmn"
)

func (a *App) FlipType(_ context.Context) pkmn.Type {
	return pkmn.Types[rand.Intn(len(pkmn.Types))]
}
