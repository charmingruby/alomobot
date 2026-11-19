package pkmn

type Region string

const (
	RegionKanto  Region = "Kanto"
	RegionJohto  Region = "Johto"
	RegionHoenn  Region = "Hoenn"
	RegionSinnoh Region = "Sinnoh"
	RegionUnova  Region = "Unova"
	RegionKalos  Region = "Kalos"
	RegionAlola  Region = "Alola"
	RegionGalar  Region = "Galar"
	RegionHisui  Region = "Hisui"
	RegionPaldea Region = "Paldea"
)

func (r Region) Valid() bool {
	switch r {
	case RegionKanto, RegionJohto, RegionHoenn, RegionSinnoh, RegionUnova,
		RegionKalos, RegionAlola, RegionGalar, RegionHisui, RegionPaldea:
		return true
	default:
		return false
	}
}

func (r Region) String() string {
	return string(r)
}

type Gen struct {
	Number     int
	Region     Region
	FirstDexID int
	LastDexID  int
}

var (
	GenI = Gen{
		Number:     1,
		Region:     RegionKanto,
		FirstDexID: 1,
		LastDexID:  151,
	}
	GenII = Gen{
		Number:     2,
		Region:     RegionJohto,
		FirstDexID: 152,
		LastDexID:  251,
	}
	GenIII = Gen{
		Number:     3,
		Region:     RegionHoenn,
		FirstDexID: 252,
		LastDexID:  386,
	}
	GenIV = Gen{
		Number:     4,
		Region:     RegionSinnoh,
		FirstDexID: 387,
		LastDexID:  493,
	}
	GenV = Gen{
		Number:     5,
		Region:     RegionUnova,
		FirstDexID: 494,
		LastDexID:  649,
	}
	GenVI = Gen{
		Number:     6,
		Region:     RegionKalos,
		FirstDexID: 650,
		LastDexID:  721,
	}
	GenVII = Gen{
		Number:     7,
		Region:     RegionAlola,
		FirstDexID: 722,
		LastDexID:  809,
	}
	GenVIII = Gen{
		Number:     8,
		Region:     RegionGalar,
		FirstDexID: 810,
		LastDexID:  905,
	}
	GenIX = Gen{
		Number:     9,
		Region:     RegionPaldea,
		FirstDexID: 906,
		LastDexID:  1025,
	}
)

var Gens = []Gen{
	GenI,
	GenII,
	GenIII,
	GenIV,
	GenV,
	GenVI,
	GenVII,
	GenVIII,
	GenIX,
}

func (g Gen) Valid() bool {
	return g.Number >= 1 && g.Number <= 9
}
