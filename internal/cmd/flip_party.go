package cmd

import (
	"context"
	"fmt"

	"github.com/charmingruby/erebor/discord/api"

	"github.com/charmingruby/alomobot/internal/cmd/message"
	"github.com/charmingruby/alomobot/internal/core"
	"github.com/charmingruby/alomobot/internal/format"
)

func flipParty() api.Command {
	const benchParam = "bench"

	return api.Command{
		Name:        "party",
		Description: "Flips a random Pokémon party",
		Parameters: []api.Parameter{
			{
				Type:        api.ParameterInteger,
				Name:        benchParam,
				Description: "Number of reserve Pokémon to include",
			},
		},
		Handler: func(ctx context.Context, i api.Interaction) (api.Response, error) {
			filter := core.FlipPartyFilter{
				Bench: int(i.Integer(benchParam)),
			}

			party, bench, err := core.FlipParty(ctx, filter)
			if err != nil {
				return api.Response{}, err
			}

			partyDs := format.List(party)
			content := fmt.Sprintf("**Party** (6/6)\n%s\n",
				partyDs,
			)

			partyPokesURL := make([]string, 0, 6)
			for _, p := range party {
				partyPokesURL = append(partyPokesURL, p.SpriteURL)
			}

			if len(bench) == 0 {
				return api.Response{
					Content:   message.Flip("party", content),
					ImageURLs: partyPokesURL,
				}, nil
			}

			benchDs := format.List(bench)
			content += fmt.Sprintf("\n**Bench** (%d)\n%s",
				len(bench),
				benchDs,
			)

			return api.Response{
				Content:   message.Flip("party", content),
				ImageURLs: partyPokesURL,
			}, nil
		},
	}
}
