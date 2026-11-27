package pool

import (
	"maps"
	"slices"

	"github.com/charmingruby/alomobot/internal/core/pkmn"
)

type PoolName string

// Pool groups Pokémon entries by contributor. Repeated entries are separate copies.
type Pool struct {
	Name     PoolName
	Subpools map[string][]pkmn.Pokemon
}

var Pools = []Pool{
	favsPool,
}

func (p Pool) Entries() []pkmn.Pokemon {
	var entries []pkmn.Pokemon
	for _, key := range slices.Sorted(maps.Keys(p.Subpools)) {
		entries = append(entries, p.Subpools[key]...)
	}

	return entries
}

func FindPool(name string) (Pool, bool) {
	idx := slices.IndexFunc(Pools, func(p Pool) bool {
		return p.Name == PoolName(name)
	})
	if idx == -1 {
		return Pool{}, false
	}

	return Pools[idx], true
}
