package deck

import "github.com/charmingruby/alomobot/internal/core/booster"

var (
	PokeBall = booster.NewSupportCard(
		"Poké Ball",
		"Discard 1 Pokémon of your choice from your round collection. Receive 1 random Pokémon.",
		"",
		booster.CaptureMechanic,
	)
	GreatBall = booster.NewSupportCard(
		"Great Ball",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 Pokémon and receive 1 of your choice.",
		"",
		booster.CaptureMechanic,
	)
	UltraBall = booster.NewSupportCard(
		"Ultra Ball",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 Pokémon of at least Uncommon rarity and receive 1 of your choice.",
		"",
		booster.CaptureMechanic,
	)
	NetBall = booster.NewSupportCard(
		"Net Ball",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 Water- or Bug-type Pokémon and receive 1 of your choice.",
		"",
		booster.CaptureMechanic,
	)
	DuskBall = booster.NewSupportCard(
		"Dusk Ball",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 Dark- or Ghost-type Pokémon and receive 1 of your choice.",
		"",
		booster.CaptureMechanic,
	)
	RepeatBall = booster.NewSupportCard(
		"Repeat Ball",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 2 Pokémon that share at least 1 type with the discarded Pokémon and receive 1 of your choice.",
		"",
		booster.CaptureMechanic,
	)
	NestBall = booster.NewSupportCard(
		"Nest Ball",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 2 Pokémon and receive 1 of your choice.",
		"",
		booster.CaptureMechanic,
	)
	DiveBall = booster.NewSupportCard(
		"Dive Ball",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 Water-type Pokémon of at least Uncommon rarity and receive 1 of your choice.",
		"",
		booster.CaptureMechanic,
	)
	LureBall = booster.NewSupportCard(
		"Lure Ball",
		"Discard 1 Pokémon of your choice from your round collection. Receive 1 random Water-type Pokémon of at least Rare rarity.",
		"",
		booster.CaptureMechanic,
	)
	MoonBall = booster.NewSupportCard(
		"Moon Ball",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 2 Ghost-, Dark- or Psychic-type Pokémon and receive 1 of your choice.",
		"",
		booster.CaptureMechanic,
	)
	LevelBall = booster.NewSupportCard(
		"Level Ball",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 Pokémon of equal or higher rarity than the discarded Pokémon and receive 1 of your choice.",
		"",
		booster.CaptureMechanic,
	)
	QuickBall = booster.NewSupportCard(
		"Quick Ball",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 4 Pokémon and receive 1 of your choice.",
		"Must be your first Support used this round.",
		booster.CaptureMechanic,
	)
	TimerBall = booster.NewSupportCard(
		"Timer Ball",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 4 Pokémon of at least Uncommon rarity and receive 1 of your choice.",
		"Must be your second Support used this round.",
		booster.CaptureMechanic,
	)
	SportBall = booster.NewSupportCard(
		"Sport Ball",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 Bug-type Pokémon of at least Uncommon rarity and receive 1 of your choice.",
		"",
		booster.CaptureMechanic,
	)
	DreamBall = booster.NewSupportCard(
		"Dream Ball",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 2 Psychic- or Fairy-type Pokémon of at least Uncommon rarity and receive 1 of your choice.",
		"",
		booster.CaptureMechanic,
	)
	SafariBall = booster.NewSupportCard(
		"Safari Ball",
		"Discard 1 Pokémon of your choice from your round collection. Choose a generation. Reveal 3 Pokémon introduced in that generation and receive 1 of your choice.",
		"",
		booster.CaptureMechanic,
	)
	LuxuryBall = booster.NewSupportCard(
		"Luxury Ball",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 2 Pokémon of at least Rare rarity and receive 1 of your choice.",
		"",
		booster.CaptureMechanic,
	)
	MasterBall = booster.NewSupportCard(
		"Master Ball",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 Pokémon of at least Epic rarity and receive 1 of your choice. You cannot choose a species outside the revealed options.",
		"",
		booster.CaptureMechanic,
	)
)

