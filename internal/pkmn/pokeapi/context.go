package pokeapi

import (
	"context"
	"sync"
)

var defaultPokeAPI = sync.OnceValue(New)

type pokeAPICtxKey struct{}

func WithPokeAPI(ctx context.Context, api *PokeAPI) context.Context {
	return context.WithValue(ctx, pokeAPICtxKey{}, api)
}

func PokeAPIFromContext(ctx context.Context) *PokeAPI {
	if ctx != nil {
		if api, ok := ctx.Value(pokeAPICtxKey{}).(*PokeAPI); ok && api != nil {
			return api
		}
	}

	return defaultPokeAPI()
}
