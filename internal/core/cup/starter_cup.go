package cup

const StarterCup CupName = "Starter"

var starterCup = Cup{
	Name: StarterCup,
	Guide: GuideCup{
		HowItWorks:   "Battles use starter Pokémon and their evolutions.",
		Restrictions: "Only Pokémon from starter evolutionary lines are allowed.",
		TeamBuilding: "Choose Pokémon from any stage of a starter's evolutionary line. For example, Charmander, Charmeleon, and Charizard are all allowed.",
	},
}
