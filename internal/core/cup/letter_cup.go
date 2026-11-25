package cup

const LetterCup CupName = "Letter"

var letterCup = Cup{
	Name: LetterCup,
	Guide: GuideCup{
		HowItWorks:   "A letter is randomly drawn, and every Pokémon on your team must have a species name starting with it.",
		Restrictions: "Use species names, not nicknames, to determine eligibility. Pokémon whose names start with another letter are not allowed.",
		TeamBuilding: "Wait for the letter draw, then build your team using matching Pokémon. For example, if B is drawn, Bulbasaur, Butterfree, and Blastoise are allowed.",
	},
}
