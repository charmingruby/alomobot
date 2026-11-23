package cmd

import (
	"context"
	"fmt"

	"github.com/charmingruby/erebor/discord/command"

	"github.com/charmingruby/alomobot/internal/cmd/message"
	"github.com/charmingruby/alomobot/internal/core"
	"github.com/charmingruby/alomobot/internal/format"
)

func flipParty() command.Definition {
	const benchParam = "bench"

	return command.Definition{
		Name:        "flip-party",
		Description: "Flips a random Pokémon party",
		Parameters: []*command.Parameter{
			command.Integer(command.Parameter{
				Name:        benchParam,
				Description: "Number of reserve Pokémon to include",
			}),
		},
		Handler: func(ctx context.Context, i command.Interaction) (string, error) {
			filter := core.FlipPartyFilter{}

			for _, opt := range i.Options {
				switch opt.Name {
				case benchParam:
					filter.Bench = int(opt.IntValue())
				}
			}

			party, bench, err := core.FlipParty(ctx, filter)
			if err != nil {
				return "", err
			}

			partyDs := format.List(party)
			content := fmt.Sprintf("**Party** (6/6)\n%s\n",
				partyDs,
			)

			if len(bench) == 0 {
				return message.Flip("party", content), nil
			}

			benchDs := format.List(bench)
			content += fmt.Sprintf("\n**Bench** (%d)\n%s",
				len(bench),
				benchDs,
			)

			return message.Flip("party", content), nil
		},
	}
}
