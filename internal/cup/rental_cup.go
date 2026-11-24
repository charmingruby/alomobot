package cup

const RentalCup NameCup = "Rental"

var rentalCup = Cup{
	Name: RentalCup,
	Guide: GuideCup{
		HowItWorks:   "Start with a normal team. After every battle, the winner takes the loser's team, and the loser takes the winner's team.",
		Restrictions: "Exchange the entire team after each battle. Continue the tournament with the team received from your opponent.",
		TeamBuilding: "Build your initial team following the general rules. After each battle, swap teams with your opponent and use the received team in your next battle.",
	},
}
