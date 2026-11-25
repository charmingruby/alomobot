package cup

const StandardCup CupName = "Standard"

var standardCup = Cup{
	Name: StandardCup,
	Guide: GuideCup{
		HowItWorks:   "Battles follow the general tournament rules, with no additional cup-specific restrictions.",
		Restrictions: "All restrictions and bans from the general rules apply.",
		TeamBuilding: "Build your team according to the general tournament rules.",
	},
}
