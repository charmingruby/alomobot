package cmd

import (
	"context"
	"fmt"

	"github.com/charmingruby/erebor/discord/command"

	"github.com/charmingruby/alomobot/internal/core"
)

func FlipGenDef() command.Definition {
	return command.Definition{
		Name:        "flip-gen",
		Description: "Flips a random Pokémon gen",
		Handler: func(ctx context.Context, _ command.Interaction) (string, error) {
			return fmt.Sprintf("%d", core.FlipGen(ctx).Number), nil
		},
	}
}
