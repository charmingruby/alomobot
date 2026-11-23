package cmd

import (
	"context"

	"github.com/charmingruby/erebor/discord/api"

	"github.com/charmingruby/alomobot/internal/cmd/message"
	"github.com/charmingruby/alomobot/internal/core"
)

func flipGen() api.Command {
	return api.Command{
		Name:        "gen",
		Description: "Flips a random Pokémon gen",
		Handler: func(ctx context.Context, _ api.Interaction) (api.Response, error) {
			return api.Response{Content: message.Flip("generation", core.FlipGen(ctx).Render())}, nil
		},
	}
}
