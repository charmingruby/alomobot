package cmd

import (
	"github.com/charmingruby/alomobot/internal/pkmn/pokeapi"
	"github.com/charmingruby/erebor/discord"
)

func Register(b *discord.Bot) error {
	api := pokeapi.New()

	if err := b.RegisterCommand(FlipDef(api)); err != nil {
		return err
	}

	return nil
}
