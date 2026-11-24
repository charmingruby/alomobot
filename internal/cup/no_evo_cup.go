package cup

const NoEvoCup NameCup = "No Evo"

var noEvoCup = Cup{
	Name: NoEvoCup,
	Guide: GuideCup{
		HowItWorks:   "Only Pokémon that are not part of an evolutionary line can participate.",
		Restrictions: "Pokémon that can evolve or evolve from another Pokémon are not allowed. Fully evolved Pokémon from evolutionary lines are not eligible.",
		TeamBuilding: "Choose Pokémon with neither pre-evolutions nor evolutions. For example, Lapras and Ditto are allowed, but Charizard is not.",
	},
}
