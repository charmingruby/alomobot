package cmd

import (
	"context"

	"github.com/charmingruby/erebor/discord/command"

	"github.com/charmingruby/alomobot/internal/cmd/message"
	"github.com/charmingruby/alomobot/internal/core"
)

func flipGen() command.Definition {
	return command.Definition{
		Name:        "flip-gen",
		Description: "Flips a random Pokémon gen",
		Handler: func(ctx context.Context, _ command.Interaction) (string, error) {
			return message.Flip("generation", core.FlipGen(ctx).Render()), nil
		},
	}
}
