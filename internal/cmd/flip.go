package cmd

import (
	"context"

	"github.com/charmingruby/alomobot/internal/core"
	"github.com/charmingruby/alomobot/internal/pkmn"
	"github.com/charmingruby/erebor/discord/command"
)

var ()

func FlipDef() command.Definition {
	resouceParam := "resource"

	choices := []*command.Choice{
		{
			Name:  pkmn.ResourceType.String(),
			Value: pkmn.ResourceType,
		},
		{
			Name:  pkmn.ResourceCup.String(),
			Value: pkmn.ResourceCup,
		},
		{
			Name:  pkmn.ResourcePokemon.String(),
			Value: pkmn.ResourcePokemon,
		},
		{
			Name:  pkmn.ResourceTeam.String(),
			Value: pkmn.ResourceTeam,
		},
		{
			Name:  pkmn.ResourceGeneration.String(),
			Value: pkmn.ResourceGeneration,
		},
		{
			Name:  pkmn.ResourceRegion.String(),
			Value: pkmn.ResourceRegion,
		},
	}

	return command.Definition{
		Name:        "flip",
		Description: "Flips a random resource",
		Parameters: []*command.Parameter{
			command.String(command.Parameter{
				Name:        resouceParam,
				Description: "Kind of resource",
				Required:    true,
				Choices:     choices,
			}),
		},
		Handler: func(ctx context.Context, i command.Interaction) (string, error) {
			for _, opt := range i.Options {
				if opt.Name == resouceParam {
					kind := pkmn.Resource(opt.Value.(string))

					if !kind.Valid() {
						return "", ErrNonexistentKind
					}

					switch kind {
					case pkmn.ResourceType:
						return core.FlipType().String(), nil
					case pkmn.ResourceCup:
						return "todo", nil
					case pkmn.ResourcePokemon:
						return "todo", nil
					case pkmn.ResourceTeam:
						return "todo", nil
					case pkmn.ResourceGeneration:
						return core.FlipGeneration().String(), nil
					case pkmn.ResourceRegion:
						return core.FlipRegion().String(), nil
					}
				}
			}

			return "", ErrInvalidParam
		},
	}

}
