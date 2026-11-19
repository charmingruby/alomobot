package cmd

import (
	"context"

	"github.com/charmingruby/alomobot/internal/core"
	"github.com/charmingruby/erebor/discord/command"
)

func FlipTypeDef() command.Definition {
	return command.Definition{
		Name:        "flip-type",
		Description: "Flips a random Pokémon type",
		Handler: func(ctx context.Context, _ command.Interaction) (string, error) {
			return core.FlipType(ctx).String(), nil
		},
	}
}
