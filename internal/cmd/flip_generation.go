package cmd

import (
	"context"

	"github.com/charmingruby/alomobot/internal/core"
	"github.com/charmingruby/erebor/discord/command"
)

func FlipGenerationDef() command.Definition {
	return command.Definition{
		Name:        "flip-generation",
		Description: "Flips a random Pokémon generation",
		Handler: func(ctx context.Context, _ command.Interaction) (string, error) {
			return core.FlipGeneration(ctx).String(), nil
		},
	}
}
