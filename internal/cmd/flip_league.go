package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmingruby/erebor/discord/command"

	"github.com/charmingruby/alomobot/internal/core"
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

			var message strings.Builder
			for idx, r := range league {
				if idx > 0 {
					message.WriteString("\n")
				}

				fmt.Fprintf(&message, "%d. %s", idx+1, r.String())
			}

			return message.String(), nil
		},
	}
}
