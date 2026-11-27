package cup

const SingleTypeCup CupName = "Single Type"

var singleTypeCup = Cup{
	Name: SingleTypeCup,
	Guide: GuideCup{
		HowItWorks:   "Only Pokémon with exactly one type can participate.",
		Restrictions: "Dual-type Pokémon are not allowed.",
		TeamBuilding: "Choose Pokémon with a single type. For example, Pikachu is allowed because it is only Electric, but Bulbasaur is not because it is Grass/Poison.",
	},
}
