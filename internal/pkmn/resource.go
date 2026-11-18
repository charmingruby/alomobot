package pkmn

type Resource string

const (
	ResourceType       Resource = "type"
	ResourceCup        Resource = "cup"
	ResourcePokemon    Resource = "pokemon"
	ResourceTeam       Resource = "team"
	ResourceGeneration Resource = "generation"
	ResourceRegion     Resource = "region"
)

func (k Resource) Valid() bool {
	switch k {
	case ResourceType, ResourceCup, ResourcePokemon, ResourceTeam, ResourceGeneration, ResourceRegion:
		return true
	default:
		return false
	}
}

func (k Resource) String() string {
	return string(k)
}
