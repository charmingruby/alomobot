package cmd

import (
	"context"

	"github.com/charmingruby/erebor/discord/command"
)

func FlipCupDef() command.Definition {
	return command.Definition{
		Name:        "flip-cup",
		Description: "Flips a random cup",
		Handler: func(_ context.Context, _ command.Interaction) (string, error) {
			return "todo", nil
		},
	}
}
