package core

import (
	"context"
	"math/rand"

	"github.com/charmingruby/alomobot/internal/pkmn"
)

var cups = []pkmn.Cup{
	pkmn.CupAsh,
	pkmn.CupLittle,
	pkmn.CupDexLastNumber,
	pkmn.CupCharactersAmount,
	pkmn.CupRental,
	pkmn.CupGeneration,
	pkmn.CupMonotype,
	pkmn.CupRandomParty,
	pkmn.CupBooster,
	pkmn.CupNoEvo,
	pkmn.CupSingleType,
}

func FlipCup(_ context.Context) pkmn.Cup {
	return cups[rand.Intn(len(cups))]
}
