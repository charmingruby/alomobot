package app

import (
	"context"

	"github.com/charmingruby/alomobot/internal/core/pkmn"
	"github.com/charmingruby/alomobot/internal/core/pkmn/pokeapi"
)

type pokemonAPI interface {
	GetPokemonByDexID(context.Context, int) (pkmn.Pokemon, error)
}

type App struct {
	api pokemonAPI
}

func New(api *pokeapi.PokeAPI) *App {
	return &App{
		api: api,
	}
}
