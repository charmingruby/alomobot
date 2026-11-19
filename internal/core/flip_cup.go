package core

import (
	"context"
	"math/rand"

	"github.com/charmingruby/alomobot/internal/pkmn"
)

func FlipCup(_ context.Context) pkmn.Cup {
	return pkmn.Cups[rand.Intn(len(pkmn.Cups))]
}
