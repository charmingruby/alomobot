package cmd

import (
	"context"
	"fmt"

	"github.com/charmingruby/erebor/discord/api"

	"github.com/charmingruby/alomobot/internal/cmd/template"
	"github.com/charmingruby/alomobot/internal/core"
	"github.com/charmingruby/alomobot/internal/format"
)

func flipLeague() api.Command {
	const (
		cupsParam       = "cups"
		repeatableParam = "repeatable"
		leaguePattern   = `**League**(%d)
%s`
	)

	return api.Command{
		Name:        "league",
		Description: "Flips a random league of cups",
		Parameters: []api.Parameter{
			{
				Type:        api.ParameterInteger,
				Name:        cupsParam,
				Description: "Number of cups before the finals",
			},
			{
				Type:        api.ParameterBoolean,
				Name:        repeatableParam,
				Description: "Allow repeated cups",
			},
		},
		Handler: func(ctx context.Context, i api.Interaction) (api.Response, error) {
			filter := core.FlipLeagueFilter{
				Cups:       int(i.Integer(cupsParam)),
				Repeatable: i.Boolean(repeatableParam),
			}

			filter = filter.Normalize()

			league := core.FlipLeague(ctx, filter)

			content := fmt.Sprintf(
				leaguePattern,
				len(league),
				format.List(league),
			)

			return api.Response{Content: template.Flip("league", content)}, nil
		},
	}
}
