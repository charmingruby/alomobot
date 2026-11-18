package core

import (
	"context"
	"math/rand"

	"github.com/charmingruby/alomobot/internal/pkmn"
)

var generations = []pkmn.Generation{
	pkmn.GenerationI,
	pkmn.GenerationII,
	pkmn.GenerationIII,
	pkmn.GenerationIV,
	pkmn.GenerationV,
	pkmn.GenerationVI,
	pkmn.GenerationVII,
	pkmn.GenerationVIII,
	pkmn.GenerationIX,
}

func FlipGeneration(_ context.Context) pkmn.Generation {
	return generations[rand.Intn(len(generations))]
}
