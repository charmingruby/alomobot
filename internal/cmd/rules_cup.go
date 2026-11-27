package cmd

import (
	"context"
	"errors"

	"github.com/charmingruby/erebor/discord/api"

	"github.com/charmingruby/alomobot/internal/app"
	"github.com/charmingruby/alomobot/internal/cmd/template"
	"github.com/charmingruby/alomobot/internal/core/cup"
)

func (c *cmd) rulesCup() api.Command {
	const nameParam = "name"

	choices := make([]api.Choice, 0, len(cup.Cups))
	for name := range cup.Cups {
		choices = append(choices, api.Choice{
			Name:  string(name),
			Value: string(name),
		})
	}

	return api.Command{
		Name:        "rules",
		Description: "Shows global rules or the rules for a selected cup",
		Parameters: []api.Parameter{
			{
				Type:        api.ParameterString,
				Name:        nameParam,
				Description: "Cup you want to learn about; omit to see global rules",
				Required:    false,
				Choices:     choices,
			},
		},
		Handler: func(ctx context.Context, i api.Interaction) (api.Response, error) {
			name := i.String(nameParam)
			if name == "" {
				return api.Response{Content: template.RulesGlobal()}, nil
			}

			selected, err := c.app.RulesCup(ctx, name)
			if errors.Is(err, app.ErrCupNotFound) {
				return api.Response{
					Content: "Cup not found. Select one of the available cups in `/cup rules`.",
				}, nil
			}

			if err != nil {
				return api.Response{}, err
			}

			return api.Response{Content: template.RulesCup(selected)}, nil
		},
	}
}
