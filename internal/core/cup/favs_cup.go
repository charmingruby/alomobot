package cup

const FavsCup CupName = "Favs"

var favsCup = Cup{
	Name: FavsCup,
	Guide: GuideCup{
		HowItWorks:   "Each player contributes their favorite Pokémon to the shared Favs pool. Use /pool flip to randomly deal a team of six to each participating player.",
		Restrictions: "Draw without replacement from the shared pool. Shared favorites count as separate available copies. Unused entries remain outside the teams.",
		TeamBuilding: "Use /pool flip with pool:Favs, groups set to the participating players, and size:6. Build your team using the Pokémon dealt to your group, following the general tournament rules.",
	},
}
