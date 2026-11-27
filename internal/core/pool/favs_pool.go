package pool

import "github.com/charmingruby/alomobot/internal/core/pkmn"

const FavsPool PoolName = "Favs"

var favsPool = Pool{
	Name: FavsPool,
	Subpools: map[string][]pkmn.Pokemon{
		"G": {
			{Name: "incineroar", ID: 727},
			{Name: "scizor", ID: 212},
			{Name: "gligar", ID: 207},
			{Name: "slowpoke", ID: 79},
			{Name: "alomomola", ID: 594},
			{Name: "reshiram", ID: 643},
		},
		"F": {
			{Name: "tyranitar", ID: 248},
			{Name: "garchomp", ID: 445},
			{Name: "blaziken", ID: 257},
			{Name: "excadrill", ID: 530},
			{Name: "ninetales-alola", ID: 10104},
			{Name: "rayquaza", ID: 384},
		},
	},
}
