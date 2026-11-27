package app

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
	"unicode"

	"github.com/charmingruby/alomobot/internal/core/pkmn"
	"github.com/charmingruby/alomobot/internal/core/pkmn/pokeapi"
	"github.com/charmingruby/alomobot/internal/core/pool"
)

var ErrInvalidPool = errors.New("invalid pool draw")

type FlipPoolFilter struct {
	Pool   string
	Custom string
	Groups string
	Size   int
}

type PoolGroup struct {
	Name    string
	Pokemon []pkmn.Pokemon
}

type FlipPoolResult struct {
	Name   string
	Groups []PoolGroup
}

func (a *App) FlipPool(ctx context.Context, filter FlipPoolFilter) (FlipPoolResult, error) {
	name, custom := strings.TrimSpace(filter.Pool), strings.TrimSpace(filter.Custom)
	if (name == "") == (custom == "") {
		return FlipPoolResult{}, fmt.Errorf("%w: provide exactly one of pool or custom", ErrInvalidPool)
	}

	if filter.Size < 1 {
		return FlipPoolResult{}, fmt.Errorf("%w: size must be at least 1", ErrInvalidPool)
	}

	var entries []pkmn.Pokemon

	if name != "" {
		selected, ok := pool.FindPool(name)
		if !ok {
			return FlipPoolResult{}, fmt.Errorf("%w: unknown pool; choose a registered pool", ErrInvalidPool)
		}

		entries = selected.Entries()
	} else {
		for entry := range strings.SplitSeq(custom, ",") {
			entry = strings.ToLower(strings.TrimSpace(entry))
			if entry == "" {
				return FlipPoolResult{}, fmt.Errorf("%w: custom must contain comma-separated Pokémon names without empty entries", ErrInvalidPool)
			}

			entries = append(entries, pkmn.Pokemon{Name: entry})
		}
	}

	var groups []PoolGroup
	seen := map[string]bool{}

	for group := range strings.SplitSeq(filter.Groups, ",") {
		group = strings.TrimSpace(group)
		if group == "" {
			return FlipPoolResult{}, fmt.Errorf("%w: groups must contain comma-separated identifiers without empty entries", ErrInvalidPool)
		}

		if len(group) > 32 || strings.ContainsFunc(group, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '-' && r != '_'
		}) {
			return FlipPoolResult{}, fmt.Errorf("%w: group identifiers must be up to 32 characters using letters, numbers, - or _", ErrInvalidPool)
		}

		key := strings.ToLower(group)
		if seen[key] {
			return FlipPoolResult{}, fmt.Errorf("%w: group identifiers must be unique", ErrInvalidPool)
		}

		seen[key] = true
		groups = append(groups, PoolGroup{Name: group})
	}

	if filter.Size > len(entries)/len(groups) {
		return FlipPoolResult{}, fmt.Errorf("%w: not enough Pokémon: %d entries available for %d groups of %d",
			ErrInvalidPool, len(entries), len(groups), filter.Size)
	}

	// Resolve every entry before drawing, so invalid custom names cannot be
	// silently ignored just because they were not selected.
	cache := map[string]pkmn.Pokemon{}

	for i, entry := range entries {
		identifier := entry.Name
		if entry.ID > 0 {
			identifier = strconv.Itoa(entry.ID)
		}

		if p, ok := cache[identifier]; ok {
			entries[i] = p
			continue
		}

		p, err := a.api.GetPokemon(ctx, identifier)
		if errors.Is(err, pokeapi.ErrPokemonNotFound) {
			return FlipPoolResult{}, fmt.Errorf("%w: could not resolve Pokémon %q; use PokéAPI names such as ninetales-alola or slowpoke-galar",
				ErrInvalidPool, identifier)
		}

		if err != nil {
			return FlipPoolResult{}, err
		}

		cache[identifier] = p
		entries[i] = p
	}

	rand.Shuffle(len(entries), func(i, j int) {
		entries[i], entries[j] = entries[j], entries[i]
	})

	for i := range groups {
		start := i * filter.Size
		groups[i].Pokemon = entries[start : start+filter.Size]
	}

	return FlipPoolResult{
		Name:   name,
		Groups: groups,
	}, nil
}
