package booster

import (
	"context"

	"github.com/charmingruby/alomobot/internal/core/booster/deck"
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
	ID    int
	Types []pkmn.Type
}

type SupportAttr struct{}

type EnergyAttr struct{}

type Card struct {
	Name   string
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
	PokemonCards  []Card
	SupportCards  []Card
	EnergyCards   []Card
	Slots         []Slot
	RarityWeights map[Rarity]int
}

type Set interface {
	Open(ctx context.Context) ([]Card, error)
}

func NewPokemonCard(def deck.Pokemon, rarity Rarity) Card {
	return Card{
		Name:   def.Name,
		Type:   PokemonCardType,
		Rarity: rarity,
		PokemonAttr: &PokemonAttr{
			ID:    def.ID,
			Types: def.Types,
		},
	}
}
