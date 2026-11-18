package misc

import (
	"github.com/charmingruby/erebor/discord"
)

func RegisterCommands(b *discord.Bot) error {
	if err := b.RegisterCommand(TypeFlipDef()); err != nil {
		return err
	}

	return nil
}
