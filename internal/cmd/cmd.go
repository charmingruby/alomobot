package cmd

import "github.com/charmingruby/erebor/discord/api"

func Commands() []api.Command {
	return []api.Command{
		{
			Name:        "rules",
			Description: "Explains tournament rules and formats",
			Subcommands: []api.Command{
				rulesGlobal(),
				rulesCup(),
			},
		},
		{
			Name:        "flip",
			Description: "Flips random Pokémon resources",
			Subcommands: []api.Command{
				flipType(),
				flipCup(),
				flipPokemon(),
				flipParty(),
				flipGen(),
				flipLeague(),
			},
		},
	}
}
