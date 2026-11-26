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
		pattern  = "📦 Opened a booster from **%s** set!"
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
			var pokemonCards, supportCards, energyCards []booster.Card
			for _, card := range result.Cards {
				switch card.Type {
				case booster.PokemonCardType:
					pokemonCards = append(pokemonCards, card)
					sprites = append(sprites, card.PokemonAttr.SpriteURL)
				case booster.SupportCardType:
					supportCards = append(supportCards, card)
				case booster.EnergyCardType:
					energyCards = append(energyCards, card)
				}
			}

			content := fmt.Sprintf(pattern, result.Set.Name)

			if len(pokemonCards) > 0 {
				content += fmt.Sprintf("\n\n**Pokémon** (%d)\n%s", len(pokemonCards), format.List(pokemonCards))
			}

			if len(supportCards) > 0 {
				content += fmt.Sprintf("\n\n**Support** (%d)\n%s", len(supportCards), format.List(supportCards))
			}

			if len(energyCards) > 0 {
				content += fmt.Sprintf("\n\n**Energy** (%d)\n%s", len(energyCards), format.List(energyCards))
			}

			return api.Response{
				Content:   content,
				ImageURLs: sprites,
			}, nil
		},
	}
}
