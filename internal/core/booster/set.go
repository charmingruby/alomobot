package booster

import (
	"cmp"
	"context"
	"fmt"
	"math/rand/v2"
	"slices"
)

type Slot struct {
	Types     []CardType
	MinRarity Rarity
}

type SetName string

type Set struct {
	Name          SetName
	PokemonCards  []Card
	SupportCards  []Card
	EnergyCards   []Card
	Slots         []Slot
	RarityWeights map[Rarity]int
}

func (s Set) String() string {
	return string(s.Name)
}

func (s *Set) Open(ctx context.Context) ([]Card, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	pool := slices.Concat(s.PokemonCards, s.SupportCards, s.EnergyCards)
	cards := make([]Card, 0, len(s.Slots))

	for i, slot := range s.Slots {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		minRarity := max(0, slices.Index(rarities, slot.MinRarity))
		eligible := make(map[Rarity][]Card)
		totalWeight := 0

		// Filter the cards allowed in this slot.
		for _, card := range pool {
			if !slices.Contains(slot.Types, card.Type) ||
				slices.Index(rarities, card.Rarity) < minRarity ||
				s.RarityWeights[card.Rarity] <= 0 {
				continue
			}

			if len(eligible[card.Rarity]) == 0 {
				totalWeight += s.RarityWeights[card.Rarity]
			}

			eligible[card.Rarity] = append(eligible[card.Rarity], card)
		}

		if totalWeight == 0 {
			return nil, fmt.Errorf("slot %d: no eligible cards with positive rarity weight", i+1)
		}

		// Draw a rarity by weight, then a card within that rarity.
		roll := rand.IntN(totalWeight)
		for rarity, candidates := range eligible {
			roll -= s.RarityWeights[rarity]
			if roll >= 0 {
				continue
			}

			card := candidates[rand.IntN(len(candidates))]

			switch card.Type {
			case PokemonCardType:
				card.PokemonAttr = new(*card.PokemonAttr)

				if card.PokemonAttr.SecondaryType != nil {
					card.PokemonAttr.SecondaryType = new(*card.PokemonAttr.SecondaryType)
				}
			case SupportCardType:
				card.SupportAttr = new(*card.SupportAttr)
			case EnergyCardType:
				card.EnergyAttr = new(*card.EnergyAttr)
			}

			cards = append(cards, card)
			break
		}
	}

	slices.SortStableFunc(cards, func(a, b Card) int {
		return cmp.Compare(slices.Index(rarities, a.Rarity), slices.Index(rarities, b.Rarity))
	})

	return cards, nil
}
