package app

import (
	"context"
	"errors"
	"slices"
	"strings"

	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/set"
)

var (
	ErrEnergyCardNotFound     = errors.New("energy card not found in set")
	ErrEnergyCardPageNotFound = errors.New("energy card page not found")
)

type SearchEnergyCardsFilter struct {
	Set    string
	Name   string
	Rarity booster.Rarity
	Page   int
}

type SearchEnergyCardsResult struct {
	Set      booster.Set
	Cards    []booster.Card
	Page     int
	Pages    int
	Total    int
	Detailed bool
}

func (a *App) SearchEnergyCards(ctx context.Context, filter SearchEnergyCardsFilter) (SearchEnergyCardsResult, error) {
	const pageSize = 25

	if err := ctx.Err(); err != nil {
		return SearchEnergyCardsResult{}, err
	}

	selected, ok := set.FindSet(filter.Set)
	if !ok {
		return SearchEnergyCardsResult{}, ErrSetNotFound
	}

	filter.Name = strings.TrimSpace(filter.Name)
	cards := slices.DeleteFunc(slices.Clone(selected.EnergyCards), func(card booster.Card) bool {
		return (filter.Rarity != "" && card.Rarity != filter.Rarity) ||
			(filter.Name != "" && !strings.EqualFold(card.Name, filter.Name))
	})

	detailed := filter.Name != ""
	if detailed {
		if len(cards) == 0 {
			return SearchEnergyCardsResult{}, ErrEnergyCardNotFound
		}
		cards = cards[:1]
		filter.Page = 1
	}

	page := max(1, filter.Page)
	pages := max(1, (len(cards)+pageSize-1)/pageSize)
	if page > pages {
		return SearchEnergyCardsResult{}, ErrEnergyCardPageNotFound
	}
	start := (page - 1) * pageSize
	end := min(start+pageSize, len(cards))

	return SearchEnergyCardsResult{
		Set: selected, Cards: cards[start:end],
		Page: page, Pages: pages, Total: len(cards), Detailed: detailed,
	}, nil
}
