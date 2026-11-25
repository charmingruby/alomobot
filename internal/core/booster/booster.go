package booster

import (
	"context"

	"github.com/charmingruby/alomobot/internal/core/pkmn"
)

type CardType string

const (
	EnergyCardType  CardType = "energy"
	PokemonCardType CardType = "pokemon"
	SupportCardType CardType = "support"
)

type Rarity string

const (
	CommonRarity    Rarity = "common"
	UncommonRarity  Rarity = "uncommon"
	RareRarity      Rarity = "rare"
	EpicRarity      Rarity = "epic"
	LegendaryRarity Rarity = "legendary"
)

type PokemonAttr struct {
	Types []pkmn.Type
}

type SupportAttr struct{}

type EnergyAttr struct{}

type Card struct {
	Name   string
	Set    string
	Type   CardType
	Rarity Rarity

	PokemonAttr *PokemonAttr
	SupportAttr *SupportAttr
	EnergyAttr  *EnergyAttr
}

type Slot struct {
	Types     []CardType
	MinRarity Rarity
}

type SetDefinition struct {
	Cards         []Card
	Slots         []Slot
	RarityWeights map[Rarity]int
}

type Set interface {
	Open(ctx context.Context, packs int) ([]Card, error)
}
