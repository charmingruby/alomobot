package cmd

import (
	"context"

	"github.com/charmingruby/erebor/discord/api"

	"github.com/charmingruby/alomobot/internal/cmd/template"
)

func (c *cmd) flipGen() api.Command {
	return api.Command{
		Name:        "flip",
		Description: "Flips a random Pokémon gen",
		Handler: func(ctx context.Context, _ api.Interaction) (api.Response, error) {
			return api.Response{Content: template.Flip("generation", c.app.FlipGen(ctx).Render())}, nil
		},
	}
}
