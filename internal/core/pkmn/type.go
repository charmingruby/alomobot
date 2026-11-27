package pkmn

import (
	"fmt"

	"github.com/charmingruby/alomobot/internal/format"
)

type Type string

const (
	TypeNormal   Type = "Normal"
	TypeFire     Type = "Fire"
	TypeWater    Type = "Water"
	TypeElectric Type = "Electric"
	TypeGrass    Type = "Grass"
	TypeIce      Type = "Ice"
	TypeFighting Type = "Fighting"
	TypePoison   Type = "Poison"
	TypeGround   Type = "Ground"
	TypeFlying   Type = "Flying"
	TypePsychic  Type = "Psychic"
	TypeBug      Type = "Bug"
	TypeRock     Type = "Rock"
	TypeGhost    Type = "Ghost"
	TypeDragon   Type = "Dragon"
	TypeDark     Type = "Dark"
	TypeSteel    Type = "Steel"
	TypeFairy    Type = "Fairy"
)

var Types = []Type{
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

func (t Type) String() string {
	return string(t)
}

func (t Type) Render() string {
	return fmt.Sprintf("`%s`", format.Capitalize(t.String()))
}
