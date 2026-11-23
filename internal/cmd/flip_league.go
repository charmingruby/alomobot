package cmd

import (
	"context"
	"fmt"

	"github.com/charmingruby/erebor/discord/command"

	"github.com/charmingruby/alomobot/internal/cmd/message"
	"github.com/charmingruby/alomobot/internal/core"
	"github.com/charmingruby/alomobot/internal/format"
)

func FlipLeagueDef() command.Definition {
	cupsParam := "cups"
	repeatableParam := "repeatable"

	return command.Definition{
		Name:        "flip-league",
		Description: "Flips a random league of cups",
		Parameters: []*command.Parameter{
			command.Integer(command.Parameter{
				Name:        cupsParam,
				Description: "Number of cups before the finals",
			}),
			command.Boolean(command.Parameter{
				Name:        repeatableParam,
				Description: "Allow repeated cups",
			}),
		},
		Handler: func(ctx context.Context, i command.Interaction) (string, error) {
			filter := core.FlipLeagueFilter{}

			for _, opt := range i.Options {
				switch opt.Name {
				case cupsParam:
					filter.Cups = int(opt.IntValue())
				case repeatableParam:
					filter.Repeatable = opt.BoolValue()
				}
			}

			filter = filter.Normalize()

			league := core.FlipLeague(ctx, filter)

			content := fmt.Sprintf("**League**(%d)\n%s", len(league), format.List(league))

			return message.Flip("league", content), nil
		},
	}
}
