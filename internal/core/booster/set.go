package booster

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"slices"

	"github.com/charmingruby/alomobot/internal/core/pkmn"
)

const MaxDrawQuantity = 10

var (
	ErrInvalidDrawFilter = errors.New("invalid draw filter")
	ErrInsufficientCards = errors.New("not enough eligible cards")
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
	RarityWeights map[Rarity]float64
}

type DrawFilter struct {
	Quantity  int
	MinRarity Rarity
	MaxRarity Rarity
	Types     []pkmn.Type
}

func (f DrawFilter) Normalize() DrawFilter {
	if f.Quantity == 0 {
		f.Quantity = 1
	}

	return f
}

func (f DrawFilter) Validate() error {
	if f.Quantity < 1 {
		return ErrInvalidDrawFilter
	}

	if f.MinRarity != "" && f.MaxRarity != "" &&
		slices.Index(Rarities, f.MinRarity) > slices.Index(Rarities, f.MaxRarity) {
		return ErrInvalidDrawFilter
	}

	for _, kind := range f.Types {
		if !slices.Contains(pkmn.Types, kind) {
			return ErrInvalidDrawFilter
		}
	}

	return nil
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

		minRarity := max(0, slices.Index(Rarities, slot.MinRarity))
		eligible := make(map[Rarity][]Card)
		totalWeight := 0.0

		// Filter the cards allowed in this slot.
		for _, card := range pool {
			if !slices.Contains(slot.Types, card.Type) ||
				slices.Index(Rarities, card.Rarity) < minRarity ||
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

		card, err := s.drawWeighted(eligible, totalWeight)
		if err != nil {
			return nil, fmt.Errorf("slot %d: %w", i+1, err)
		}
		cards = append(cards, card)
	}

	slices.SortStableFunc(cards, func(a, b Card) int {
		return cmp.Compare(slices.Index(Rarities, a.Rarity), slices.Index(Rarities, b.Rarity))
	})

	return cards, nil
}

// Draw reveals distinct Pokémon cards using the set's rarity weights.
// Types match either Pokémon type, with OR semantics between requested types.
func (s Set) Draw(ctx context.Context, filter DrawFilter) ([]Card, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	filter = filter.Normalize()
	if err := filter.Validate(); err != nil {
		return nil, err
	}

	minRarity, maxRarity := max(0, slices.Index(Rarities, filter.MinRarity)), len(Rarities)-1
	if filter.MaxRarity != "" {
		maxRarity = slices.Index(Rarities, filter.MaxRarity)
	}

	eligible := make(map[Rarity][]Card)
	seen := make(map[string]bool)
	totalWeight, count := 0.0, 0

	for _, card := range s.PokemonCards {
		rank := slices.Index(Rarities, card.Rarity)

		if card.Type != PokemonCardType || card.PokemonAttr == nil ||
			rank < minRarity || rank > maxRarity || s.RarityWeights[card.Rarity] <= 0 ||
			seen[card.Name] {
			continue
		}

		attr := card.PokemonAttr
		if len(filter.Types) > 0 && !slices.Contains(filter.Types, attr.PrimaryType) &&
			(attr.SecondaryType == nil || !slices.Contains(filter.Types, *attr.SecondaryType)) {
			continue
		}

		if len(eligible[card.Rarity]) == 0 {
			totalWeight += s.RarityWeights[card.Rarity]
		}

		eligible[card.Rarity] = append(eligible[card.Rarity], card)
		seen[card.Name] = true
		count++
	}

	if count < filter.Quantity {
		return nil, ErrInsufficientCards
	}

	cards := make([]Card, 0, filter.Quantity)
	for range filter.Quantity {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		card, err := s.drawWeighted(eligible, totalWeight)
		if err != nil {
			return nil, err
		}

		cards = append(cards, card)
		eligible[card.Rarity] = slices.DeleteFunc(eligible[card.Rarity], func(candidate Card) bool {
			return candidate.Name == card.Name
		})

		if len(eligible[card.Rarity]) == 0 {
			totalWeight -= s.RarityWeights[card.Rarity]
		}
	}

	slices.SortStableFunc(cards, func(a, b Card) int {
		return cmp.Compare(slices.Index(Rarities, a.Rarity), slices.Index(Rarities, b.Rarity))
	})

	return cards, nil
}

// drawWeighted draws a rarity by weight, then a card uniformly within that rarity.
// The returned card's attributes are independent of the set's catalog.
func (s Set) drawWeighted(eligible map[Rarity][]Card, totalWeight float64) (Card, error) {
	if totalWeight <= 0 {
		return Card{}, errors.New("total rarity weight must be positive")
	}

	roll := rand.Float64() * totalWeight

	for _, rarity := range Rarities {
		candidates := eligible[rarity]
		if len(candidates) == 0 {
			continue
		}

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

		return card, nil
	}

	return Card{}, errors.New("no eligible card found for weighted draw")
}
