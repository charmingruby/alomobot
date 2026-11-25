package cmd

import (
	"github.com/charmingruby/alomobot/internal/app"
	"github.com/charmingruby/erebor/discord/api"
)

type cmd struct {
	app *app.App
}

func Commands(app *app.App) []api.Command {
	c := cmd{
		app: app,
	}

	return []api.Command{
		{
			Name:        "rules",
			Description: "Explains tournament rules and formats",
			Subcommands: []api.Command{
				c.rulesGlobal(),
				c.rulesCup(),
			},
		},
		{
			Name:        "flip",
			Description: "Flips random Pokémon resources",
			Subcommands: []api.Command{
				c.flipType(),
				c.flipCup(),
				c.flipPokemon(),
				c.flipParty(),
				c.flipGen(),
				c.flipLeague(),
			},
		},
	}
}
