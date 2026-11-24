package cmd

import (
	"context"

	"github.com/charmingruby/erebor/discord/api"

	"github.com/charmingruby/alomobot/internal/cmd/template"
	"github.com/charmingruby/alomobot/internal/core"
)

func flipType() api.Command {
	return api.Command{
		Name:        "type",
		Description: "Flips a random Pokémon type",
		Handler: func(ctx context.Context, _ api.Interaction) (api.Response, error) {
			return api.Response{Content: template.Flip("type", core.FlipType(ctx).Render())}, nil
		},
	}
}
