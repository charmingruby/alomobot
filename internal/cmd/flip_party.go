package cmd

import (
	"context"
	"fmt"

	"github.com/charmingruby/erebor/discord/api"

	"github.com/charmingruby/alomobot/internal/app"
	"github.com/charmingruby/alomobot/internal/cmd/template"
	"github.com/charmingruby/alomobot/internal/format"
)

func (c *cmd) flipParty() api.Command {
	const (
		benchParam   = "bench"
		partyPattern = `**Party** (6/6)
%s
`
		benchPattern = `
**Bench** (%d)
%s`
	)

	return api.Command{
		Name:        "flip",
		Description: "Flips a random Pokémon party",
		Parameters: []api.Parameter{
			{
				Type:        api.ParameterInteger,
				Name:        benchParam,
				Description: "Number of reserve Pokémon to include",
			},
		},
		Handler: func(ctx context.Context, i api.Interaction) (api.Response, error) {
			filter := app.FlipPartyFilter{
				Bench: int(i.Integer(benchParam)),
			}

			party, bench, err := c.app.FlipParty(ctx, filter)
			if err != nil {
				return api.Response{}, err
			}

			partyDs := format.List(party)
			content := fmt.Sprintf(
				partyPattern,
				partyDs,
			)

			partyPokesURL := make([]string, 0, 6)
			for _, p := range party {
				partyPokesURL = append(partyPokesURL, p.SpriteURL)
			}

			if len(bench) == 0 {
				return api.Response{
					Content:   template.Flip("party", content),
					ImageURLs: partyPokesURL,
				}, nil
			}

			benchDs := format.List(bench)
			content += fmt.Sprintf(
				benchPattern,
				len(bench),
				benchDs,
			)

			return api.Response{
				Content:   template.Flip("party", content),
				ImageURLs: partyPokesURL,
			}, nil
		},
	}
}
