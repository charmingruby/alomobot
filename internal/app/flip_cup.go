package app

import (
	"context"
	"maps"
	"math/rand/v2"
	"slices"

	"github.com/charmingruby/alomobot/internal/core/cup"
)

func (a *App) FlipCup(_ context.Context) cup.Cup {
	cups := slices.Collect(maps.Values(cup.Cups))
	return cups[rand.IntN(len(cups))]
}
