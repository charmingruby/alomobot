package pkmn

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

func (t Type) String() string {
	return string(t)
}
