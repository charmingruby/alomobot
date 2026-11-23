package cmd

import (
	"context"

	"github.com/charmingruby/erebor/discord/command"

	"github.com/charmingruby/alomobot/internal/core"
)

func FlipPokemonDef() command.Definition {
	return command.Definition{
		Name:        "flip-pokemon",
		Description: "Flips a random Pokémon",
		Handler: func(ctx context.Context, _ command.Interaction) (string, error) {
			poke, err := core.FlipPokemon(ctx)
			if err != nil {
				return "", err
			}

			return poke.Render(), nil
		},
	}
}
