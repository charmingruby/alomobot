package cup

const ShinyCup CupName = "Shiny"

var shinyCup = Cup{
	Name: ShinyCup,
	Guide: GuideCup{
		HowItWorks:   "Only shiny Pokémon can participate.",
		Restrictions: "Non-shiny Pokémon are not allowed. No other cup-specific restrictions apply.",
		TeamBuilding: "Build a team using only shiny Pokémon, following the general tournament rules.",
	},
}
