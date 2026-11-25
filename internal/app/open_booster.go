package app

import (
	"context"
	"errors"

	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/set"
)

var ErrSetNotFound = errors.New("set not found")

type OpenBoosterResult struct {
	Set   booster.Set
	Cards []booster.Card
}

func (a *App) OpenBooster(ctx context.Context, name string) (OpenBoosterResult, error) {
	selected, ok := set.FindSet(name)
	if !ok {
		return OpenBoosterResult{}, ErrSetNotFound
	}

	cards, err := selected.Open(ctx)
	if err != nil {
		return OpenBoosterResult{}, err
	}

	return OpenBoosterResult{Set: selected, Cards: cards}, nil
}
