package cup

const MonogenCup NameCup = "Monogen"

var monogenCup = Cup{
	Name: MonogenCup,
	Guide: GuideCup{
		HowItWorks:   "All Pokémon on your team must come from the same generation.",
		Restrictions: "Do not mix Pokémon from different generations.",
		TeamBuilding: "Use `/flip gen` to draw your team's generation, then build your team using Pokémon introduced in it. For example, if Generation I is drawn, Bulbasaur and Pikachu are both allowed.",
	},
}
