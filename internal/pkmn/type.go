package pkmn

import "math/rand"

type Type string

const (
	TypeNormal   Type = "normal"
	TypeFire     Type = "fire"
	TypeWater    Type = "water"
	TypeElectric Type = "electric"
	TypeGrass    Type = "grass"
	TypeIce      Type = "ice"
	TypeFighting Type = "fighting"
	TypePoison   Type = "poison"
	TypeGround   Type = "ground"
	TypeFlying   Type = "flying"
	TypePsychic  Type = "psychic"
	TypeBug      Type = "bug"
	TypeRock     Type = "rock"
	TypeGhost    Type = "ghost"
	TypeDragon   Type = "dragon"
	TypeDark     Type = "dark"
	TypeSteel    Type = "steel"
	TypeFairy    Type = "fairy"
)

var types = []Type{
	TypeNormal,
	TypeFire,
	TypeWater,
	TypeElectric,
	TypeGrass,
	TypeIce,
	TypeFighting,
	TypePoison,
	TypeGround,
	TypeFlying,
	TypePsychic,
	TypeBug,
	TypeRock,
	TypeGhost,
	TypeDragon,
	TypeDark,
	TypeSteel,
	TypeFairy,
}

func RandomType() Type {
	return types[rand.Intn(len(types))]
}
