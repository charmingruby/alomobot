package core

import (
	"context"
	"math/rand"

	"github.com/charmingruby/alomobot/internal/pkmn"
)

func FlipGen(_ context.Context) pkmn.Gen {
	return pkmn.Gens[rand.Intn(len(pkmn.Gens))]
}