var (
	ProfessorOak = booster.NewSupportCard(
		"Professor Oak",
		"Discard 2 Pokémon of your choice from your round collection. For each discarded Pokémon, receive 1 random Pokémon of the same rarity.",
		"You must have at least 2 Pokémon available to discard.",
		booster.RenewalRiskMechanic,
	)
	ProfessorBirch = booster.NewSupportCard(
		"Professor Birch",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 2 Pokémon of the same rarity and receive 1 of your choice.",
		"",
		booster.RenewalRiskMechanic,
	)
	Gambler = booster.NewSupportCard(
		"Gambler",
		"Discard 1 Common or Uncommon Pokémon of your choice from your round collection. With equal probability, receive either 1 random Common Pokémon or 1 random Rare Pokémon.",
		"You must have a Common or Uncommon Pokémon available to discard.",
		booster.RenewalRiskMechanic,
	)
	ProfessorSycamore = booster.NewSupportCard(
		"Professor Sycamore",
		"Discard 3 Pokémon of your choice from your round collection. For each discarded Pokémon, receive 1 random Pokémon of the same rarity.",
		"You must have at least 3 Pokémon available to discard.",
		booster.RenewalRiskMechanic,
	)
)

var (
	Ilima = booster.NewSupportCard(
		"Ilima",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 normal-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	Blaine = booster.NewSupportCard(
		"Blaine",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 fire-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	Fisherman = booster.NewSupportCard(
		"Fisherman",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 water-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	LtSurge = booster.NewSupportCard(
		"Lt. Surge",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 electric-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	Erika = booster.NewSupportCard(
		"Erika",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 grass-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	Pryce = booster.NewSupportCard(
		"Pryce",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 ice-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	Bruno = booster.NewSupportCard(
		"Bruno",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 fighting-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	Koga = booster.NewSupportCard(
		"Koga",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 poison-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	Giovanni = booster.NewSupportCard(
		"Giovanni",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 ground-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	Falkner = booster.NewSupportCard(
		"Falkner",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 flying-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	Sabrina = booster.NewSupportCard(
		"Sabrina",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 psychic-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	BugCatcher = booster.NewSupportCard(
		"Bug Catcher",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 bug-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	Brock = booster.NewSupportCard(
		"Brock",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 rock-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	Morty = booster.NewSupportCard(
		"Morty",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 ghost-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	Clair = booster.NewSupportCard(
		"Clair",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 dragon-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	Karen = booster.NewSupportCard(
		"Karen",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 dark-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	Jasmine = booster.NewSupportCard(
		"Jasmine",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 steel-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)
	Valerie = booster.NewSupportCard(
		"Valerie",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 fairy-type Pokémon of up to Rare rarity and receive 1 of your choice.",
		"",
		booster.TypeSpecialistMechanic,
	)

	PokemonBreeder = booster.NewSupportCard(
		"Pokémon Breeder",
		"Discard 1 Pokémon of your choice from your round collection. Choose and receive any eligible Pokémon of up to Epic rarity that shares at least 1 Egg Group with it. Undiscovered does not count as a shared Egg Group, and Ditto is not a wildcard.",
		"",
		booster.SynergyMechanic,
	)
	ProfessorElm = booster.NewSupportCard(
		"Professor Elm",
		"Choose a generation represented by at least 2 Pokémon in your team. Discard 1 Pokémon of your choice from your round collection. Reveal 3 Pokémon introduced in that generation, of at least Uncommon rarity, and receive 1 of your choice.",
		"Your team must contain at least 2 Pokémon introduced in the same generation.",
		booster.SynergyMechanic,
	)
	Bill = booster.NewSupportCard(
		"Bill",
		"Discard 1 Pokémon of your choice from your round collection. Reveal 3 Pokémon of up to Uncommon rarity that have no types present in your team before this effect, and receive 1 of your choice.",
		"You must have at least 3 Pokémon in your team.",
		booster.SynergyMechanic,
	)
	PokemonRanger = booster.NewSupportCard(
		"Pokémon Ranger",
		"Choose a type shared by at least 2 Pokémon in your team. Discard 1 Pokémon of your choice from your round collection. Reveal 3 Pokémon of that type, of up to Rare rarity, and receive 1 of your choice.",
		"Your team must contain at least 2 Pokémon sharing a type.",
		booster.SynergyMechanic,
	)
)
