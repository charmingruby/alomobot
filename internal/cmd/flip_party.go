package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmingruby/erebor/discord/command"

	"github.com/charmingruby/alomobot/internal/core"
)

func FlipPartyDef() command.Definition {
	return command.Definition{
		Name:        "flip-party",
		Description: "Flips a random Pokémon party",
		Handler: func(ctx context.Context, _ command.Interaction) (string, error) {
			team, err := core.FlipParty(ctx)
			if err != nil {
				return "", err
			}

			var message strings.Builder
			for idx, p := range team {
				if idx > 0 {
					message.WriteString("\n")
				}

				fmt.Fprintf(&message, "%d. %s#%d", idx+1, p.Name, p.ID)
			}

			return message.String(), nil
		},
	}
}
