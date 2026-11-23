package cmd

import (
	"fmt"

	"github.com/charmingruby/erebor/discord"
	"github.com/charmingruby/erebor/discord/command"
)

func Register(b *discord.Bot) error {
	defs := []command.Definition{
		flipType(),
		flipCup(),
		flipPokemon(),
		flipParty(),
		flipGen(),
		flipLeague(),
	}

	for _, def := range defs {
		if err := b.RegisterCommand(def); err != nil {
			return fmt.Errorf("register command %q: %w", def.Name, err)
		}
	}

	return nil
}
