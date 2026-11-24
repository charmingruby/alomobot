package cup

const MonogenCup NameCup = "Monogen"

var monogenCup = Cup{
	Name: MonogenCup,
	Guide: GuideCup{
		HowItWorks:   "All Pokémon on your team must come from the same generation.",
		Restrictions: "Do not mix Pokémon from different generations.",
		TeamBuilding: "Choose a generation and build your team using Pokémon introduced in it. For example, Bulbasaur and Pikachu are both from Generation I.",
	},
}
