package core

import (
	"context"
	"math/rand"

	"github.com/charmingruby/alomobot/internal/pkmn"
)

func FlipType(_ context.Context) pkmn.Type {
	return pkmn.Types[rand.Intn(len(pkmn.Types))]
}
