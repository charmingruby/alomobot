package cmd

import (
	"github.com/charmingruby/erebor/discord"
)

func Register(b *discord.Bot) error {
	if err := b.RegisterCommand(FlipDef()); err != nil {
		return err
	}

	return nil
}
