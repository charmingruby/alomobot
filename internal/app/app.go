package app

import "github.com/charmingruby/alomobot/internal/core/pkmn/pokeapi"

type App struct {
	api *pokeapi.PokeAPI
}

func New(api *pokeapi.PokeAPI) *App {
	return &App{
		api: api,
	}
}
