package cmd

import (
	"github.com/charmingruby/alomobot/internal/pkmn/pokeapi"
	"github.com/charmingruby/erebor/discord"
	"github.com/charmingruby/erebor/discord/command"
)

func Register(b *discord.Bot) error {
	api := pokeapi.New()

	definitions := []command.Definition{
		FlipTypeDef(),
		FlipCupDef(),
		FlipPokemonDef(api),
		FlipPartyDef(api),
		FlipGenerationDef(),
		FlipRegionDef(),
		FlipLeagueDef(),
	}

	for _, def := range definitions {
		if err := b.RegisterCommand(def); err != nil {
			return err
		}
	}

	return nil
}
