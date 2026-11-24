package cmd

import (
	"context"

	"github.com/charmingruby/erebor/discord/api"

	"github.com/charmingruby/alomobot/internal/cmd/message"
	"github.com/charmingruby/alomobot/internal/cup"
)

func helpCup() api.Command {
	const nameParam = "name"

	choices := make([]api.Choice, 0, len(cup.Cups))
	for _, cup := range cup.Cups {
		choices = append(choices, api.Choice{
			Name:  cup.String(),
			Value: cup.String(),
		})
	}

	return api.Command{
		Name:        "cup",
		Description: "Explains how a cup works and how to build your team",
		Parameters: []api.Parameter{
			{
				Type:        api.ParameterString,
				Name:        nameParam,
				Description: "Cup you want to learn about",
				Required:    true,
				Choices:     choices,
			},
		},
		Handler: func(_ context.Context, i api.Interaction) (api.Response, error) {
			selected, ok := cup.FindCup(i.String(nameParam))
			if !ok {
				return api.Response{
					Content: "Cup not found. Select one of the available cups in `/help cup`.",
				}, nil
			}

			return api.Response{Content: message.HelpCup(selected)}, nil
		},
	}
}
