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

type SupportMechanic string

const (
	CaptureMechanic        SupportMechanic = "Capture"
	RenewalRiskMechanic    SupportMechanic = "Renewal Risk"
	TypeSpecialistMechanic SupportMechanic = "Type Specialist"
	SynergyMechanic        SupportMechanic = "Synergy"
)

type SupportAttr struct {
	Mechanic  SupportMechanic
	Effect    string
	Condition string
}

type PokemonAttr struct {
	pkmn.Pokemon
}

type EnergyAttr struct {
	Type      pkmn.Type
	Effect    string
	Condition string
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
		description := c.SupportAttr.Effect
		if c.SupportAttr.Condition != "" {
			description = c.SupportAttr.Condition + " " + description
		}
		content += fmt.Sprintf(": _%s_", description)
	case EnergyCardType:
		description := c.EnergyAttr.Effect
		if c.EnergyAttr.Condition != "" {
			description = c.EnergyAttr.Condition + " " + description
		}
		content += fmt.Sprintf(": _%s_", description)
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

func NewSupportCard(name, effect, condition string, mechanic SupportMechanic) Card {
	return Card{
		Name: name,
		Type: SupportCardType,
		SupportAttr: &SupportAttr{
			Effect:    effect,
			Condition: condition,
			Mechanic:  mechanic,
		},
	}
}

func NewEnergyCard(name, effect string) Card {
	return Card{
		Name: name,
		Type: EnergyCardType,
		EnergyAttr: &EnergyAttr{
			Effect: effect,
		},
	}
}
