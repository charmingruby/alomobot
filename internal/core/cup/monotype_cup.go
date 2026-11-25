package cup

const MonotypeCup CupName = "Monotype"

var monotypeCup = Cup{
	Name: MonotypeCup,
	Guide: GuideCup{
		HowItWorks:   "All Pokémon on your team must share at least one type.",
		Restrictions: "Dual-type Pokémon are allowed as long as they share the team's chosen type.",
		TeamBuilding: "Use `/flip type` to draw your team's type, then select Pokémon that have it. For example, if Water is drawn, Squirtle and Gyarados are both allowed.",
	},
}
