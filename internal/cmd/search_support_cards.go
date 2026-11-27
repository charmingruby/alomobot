package cmd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/charmingruby/erebor/discord/api"

	"github.com/charmingruby/alomobot/internal/app"
	"github.com/charmingruby/alomobot/internal/cmd/template"
	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/set"
)

func (c *cmd) searchSupportCards() api.Command {
	setChoices := make([]api.Choice, 0, len(set.Sets))
	for name := range set.Sets {
		setChoices = append(setChoices, api.Choice{Name: string(name), Value: string(name)})
	}

	rarityChoices := make([]api.Choice, 0, len(booster.Rarities))
	for _, rarity := range booster.Rarities {
		rarityChoices = append(rarityChoices, api.Choice{
			Name:  string(rarity),
			Value: string(rarity),
		})
	}

	return api.Command{
		Name:        "support",
		Description: "Lists or reads Support cards and their effects",
		Parameters: []api.Parameter{
			{
				Type: api.ParameterString, Name: "set",
				Description: "Set to browse", Required: true, Choices: setChoices,
			},
			{
				Type: api.ParameterString, Name: "name",
				Description: "Exact card name to read; omit to list cards",
			},
			{
				Type: api.ParameterString, Name: "rarity",
				Description: "Card rarity; omit to show all rarities",
				Choices:     rarityChoices,
			},
			{
				Type: api.ParameterInteger, Name: "page",
				Description: "Page number (25 cards per page); ignored when reading a named card", MinValue: new(1.0),
			},
		},
		Handler: func(ctx context.Context, i api.Interaction) (api.Response, error) {
			result, err := c.app.SearchSupportCards(ctx, app.SearchSupportCardsFilter{
				Set: i.String("set"), Name: i.String("name"), Rarity: booster.Rarity(i.String("rarity")),
				Page: int(i.Integer("page")),
			})

			switch {
			case errors.Is(err, app.ErrSetNotFound):
				return api.Response{Content: "Set not found. Select one of the available sets."}, nil
			case errors.Is(err, app.ErrSupportCardNotFound):
				return api.Response{Content: "Card not found with these filters. Omit `name` to list the available cards."}, nil
			case errors.Is(err, app.ErrSupportCardPageNotFound):
				return api.Response{Content: "Page not found. Use `page:1` to see the available pages."}, nil
			case err != nil:
				return api.Response{}, err
			}

			if result.Total == 0 {
				return api.Response{Content: fmt.Sprintf("No Support cards found for these filters in **%s**.", result.Set)}, nil
			}

			if result.Detailed {
				return api.Response{
					Content: fmt.Sprintf("🃏 **%s** — Support\n\n%s", result.Set, result.Cards[0].Render()),
				}, nil
			}

			parameters := []string{fmt.Sprintf("Set `%s`", result.Set)}
			if rarity := i.String("rarity"); rarity != "" {
				parameters = append(parameters, fmt.Sprintf("Rarity `%s`", rarity))
			}

			return api.Response{
				Content: template.SearchCards(strings.Join(parameters, " • "), result.Cards, result.Total, result.Page, result.Pages),
			}, nil
		},
	}
}
