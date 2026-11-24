package cup

const LetterCup NameCup = "Letter"

var letterCup = Cup{
	Name: LetterCup,
	Guide: GuideCup{
		HowItWorks:   "[Explain how Letter Cup works.]",
		Restrictions: "- [List allowed Pokémon and bans.]",
		TeamBuilding: "1. [Describe the steps to build your team.]",
	},
}
