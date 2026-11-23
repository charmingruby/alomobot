package cmd

import (
	"context"

	"github.com/charmingruby/erebor/discord/command"

	"github.com/charmingruby/alomobot/internal/cmd/message"
	"github.com/charmingruby/alomobot/internal/core"
)

func FlipTypeDef() command.Definition {
	return command.Definition{
		Name:        "flip-type",
		Description: "Flips a random Pokémon type",
		Handler: func(ctx context.Context, _ command.Interaction) (string, error) {
			return message.Flip("type", core.FlipType(ctx).Render()), nil
		},
	}
}
