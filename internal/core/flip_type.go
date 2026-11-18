package core

import (
	"math/rand"

	"github.com/charmingruby/alomobot/internal/pkmn"
)

var types = []pkmn.Type{
	pkmn.TypeNormal,
	pkmn.TypeFire,
	pkmn.TypeWater,
	pkmn.TypeElectric,
	pkmn.TypeGrass,
	pkmn.TypeIce,
	pkmn.TypeFighting,
	pkmn.TypePoison,
	pkmn.TypeGround,
	pkmn.TypeFlying,
	pkmn.TypePsychic,
	pkmn.TypeBug,
	pkmn.TypeRock,
	pkmn.TypeGhost,
	pkmn.TypeDragon,
	pkmn.TypeDark,
	pkmn.TypeSteel,
	pkmn.TypeFairy,
}

func FlipType() pkmn.Type {
	return types[rand.Intn(len(types))]
}
