package cup

const StandardCup NameCup = "Standard"

var standardCup = Cup{
	Name: StandardCup,
	Guide: GuideCup{
		HowItWorks:   "[Explain how Standard Cup works.]",
		Restrictions: "- [List allowed Pokémon and bans.]",
		TeamBuilding: "1. [Describe the steps to build your team.]",
	},
}
