package cmd

import (
	"context"

	"github.com/charmingruby/erebor/discord/api"

	"github.com/charmingruby/alomobot/internal/cmd/message"
	"github.com/charmingruby/alomobot/internal/core"
)

func flipCup() api.Command {
	return api.Command{
		Name:        "cup",
		Description: "Flips a random cup",
		Handler: func(ctx context.Context, _ api.Interaction) (api.Response, error) {
			return api.Response{Content: message.Flip("cup", core.FlipCup(ctx).String())}, nil
		},
	}
}
