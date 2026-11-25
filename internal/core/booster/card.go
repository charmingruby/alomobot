package booster

import (
	"fmt"

	"github.com/charmingruby/alomobot/internal/core/booster/deck"
	"github.com/charmingruby/alomobot/internal/core/pkmn"
	"github.com/charmingruby/alomobot/internal/format"
)

type CardType string

const (
	EnergyCardType  CardType = "energy"
	PokemonCardType CardType = "pokémon"
	SupportCardType CardType = "support"
)

type PokemonAttr struct {
	pkmn.Pokemon
}

type SupportAttr struct{}

type EnergyAttr struct{}

type Card struct {
	Name        string
	Type        CardType
	Rarity      Rarity
	PokemonAttr *PokemonAttr
	SupportAttr *SupportAttr
	EnergyAttr  *EnergyAttr
}

func NewPokemonCard(def deck.Pokemon, rarity Rarity) Card {
	attr := &PokemonAttr{
		Pokemon: pkmn.Pokemon{
			ID:          def.ID,
			Name:        def.Name,
			PrimaryType: def.Types[0],
		},
	}
	if len(def.Types) > 1 {
		attr.SecondaryType = new(def.Types[1])
	}

	return Card{
		Name:        def.Name,
		Type:        PokemonCardType,
		Rarity:      rarity,
		PokemonAttr: attr,
	}
}

func (c Card) Render() string {
	label := format.Capitalize(string(c.Type))

	content := fmt.Sprintf("%s **%s** (`%s`)", c.Rarity.Emoji(), c.Name, label)

	if c.Type == PokemonCardType {
		content += fmt.Sprintf(" • %s", c.PokemonAttr.PrimaryType.Render())

		if c.PokemonAttr.SecondaryType != nil {
			content += fmt.Sprintf(" / %s", c.PokemonAttr.SecondaryType.Render())
		}
	}

	return content
}
