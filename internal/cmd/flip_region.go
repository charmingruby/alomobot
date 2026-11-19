package cmd

import (
	"context"

	"github.com/charmingruby/alomobot/internal/core"
	"github.com/charmingruby/erebor/discord/command"
)

func FlipRegionDef() command.Definition {
	return command.Definition{
		Name:        "flip-region",
		Description: "Flips a random Pokémon region",
		Handler: func(ctx context.Context, _ command.Interaction) (string, error) {
			return core.FlipRegion(ctx).String(), nil
		},
	}
}
