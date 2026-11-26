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
	ErrSupportCardNotFound     = errors.New("support card not found in set")
	ErrSupportCardPageNotFound = errors.New("support card page not found")
)

type SearchSupportCardsFilter struct {
	Set    string
	Name   string
	Rarity booster.Rarity
	Page   int
}

type SearchSupportCardsResult struct {
	Set      booster.Set
	Cards    []booster.Card
	Page     int
	Pages    int
	Total    int
	Detailed bool
}

func (a *App) SearchSupportCards(ctx context.Context, filter SearchSupportCardsFilter) (SearchSupportCardsResult, error) {
	const pageSize = 25

	if err := ctx.Err(); err != nil {
		return SearchSupportCardsResult{}, err
	}

	selected, ok := set.FindSet(filter.Set)
	if !ok {
		return SearchSupportCardsResult{}, ErrSetNotFound
	}

	filter.Name = strings.TrimSpace(filter.Name)
	cards := slices.DeleteFunc(slices.Clone(selected.SupportCards), func(card booster.Card) bool {
		return (filter.Rarity != "" && card.Rarity != filter.Rarity) ||
			(filter.Name != "" && !strings.EqualFold(card.Name, filter.Name))
	})

	detailed := filter.Name != ""
	if detailed {
		if len(cards) == 0 {
			return SearchSupportCardsResult{}, ErrSupportCardNotFound
		}
		cards = cards[:1]
		filter.Page = 1
	}

	page := max(1, filter.Page)
	pages := max(1, (len(cards)+pageSize-1)/pageSize)
	if page > pages {
		return SearchSupportCardsResult{}, ErrSupportCardPageNotFound
	}
	start := (page - 1) * pageSize
	end := min(start+pageSize, len(cards))

	return SearchSupportCardsResult{
		Set: selected, Cards: cards[start:end],
		Page: page, Pages: pages, Total: len(cards), Detailed: detailed,
	}, nil
}
