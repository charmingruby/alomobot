package booster

import (
	"github.com/charmingruby/alomobot/internal/core/booster/deck"
	"github.com/charmingruby/alomobot/internal/core/pkmn"
)

type CardType string

const (
	EnergyCardType  CardType = "energy"
	PokemonCardType CardType = "pokemon"
	SupportCardType CardType = "support"
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
