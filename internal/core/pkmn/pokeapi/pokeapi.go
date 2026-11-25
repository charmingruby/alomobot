package pokeapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/charmingruby/erebor/httpx"
	"github.com/charmingruby/erebor/o11y/logging"

	"github.com/charmingruby/alomobot/internal/core/pkmn"
	"github.com/charmingruby/alomobot/internal/core/pkmn/pokeapi/data"
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

func (p *PokeAPI) GetPokemonByDexID(ctx context.Context, dexID int) (pkmn.Pokemon, error) {
	poke, err := httpx.Do[data.Pokemon](ctx, p.client, httpx.Request{
		URL:            fmt.Sprintf("/pokemon/%d", dexID),
		ExpectedStatus: http.StatusOK,
	})

	if err != nil {
		logging.LoggerFromContext(ctx).ErrorContext(ctx, "pokemon not found by dex id",
			"error", err,
		)

		return pkmn.Pokemon{}, fmt.Errorf("%w: %w", ErrPokemonNotFound, err)
	}

	var primaryType pkmn.Type
	var secondaryType *pkmn.Type

	for _, t := range poke.Types {
		if t.Slot == 1 {
			primaryType = pkmn.Type(t.Type.Name)
			continue
		}

		st := pkmn.Type(t.Type.Name)
		secondaryType = &st
	}

	return pkmn.Pokemon{
		ID:            poke.ID,
		Name:          poke.Name,
		PrimaryType:   primaryType,
		SecondaryType: secondaryType,
		SpriteURL:     poke.Sprites.Other.OfficialArtwork.FrontDefault,
	}, nil
}
