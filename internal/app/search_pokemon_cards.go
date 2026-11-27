package app

import (
	"context"
	"errors"
	"slices"
	"strings"

	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/set"
	"github.com/charmingruby/alomobot/internal/core/pkmn"
)

var (
	ErrPokemonCardNotFound     = errors.New("pokemon card not found in set")
	ErrPokemonCardPageNotFound = errors.New("pokemon card page not found")
)

type SearchPokemonCardsFilter struct {
	Set    string
	Name   string
	Rarity booster.Rarity
	Page   int
	Type   pkmn.Type
}

type SearchPokemonCardsResult struct {
	Set      booster.Set
	Cards    []booster.Card
	Page     int
	Pages    int
	Total    int
	Detailed bool
}

func (a *App) SearchPokemonCards(ctx context.Context, filter SearchPokemonCardsFilter) (SearchPokemonCardsResult, error) {
	const pageSize = 25

	if err := ctx.Err(); err != nil {
		return SearchPokemonCardsResult{}, err
	}

	selected, ok := set.FindSet(filter.Set)
	if !ok {
		return SearchPokemonCardsResult{}, ErrSetNotFound
	}

	filter.Name = strings.TrimSpace(filter.Name)
	cards := slices.DeleteFunc(slices.Clone(selected.PokemonCards), func(card booster.Card) bool {
		if (filter.Rarity != "" && card.Rarity != filter.Rarity) ||
			(filter.Name != "" && !strings.EqualFold(card.Name, filter.Name)) {
			return true
		}
		if filter.Type != "" {
			attr := card.PokemonAttr
			return attr.PrimaryType != filter.Type && (attr.SecondaryType == nil || *attr.SecondaryType != filter.Type)
		}
		return false
	})

	detailed := filter.Name != ""
	if detailed {
		if len(cards) == 0 {
			return SearchPokemonCardsResult{}, ErrPokemonCardNotFound
		}
		cards = cards[:1]
		filter.Page = 1

		pokemon, err := a.api.GetPokemonByDexID(ctx, cards[0].PokemonAttr.ID)
		if err != nil {
			return SearchPokemonCardsResult{}, err
		}
		cards[0].PokemonAttr = &booster.PokemonAttr{Pokemon: pokemon}
	}

	page := max(1, filter.Page)
	pages := max(1, (len(cards)+pageSize-1)/pageSize)
	if page > pages {
		return SearchPokemonCardsResult{}, ErrPokemonCardPageNotFound
	}
	start := (page - 1) * pageSize
	end := min(start+pageSize, len(cards))

	return SearchPokemonCardsResult{
		Set: selected, Cards: cards[start:end],
		Page: page, Pages: pages, Total: len(cards), Detailed: detailed,
	}, nil
}
