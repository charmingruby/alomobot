package core

import (
	"context"
	"fmt"

	"github.com/charmingruby/alomobot/internal/pkmn"
)

type FlipLeagueFilter struct {
	Cups       int
	Repeatable bool
}

func (f FlipLeagueFilter) Normalize() FlipLeagueFilter {
	cups := 7

	if f.Cups > len(pkmn.Cups) || f.Cups <= 0 {
		f.Cups = cups
	}

	return f
}

func FlipLeague(ctx context.Context, filter FlipLeagueFilter) []pkmn.Cup {
	league := make([]pkmn.Cup, 0, filter.Cups+1)

	fmt.Printf("%+v\n", filter)

	var unique map[pkmn.Cup]struct{}
	if !filter.Repeatable {
		unique = make(map[pkmn.Cup]struct{}, filter.Cups)
	}

	for len(league) < filter.Cups {
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
