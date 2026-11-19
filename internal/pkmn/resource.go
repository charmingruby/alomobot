package pkmn

type Resource string

const (
	ResourceType       Resource = "type"
	ResourceCup        Resource = "cup"
	ResourcePokemon    Resource = "pokemon"
	ResourceParty      Resource = "party"
	ResourceGeneration Resource = "generation"
	ResourceRegion     Resource = "region"
)

func (k Resource) Valid() bool {
	switch k {
	case ResourceType, ResourceCup, ResourcePokemon, ResourceParty, ResourceGeneration, ResourceRegion:
		return true
	default:
		return false
	}
}

func (k Resource) String() string {
	return string(k)
}
