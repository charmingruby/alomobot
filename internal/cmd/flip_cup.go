package cmd

import (
	"context"

	"github.com/charmingruby/alomobot/internal/core"
	"github.com/charmingruby/erebor/discord/command"
)

func FlipCupDef() command.Definition {
	return command.Definition{
		Name:        "flip-cup",
		Description: "Flips a random cup",
		Handler: func(ctx context.Context, _ command.Interaction) (string, error) {
			return core.FlipCup(ctx).String(), nil
		},
	}
}
