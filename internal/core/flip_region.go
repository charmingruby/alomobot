package core

import (
	"math/rand"

	"github.com/charmingruby/alomobot/internal/pkmn"
)

var regions = []pkmn.Region{
	pkmn.RegionKanto,
	pkmn.RegionJohto,
	pkmn.RegionHoenn,
	pkmn.RegionSinnoh,
	pkmn.RegionUnova,
	pkmn.RegionKalos,
	pkmn.RegionAlola,
	pkmn.RegionGalar,
	pkmn.RegionHisui,
	pkmn.RegionPaldea,
}

func FlipRegion() pkmn.Region {
	return regions[rand.Intn(len(regions))]
}
