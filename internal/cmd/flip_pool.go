package cmd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/charmingruby/erebor/discord/api"

	"github.com/charmingruby/alomobot/internal/app"
	"github.com/charmingruby/alomobot/internal/core/pool"
	"github.com/charmingruby/alomobot/internal/format"
)

func (c *cmd) flipPool() api.Command {
	choices := make([]api.Choice, 0, len(pool.Pools))
	for _, p := range pool.Pools {
		choices = append(choices, api.Choice{
			Name:  string(p.Name),
			Value: string(p.Name),
		})
	}

	return api.Command{
		Name:        "flip",
		Description: "Deals Pokémon from a registered or custom pool into named groups",
		Parameters: []api.Parameter{
			{
				Type:        api.ParameterString,
				Name:        "groups",
				Required:    true,
				Description: "Comma-separated group identifiers, e.g. G,F",
			},
			{
				Type:        api.ParameterString,
				Name:        "pool",
				Choices:     choices,
				Description: "Registered pool to draw from; cannot be used with custom",
			},
			{
				Type:        api.ParameterString,
				Name:        "custom",
				Description: "Comma-separated PokéAPI names; cannot be used with pool",
			},
			{
				Type:        api.ParameterInteger,
				Name:        "size",
				Description: "Number of Pokémon per group (default: 6)",
			},
		},
		Handler: func(ctx context.Context, i api.Interaction) (api.Response, error) {
			size := 6
			if i.GetOption("size") != nil {
				size = int(i.Integer("size"))
			}

			draw, err := c.app.FlipPool(ctx, app.FlipPoolFilter{
				Pool:   i.String("pool"),
				Custom: i.String("custom"),
				Groups: i.String("groups"),
				Size:   size,
			})
			if errors.Is(err, app.ErrInvalidPool) {
				return api.Response{Content: err.Error()}, nil
			}

			if err != nil {
				return api.Response{}, err
			}

			var content strings.Builder

			for _, group := range draw.Groups {
				fmt.Fprintf(&content, "**%s** (%d/%d)\n%s\n\n",
					group.Name,
					len(group.Pokemon),
					len(group.Pokemon),
					format.List(group.Pokemon),
				)
			}

			header := "🪙 A custom pool has been rolled!"
			if draw.Name != "" {
				header = fmt.Sprintf("🪙 A new **%s** pool has been rolled!", draw.Name)
			}

			return api.Response{
				Content: header + "\n\n" + content.String(),
			}, nil
		},
	}
}
