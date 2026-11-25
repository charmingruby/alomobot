package app

import (
	"context"
	"errors"

	"github.com/charmingruby/alomobot/internal/core/cup"
)

var ErrCupNotFound = errors.New("cup not found")

func (a *App) RulesCup(_ context.Context, name string) (cup.Cup, error) {
	selected, ok := cup.FindCup(name)
	if !ok {
		return cup.Cup{}, ErrCupNotFound
	}

	return selected, nil
}
