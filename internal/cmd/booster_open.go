package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/charmingruby/erebor/discord/api"

	"github.com/charmingruby/alomobot/internal/app"
	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/set"
	"github.com/charmingruby/alomobot/internal/format"
)

func (c *cmd) boosterOpen() api.Command {
	const (
		setParam = "set"
		pattern  = "📦 Opened a booster from **%s** set!\n\n**Cards** (%d)\n%s"
	)

	choices := make([]api.Choice, 0, len(set.Sets))
	for name := range set.Sets {
		choices = append(choices, api.Choice{
			Name:  string(name),
			Value: string(name),
		})
	}

	return api.Command{
		Name:        "open",
		Description: "Opens a booster from the selected set",
		Parameters: []api.Parameter{
			{
				Type:        api.ParameterString,
				Name:        setParam,
				Description: "Set to open a booster from",
				Required:    true,
				Choices:     choices,
			},
		},
		Handler: func(ctx context.Context, i api.Interaction) (api.Response, error) {
			result, err := c.app.OpenBooster(ctx, i.String(setParam))
			if errors.Is(err, app.ErrSetNotFound) {
				return api.Response{
					Content: "Set not found. Select one of the available sets in `/booster open`.",
				}, nil
			}
			if err != nil {
				return api.Response{}, err
			}

			sprites := make([]string, 0, len(result.Cards))
			for _, card := range result.Cards {
				if card.Type == booster.PokemonCardType {
					sprites = append(sprites, card.PokemonAttr.SpriteURL)
				}
			}

			return api.Response{
				Content:   fmt.Sprintf(pattern, result.Set.Name, len(result.Cards), format.List(result.Cards)),
				ImageURLs: sprites,
			}, nil
		},
	}
}
