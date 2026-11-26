package deck

import (
	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/pkmn"
)

func energy(name string, energyType pkmn.Type, description string) booster.Card {
	card := booster.NewEnergyCard(name, description)
	card.EnergyAttr.Type = energyType
	return card
}

var (
	NormalEnergy   = energy("Normal Energy", pkmn.TypeNormal, "Provides Normal-type energy.")
	FireEnergy     = energy("Fire Energy", pkmn.TypeFire, "Provides Fire-type energy.")
	WaterEnergy    = energy("Water Energy", pkmn.TypeWater, "Provides Water-type energy.")
	ElectricEnergy = energy("Electric Energy", pkmn.TypeElectric, "Provides Electric-type energy.")
	GrassEnergy    = energy("Grass Energy", pkmn.TypeGrass, "Provides Grass-type energy.")
	IceEnergy      = energy("Ice Energy", pkmn.TypeIce, "Provides Ice-type energy.")
	FightingEnergy = energy("Fighting Energy", pkmn.TypeFighting, "Provides Fighting-type energy.")
	PoisonEnergy   = energy("Poison Energy", pkmn.TypePoison, "Provides Poison-type energy.")
	GroundEnergy   = energy("Ground Energy", pkmn.TypeGround, "Provides Ground-type energy.")
	FlyingEnergy   = energy("Flying Energy", pkmn.TypeFlying, "Provides Flying-type energy.")
	PsychicEnergy  = energy("Psychic Energy", pkmn.TypePsychic, "Provides Psychic-type energy.")
	BugEnergy      = energy("Bug Energy", pkmn.TypeBug, "Provides Bug-type energy.")
	RockEnergy     = energy("Rock Energy", pkmn.TypeRock, "Provides Rock-type energy.")
	GhostEnergy    = energy("Ghost Energy", pkmn.TypeGhost, "Provides Ghost-type energy.")
	DragonEnergy   = energy("Dragon Energy", pkmn.TypeDragon, "Provides Dragon-type energy.")
	DarkEnergy     = energy("Dark Energy", pkmn.TypeDark, "Provides Dark-type energy.")
	SteelEnergy    = energy("Steel Energy", pkmn.TypeSteel, "Provides Steel-type energy.")
	FairyEnergy    = energy("Fairy Energy", pkmn.TypeFairy, "Provides Fairy-type energy.")
	NeutralEnergy  = energy(
		"Neutral Energy",
		pkmn.TypeNormal,
		"Choose any Pokémon type when using this card. This energy counts as the chosen type.",
	)
)
