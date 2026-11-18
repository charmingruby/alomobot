package misc

import (
	"context"

	"github.com/charmingruby/alomobot/internal/pkmn"
	"github.com/charmingruby/erebor/discord/command"
)

func TypeFlipDef() command.Definition {
	return command.Definition{
		Name:        "type-flip",
		Description: "Flips a random Pokémon Type",
		Handler: func(ctx context.Context, i command.Interaction) (string, error) {
			return string(pkmn.RandomType()), nil
		},
	}
}
