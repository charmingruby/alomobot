package cmd

import "github.com/charmingruby/erebor/discord/api"

func Commands() []api.Command {
	return []api.Command{
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
