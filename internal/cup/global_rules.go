package cup

type Rule string

func (r Rule) Render() string {
	return string(r)
}

var GlobalRules = []Rule{
	"Only one battle mechanic is allowed per team: Mega Evolution, Terastallization, Gigantamax, or Z-Moves. For example, a team with Mega Evolution cannot also use Terastallization.",
	"Consumable items such as Potions and Revives cannot be used.",
	"One-hit knockout moves, such as Fissure and Guillotine, are banned.",
	"Moves that increase the user's evasion, such as Minimize, are banned.",
	"Held items cannot be repeated within a team. Each Pokémon's held item must be unique.",
	"There is no limit on the number of Legendary Pokémon, provided none are classified as Uber or Uber UU.",
}
