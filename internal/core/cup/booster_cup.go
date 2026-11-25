package cup

const BoosterCup NameCup = "Booster"

var boosterCup = Cup{
	Name: BoosterCup,
	Guide: GuideCup{
		HowItWorks:   "[Explain how Booster Cup works.]",
		Restrictions: "- [List allowed Pokémon and bans.]",
		TeamBuilding: "1. [Describe the steps to build your team.]",
	},
}
