package cmd

import (
	"context"
	"fmt"

	"github.com/charmingruby/alomobot/internal/core"
	"github.com/charmingruby/erebor/discord/command"
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

			return fmt.Sprintf("%s#%d", poke.Name, poke.ID), nil
		},
	}
}
