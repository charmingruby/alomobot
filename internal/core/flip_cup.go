package core

import (
	"context"
	"math/rand"

	"github.com/charmingruby/alomobot/internal/cup"
)

func FlipCup(_ context.Context) cup.Cup {
	return cup.Cups[rand.Intn(len(cup.Cups))]
}
