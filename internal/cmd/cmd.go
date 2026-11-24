package cmd

import "github.com/charmingruby/erebor/discord/api"

func Commands() []api.Command {
	return []api.Command{
		{
			Name:        "help",
			Description: "Explains how the formats work",
			Subcommands: []api.Command{
				helpCup(),
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
