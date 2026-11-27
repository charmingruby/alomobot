package cup

const BoosterCup CupName = "Booster"

var boosterCup = Cup{
	Name: BoosterCup,
	Guide: GuideCup{
		HowItWorks: "Open 2 booster packs containing Pokémon, Support, and Energy cards. " +
			"Build a team of 6 Pokémon and keep all remaining Pokémon on your bench. " +
			"Before using any Support or Energy cards, resolve all duplicate and Legendary Pokémon received from the packs. " +
			"These rules also apply to every Pokémon received later, including through Support and Energy cards.",
		Restrictions: "Legendary Pokémon cannot be used. Whenever you receive one, whether duplicated or not, " +
			"choose any non-Legendary Pokémon from the set that shares at least 1 of its types. " +
			"For dual-type Legendaries, either type is allowed. If the chosen Pokémon is a duplicate, apply the duplicate rule. " +
			"Using any Support or Energy card is always optional, including Energy created from duplicates.",
		TeamBuilding: "1. Open 2 packs and check the Pokémon received.\n" +
			"2. Replace each Legendary Pokémon with a non-Legendary Pokémon from the set sharing either of its types.\n" +
			"3. Convert each duplicate Pokémon into an Energy matching its primary type. " +
			"You may save this Energy or use it to draw a Pokémon of that type from the set.\n" +
			"4. After resolving the initial duplicates and Legendaries, optionally use your Support and Energy cards.\n" +
			"5. Apply the same rules to every new Pokémon received. If a draw produces another duplicate, " +
			"convert it into Energy matching its primary type again. You may repeat this process until you receive " +
			"a Pokémon you do not already have, or stop and save the Energy at any time.\n" +
			"6. Choose 6 Pokémon for your team. All remaining Pokémon form your bench.",
	},
}
