package cup

const RandomPartyCup NameCup = "Random Party"

var randomPartyCup = Cup{
	Name: RandomPartyCup,
	Guide: GuideCup{
		HowItWorks:   "You receive six randomly drawn Pokémon for your team and six more for your bench.",
		Restrictions: "Build your battle team using only the 12 Pokémon drawn for you, with six on the team and six on the bench.",
		TeamBuilding: "Use `/flip party bench:6` to draw six Pokémon for your team and six for your bench. You may swap Pokémon between them to choose which six you bring into battle.",
	},
}
