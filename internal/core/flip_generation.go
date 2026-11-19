package core

import (
	"context"
	"math/rand"

	"github.com/charmingruby/alomobot/internal/pkmn"
)

func FlipGeneration(_ context.Context) pkmn.Generation {
	return pkmn.Generations[rand.Intn(len(pkmn.Generations))]
}
