package core

import (
	"context"
	"math/rand"

	"github.com/charmingruby/alomobot/internal/pkmn"
)

func FlipRegion(_ context.Context) pkmn.Region {
	return pkmn.Regions[rand.Intn(len(pkmn.Regions))]
}
