package cup

const LittleCup NameCup = "Little"

var littleCup = Cup{
	Name: LittleCup,
	Guide: GuideCup{
		HowItWorks:   "Battles use only base-form Pokémon that can evolve.",
		Restrictions: "Only the first stage of an evolutionary line is allowed. In a three-stage line, neither the middle nor the final stage is allowed. Pokémon that cannot evolve are also not allowed.",
		TeamBuilding: "Choose base-form Pokémon with an evolution. For example, Bulbasaur is allowed, but Ivysaur and Venusaur are not.",
	},
}
