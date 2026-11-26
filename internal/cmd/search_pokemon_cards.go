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
	"github.com/charmingruby/alomobot/internal/core/pkmn"
	"github.com/charmingruby/alomobot/internal/format"
)

func (c *cmd) searchPokemonCards() api.Command {
	choices := make([]api.Choice, 0, len(set.Sets))
	for name := range set.Sets {
		choices = append(choices, api.Choice{Name: string(name), Value: string(name)})
	}

	types := make([]api.Choice, 0, len(pkmn.Types))
	for _, kind := range pkmn.Types {
		types = append(types, api.Choice{Name: format.Capitalize(kind.String()), Value: kind.String()})
	}

	return api.Command{
		Name:        "pokemon",
		Description: "Lists or reads Pokémon cards, filtered by rarity and typing",
		Parameters: []api.Parameter{
			{
				Type: api.ParameterString, Name: "set",
				Description: "Set to browse", Required: true, Choices: choices,
			},
			{
				Type: api.ParameterString, Name: "name",
				Description: "Exact card name to read; omit to list cards",
			},
			{
				Type: api.ParameterString, Name: "rarity",
				Description: "Card rarity; omit to show all rarities",
				Choices: []api.Choice{
					{Name: "Common", Value: string(booster.CommonRarity)},
					{Name: "Uncommon", Value: string(booster.UncommonRarity)},
					{Name: "Rare", Value: string(booster.RareRarity)},
					{Name: "Epic", Value: string(booster.EpicRarity)},
					{Name: "Legendary", Value: string(booster.LegendaryRarity)},
				},
			},
			{
				Type: api.ParameterInteger, Name: "page",
				Description: "Page number (25 cards per page); ignored when reading a named card", MinValue: new(1.0),
			},
			{
				Type: api.ParameterString, Name: "type",
				Description: "Pokémon typing; matches either primary or secondary type", Choices: types,
			},
		},
		Handler: func(ctx context.Context, i api.Interaction) (api.Response, error) {
			result, err := c.app.SearchPokemonCards(ctx, app.SearchPokemonCardsFilter{
				Set:    i.String("set"),
				Name:   i.String("name"),
				Rarity: booster.Rarity(i.String("rarity")),
				Page:   int(i.Integer("page")),
				Type:   pkmn.Type(i.String("type")),
			})

			switch {
			case errors.Is(err, app.ErrSetNotFound):
				return api.Response{Content: "Set not found. Select one of the available sets."}, nil
			case errors.Is(err, app.ErrPokemonCardNotFound):
				return api.Response{Content: "Card not found with these filters. Omit `name` to list the available cards."}, nil
			case errors.Is(err, app.ErrPokemonCardPageNotFound):
				return api.Response{Content: "Page not found. Use `page:1` to see the available pages."}, nil
			case err != nil:
				return api.Response{}, err
			}

			if result.Total == 0 {
				return api.Response{Content: fmt.Sprintf("No Pokémon cards found for these filters in **%s**.", result.Set)}, nil
			}

			if result.Detailed {
				card := result.Cards[0]
				response := api.Response{
					Content: fmt.Sprintf("🃏 **%s** — Pokémon\n\n%s", result.Set, card.Render()),
				}

				if card.Type == booster.PokemonCardType {
					response.ImageURLs = []string{card.PokemonAttr.SpriteURL}
				}

				return response, nil
			}

			parameters := []string{fmt.Sprintf("Set `%s`", result.Set)}

			if rarity := i.String("rarity"); rarity != "" {
				parameters = append(parameters, fmt.Sprintf("Rarity `%s`", rarity))
			}

			if kind := i.String("type"); kind != "" {
				parameters = append(parameters, fmt.Sprintf("Type `%s`", kind))
			}

			return api.Response{
				Content: template.SearchCards(strings.Join(parameters, " • "), result.Cards, result.Total, result.Page, result.Pages),
			}, nil
		},
	}
}
