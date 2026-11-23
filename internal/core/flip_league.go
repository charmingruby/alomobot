package core

import (
	"context"

	"github.com/charmingruby/alomobot/internal/pkmn"
)

const standardLeagueSize = 8

type FlipLeagueFilter struct {
	Cups       int
	Repeatable bool
}

func (f FlipLeagueFilter) Normalize() FlipLeagueFilter {
	if f.Repeatable && f.Cups >= 0 {
		return f
	}

	if f.Cups > len(pkmn.Cups) || f.Cups <= 0 {
		f.Cups = standardLeagueSize
	}

	return f
}

func FlipLeague(ctx context.Context, filter FlipLeagueFilter) []pkmn.Cup {
	league := make([]pkmn.Cup, 0, filter.Cups)

	var unique map[pkmn.Cup]struct{}
	if !filter.Repeatable {
		unique = make(map[pkmn.Cup]struct{}, filter.Cups)
	}

	for len(league) < filter.Cups-1 {
		cup := FlipCup(ctx)

		if !filter.Repeatable {
			if _, ok := unique[cup]; ok {
				continue
			}

			unique[cup] = struct{}{}
		}

		league = append(league, cup)
	}

	league = append(league, pkmn.CupStandard)

	return league
}
