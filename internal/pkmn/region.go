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
