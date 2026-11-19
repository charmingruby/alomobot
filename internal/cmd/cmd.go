package cmd

import (
	"github.com/charmingruby/erebor/discord"
	"github.com/charmingruby/erebor/discord/command"
)

func Register(b *discord.Bot) error {
	definitions := []command.Definition{
		FlipTypeDef(),
		FlipCupDef(),
		FlipPokemonDef(),
		FlipPartyDef(),
		FlipGenDef(),
		FlipLeagueDef(),
	}

	for _, def := range definitions {
		if err := b.RegisterCommand(def); err != nil {
			return err
		}
	}

	return nil
}
