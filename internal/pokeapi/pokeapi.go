package pokeapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/charmingruby/alomobot/internal/core"
	"github.com/charmingruby/alomobot/internal/pokeapi/data"
	"github.com/charmingruby/erebor/httpx"
	"github.com/charmingruby/erebor/o11y/logging"
)

const url = "https://pokeapi.co/api/v2"

var (
	ErrPokemonNotFound = errors.New("pokemon not found")
)

type PokeAPI struct {
	client *httpx.Client
}

func New() *PokeAPI {
	return &PokeAPI{
		client: httpx.NewClient(httpx.WithBaseURL(url)),
	}
}

func (p *PokeAPI) GetPokemonByDexID(ctx context.Context, dexID int) (*core.Pokemon, error) {
	poke, err := httpx.Do[data.Pokemon](ctx, p.client, httpx.Request{
		URL:            fmt.Sprintf("/pokemon/%d", dexID),
		ExpectedStatus: http.StatusOK,
	})

	if err != nil {
		logging.LoggerFromContext(ctx).ErrorContext(ctx, "pokemon not found by dex id",
			"error", err,
		)

		return nil, fmt.Errorf("%w: %w", ErrPokemonNotFound, err)
	}

	return &core.Pokemon{
		ID:        poke.ID,
		Name:      poke.Name,
		SpriteURL: poke.Sprites.FrontDefault,
	}, nil
}
