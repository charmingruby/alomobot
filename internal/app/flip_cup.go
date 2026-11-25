package app

import (
	"context"
	"math/rand"

	"github.com/charmingruby/alomobot/internal/core/cup"
)

func (a *App) FlipCup(_ context.Context) cup.Cup {
	return cup.Cups[rand.Intn(len(cup.Cups))]
}
