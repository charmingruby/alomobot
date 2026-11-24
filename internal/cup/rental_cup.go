package cup

const RentalCup NameCup = "Rental"

var rentalCup = Cup{
	Name: RentalCup,
	Guide: GuideCup{
		HowItWorks:   "[Explain how Rental Cup works.]",
		Restrictions: "- [List allowed Pokémon and bans.]",
		TeamBuilding: "1. [Describe the steps to build your team.]",
	},
}
