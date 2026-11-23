package cmd

import (
	"context"

	"github.com/charmingruby/erebor/discord/api"

	"github.com/charmingruby/alomobot/internal/cmd/message"
	"github.com/charmingruby/alomobot/internal/core"
)

func flipPokemon() api.Command {
	return api.Command{
		Name:        "pokemon",
		Description: "Flips a random Pokémon",
		Handler: func(ctx context.Context, _ api.Interaction) (api.Response, error) {
			poke, err := core.FlipPokemon(ctx)
			if err != nil {
				return api.Response{}, err
			}

			return api.Response{
				Content: message.Flip("Pokémon", poke.Render()),
				ImageURLs: []string{
					poke.SpriteURL,
				}}, nil
		},
	}
}
