package cmd

import (
	"context"

	"github.com/charmingruby/alomobot/internal/core"
	"github.com/charmingruby/erebor/discord/command"
)

var ()

func FlipDef() command.Definition {
	resouceParam := "resource"

	choices := []*command.Choice{
		{
			Name:  core.ResourceType.String(),
			Value: core.ResourceType,
		},
		{
			Name:  core.ResourceCup.String(),
			Value: core.ResourceCup,
		},
		{
			Name:  core.ResourcePokemon.String(),
			Value: core.ResourcePokemon,
		},
		{
			Name:  core.ResourceTeam.String(),
			Value: core.ResourceTeam,
		},
		{
			Name:  core.ResourceGeneration.String(),
			Value: core.ResourceGeneration,
		},
		{
			Name:  core.ResourceRegion.String(),
			Value: core.ResourceRegion,
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
					kind := core.Resource(opt.Value.(string))

					if !kind.Valid() {
						return "", ErrNonexistentKind
					}

					println(kind)

					switch kind {
					case core.ResourceType:
						return string(core.RandomType()), nil
					case core.ResourceCup:
						return "todo", nil
					case core.ResourcePokemon:
						return "todo", nil
					case core.ResourceTeam:
						return "todo", nil
					case core.ResourceGeneration:
						return "todo", nil
					case core.ResourceRegion:
						return "todo", nil
					}
				}
			}

			return "", ErrInvalidParam
		},
	}

}
