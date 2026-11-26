package booster

import (
	"fmt"

	"github.com/charmingruby/alomobot/internal/core/pkmn"
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

type SupportAttr struct {
	Name        string
	Description string
}

type EnergyAttr struct {
	Name        string
	Type        pkmn.Type
	Description string
}

type Card struct {
	Name        string
	Type        CardType
	Rarity      Rarity
	PokemonAttr *PokemonAttr
	SupportAttr *SupportAttr
	EnergyAttr  *EnergyAttr
}

func (c Card) WithRarity(rarity Rarity) Card {
	c.Rarity = rarity

	return c
}

func (c Card) Render() string {
	content := fmt.Sprintf("%s **%s**", c.Rarity.Emoji(), c.Name)

	switch c.Type {
	case PokemonCardType:
		content += fmt.Sprintf(" • %s", c.PokemonAttr.PrimaryType.Render())

		if c.PokemonAttr.SecondaryType != nil {
			content += fmt.Sprintf(" / %s", c.PokemonAttr.SecondaryType.Render())
		}
	case SupportCardType:
		content += fmt.Sprintf(": _%s_", c.SupportAttr.Description)
	case EnergyCardType:
		content += fmt.Sprintf(": _%s_", c.EnergyAttr.Description)
	}

	return content
}

func NewPokemonCard(id int, name string, types ...pkmn.Type) Card {
	attr := &PokemonAttr{
		Pokemon: pkmn.Pokemon{
			ID:          id,
			Name:        name,
			PrimaryType: types[0],
		},
	}
	if len(types) > 1 {
		attr.SecondaryType = new(types[1])
	}

	return Card{
		Name:        name,
		Type:        PokemonCardType,
		PokemonAttr: attr,
	}
}

func NewSupportCard(name, description string) Card {
	return Card{
		Name: name,
		Type: SupportCardType,
		SupportAttr: &SupportAttr{
			Name:        name,
			Description: description,
		},
	}
}

func NewEnergyCard(name, description string) Card {
	return Card{
		Name: name,
		Type: EnergyCardType,
		EnergyAttr: &EnergyAttr{
			Name:        name,
			Description: description,
		},
	}
}
