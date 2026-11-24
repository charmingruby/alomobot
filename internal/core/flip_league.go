package core

import (
	"context"

	"github.com/charmingruby/alomobot/internal/cup"
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

	if f.Cups > len(cup.Cups) || f.Cups <= 0 {
		f.Cups = standardLeagueSize
	}

	return f
}

func FlipLeague(ctx context.Context, filter FlipLeagueFilter) []cup.Cup {
	league := make([]cup.Cup, 0, filter.Cups)

	var unique map[cup.NameCup]struct{}
	if !filter.Repeatable {
		unique = make(map[cup.NameCup]struct{}, filter.Cups)
	}

	for len(league) < filter.Cups-1 {
		cup := FlipCup(ctx)

		if !filter.Repeatable {
			if _, ok := unique[cup.Name]; ok {
				continue
			}

			unique[cup.Name] = struct{}{}
		}

		league = append(league, cup)
	}

	standardCup, _ := cup.FindCup(string(cup.StandardCup))
	league = append(league, standardCup)

	return league
}
