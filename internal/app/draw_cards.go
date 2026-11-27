package app

import (
	"context"
	"strconv"

	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/set"
)

type DrawCardsFilter struct {
	Set string
	booster.DrawFilter
}

type DrawCardsResult struct {
	Set   booster.Set
	Cards []booster.Card
}

func (a *App) DrawCards(ctx context.Context, filter DrawCardsFilter) (DrawCardsResult, error) {
	selected, ok := set.FindSet(filter.Set)
	if !ok {
		return DrawCardsResult{}, ErrSetNotFound
	}

	cards, err := selected.Draw(ctx, filter.DrawFilter)
	if err != nil {
		return DrawCardsResult{}, err
	}

	for _, card := range cards {
		pokemon, err := a.api.GetPokemon(ctx, strconv.Itoa(card.PokemonAttr.ID))
		if err != nil {
			return DrawCardsResult{}, err
		}

		card.PokemonAttr.Pokemon = pokemon
	}

	return DrawCardsResult{
		Set:   selected,
		Cards: cards,
	}, nil
}
