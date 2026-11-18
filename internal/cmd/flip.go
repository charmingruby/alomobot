package cmd

import (
	"context"

	"github.com/charmingruby/alomobot/internal/core"
	"github.com/charmingruby/erebor/discord/command"
)

var ()

func FlipDef() command.Definition {
	resouceParam := "resource"

	return command.Definition{
		Name:        "flip",
		Description: "Flips a random resource",
		Parameters: []*command.Parameter{
			command.String(command.Parameter{
				Name:        resouceParam,
				Description: "Kind of resource",
				Required:    true,
				Choices: []*command.Choice{
					{
						Name:  "type",
						Value: core.TypeResouceKind,
					},
					{
						Name:  "cup",
						Value: core.CupResouceKind,
					},
				},
			}),
		},
		Handler: func(ctx context.Context, i command.Interaction) (string, error) {
			for _, opt := range i.Options {
				if opt.Name == resouceParam {
					kind := core.ResourceKind(opt.Value.(string))

					if !kind.Valid() {
						return "", ErrNonexistentKind
					}

					switch kind {
					case core.TypeResouceKind:
						return string(core.RandomType()), nil
					case core.CupResouceKind:
						return "todo", nil
					}
				}
			}

			return "", ErrInvalidParam
		},
	}

}
