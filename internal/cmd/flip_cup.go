package cmd

import (
	"context"

	"github.com/charmingruby/erebor/discord/command"

	"github.com/charmingruby/alomobot/internal/cmd/message"
	"github.com/charmingruby/alomobot/internal/core"
)

func FlipCupDef() command.Definition {
	return command.Definition{
		Name:        "flip-cup",
		Description: "Flips a random cup",
		Handler: func(ctx context.Context, _ command.Interaction) (string, error) {
			return message.Flip("cup", core.FlipCup(ctx).String()), nil
		},
	}
}
