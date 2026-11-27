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
			Name:        "booster",
			Description: "Opens boosters and searches set cards",
			Subcommands: []api.Command{
				c.openBooster(),
				c.drawCards(),
				{
					Name:        "search",
					Description: "Searches cards by set, name, rarity and type",
					Subcommands: []api.Command{
						c.searchPokemonCards(),
						c.searchSupportCards(),
						c.searchEnergyCards(),
					},
				},
			},
		},
		{
			Name:        "rules",
			Description: "Explains tournament rules and formats",
			Subcommands: []api.Command{
				c.rulesGlobal(),
				c.rulesCup(),
			},
		},
		{
			Name:        "pokemon",
			Description: "Pokémon commands",
			Subcommands: []api.Command{
				c.flipPokemon(),
			},
		},
		{
			Name:        "type",
			Description: "Pokémon type commands",
			Subcommands: []api.Command{
				c.flipType(),
			},
		},
		{
			Name:        "cup",
			Description: "Tournament cup commands",
			Subcommands: []api.Command{
				c.flipCup(),
			},
		},
		{
			Name:        "party",
			Description: "Pokémon party commands",
			Subcommands: []api.Command{
				c.flipParty(),
			},
		},
		{
			Name:        "gen",
			Description: "Pokémon generation commands",
			Subcommands: []api.Command{
				c.flipGen(),
			},
		},
		{
			Name:        "league",
			Description: "Tournament league commands",
			Subcommands: []api.Command{
				c.flipLeague(),
			},
		},
	}
}
