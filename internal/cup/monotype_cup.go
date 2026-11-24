package cup

const MonotypeCup NameCup = "Monotype"

var monotypeCup = Cup{
	Name: MonotypeCup,
	Guide: GuideCup{
		HowItWorks:   "All Pokémon on your team must share at least one type.",
		Restrictions: "Dual-type Pokémon are allowed as long as they share the team's chosen type.",
		TeamBuilding: "Choose a type and select Pokémon that have it. For example, Squirtle and Gyarados both share the Water type.",
	},
}
