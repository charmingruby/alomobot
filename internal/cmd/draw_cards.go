package cmd

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/charmingruby/erebor/discord/api"

	"github.com/charmingruby/alomobot/internal/app"
	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/set"
	"github.com/charmingruby/alomobot/internal/core/pkmn"
	"github.com/charmingruby/alomobot/internal/format"
)

func (c *cmd) drawCards() api.Command {
	const (
		setParam       = "set"
		quantityParam  = "quantity"
		minRarityParam = "min_rarity"
		maxRarityParam = "max_rarity"
		typesParam     = "types"
		headingPattern = "🃏 Revealed **%d Pokémon**"
		cardsPattern   = "\n%s\n\n%s\n\nChoose and discard manually according to the Trainer's effect."
	)

	setChoices := make([]api.Choice, 0, len(set.Sets))
	for name := range set.Sets {
		setChoices = append(setChoices, api.Choice{
			Name:  string(name),
			Value: string(name),
		})
	}

	rarityChoices := make([]api.Choice, 0, len(booster.Rarities))
	for _, rarity := range booster.Rarities {
		rarityChoices = append(rarityChoices, api.Choice{
			Name:  string(rarity),
			Value: string(rarity),
		})
	}

	return api.Command{
		Name:        "draw",
		Description: "Reveals distinct Pokémon cards with rarity and type filters",
		Parameters: []api.Parameter{
			{
				Type:        api.ParameterString,
				Name:        setParam,
				Description: "Set to draw from",
				Required:    true,
				Choices:     setChoices,
			},
			{
				Type:        api.ParameterInteger,
				Name:        quantityParam,
				Description: "Number of distinct Pokémon to reveal (1-10)",
				Required:    true,
				MinValue:    new(1.0),
				MaxValue:    float64(booster.MaxDrawQuantity),
			},
			{
				Type:        api.ParameterString,
				Name:        minRarityParam,
				Description: "Minimum rarity (inclusive)",
				Choices:     rarityChoices,
			},
			{
				Type:        api.ParameterString,
				Name:        maxRarityParam,
				Description: "Maximum rarity (inclusive)",
				Choices:     rarityChoices,
			},
			{
				Type:        api.ParameterString,
				Name:        typesParam,
				Description: "Comma-separated types, e.g. Water,Bug; matches either Pokémon type",
			},
		},
		Handler: func(ctx context.Context, i api.Interaction) (api.Response, error) {
			filter := app.DrawCardsFilter{
				Set: i.String(setParam),
				DrawFilter: booster.DrawFilter{
					Quantity:  int(i.Integer(quantityParam)),
					MinRarity: booster.Rarity(i.String(minRarityParam)),
					MaxRarity: booster.Rarity(i.String(maxRarityParam)),
				},
			}

			types := make([]string, 0)
			if raw := i.String(typesParam); raw != "" {
				for kind := range strings.SplitSeq(raw, ",") {
					kind = format.Capitalize(strings.ToLower(strings.TrimSpace(kind)))
					filter.Types = append(filter.Types, pkmn.Type(kind))
					types = append(types, kind)
				}
			}

			result, err := c.app.DrawCards(ctx, filter)

			switch {
			case errors.Is(err, app.ErrSetNotFound):
				return api.Response{Content: "Set not found. Select one of the available sets."}, nil
			case errors.Is(err, booster.ErrInvalidDrawFilter):
				return api.Response{Content: "Invalid filters. Use quantity 1-10, minimum rarity no higher than maximum, and valid comma-separated Pokémon types in English (e.g. water,bug)."}, nil
			case errors.Is(err, booster.ErrInsufficientCards):
				return api.Response{Content: "Not enough distinct Pokémon match these filters. Reduce the quantity or broaden the filters."}, nil
			case err != nil:
				return api.Response{}, err
			}

			parameters := []string{fmt.Sprintf("Set `%s`", result.Set)}

			if filter.MinRarity != "" {
				parameters = append(parameters, fmt.Sprintf("Min rarity `%s`", filter.MinRarity))
			}

			if filter.MaxRarity != "" {
				parameters = append(parameters, fmt.Sprintf("Max rarity `%s`", filter.MaxRarity))
			}

			if len(filter.Types) > 0 {
				parameters = append(parameters, fmt.Sprintf("Types `%s`", strings.Join(types, ", ")))
			}

			sprites := make([]string, 0, len(result.Cards))
			for _, card := range result.Cards {
				sprites = append(sprites, card.PokemonAttr.SpriteURL)
			}

			content := fmt.Sprintf(headingPattern, len(result.Cards))
			content += fmt.Sprintf(cardsPattern, strings.Join(parameters, " • "), format.List(result.Cards))

			return api.Response{
				Content:   content,
				ImageURLs: sprites,
			}, nil
		},
	}
}
