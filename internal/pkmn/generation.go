package pkmn

type Generation string

const (
	GenerationI    Generation = "I"
	GenerationII   Generation = "II"
	GenerationIII  Generation = "III"
	GenerationIV   Generation = "IV"
	GenerationV    Generation = "V"
	GenerationVI   Generation = "VI"
	GenerationVII  Generation = "VII"
	GenerationVIII Generation = "VIII"
	GenerationIX   Generation = "IX"
)

func (g Generation) Valid() bool {
	switch g {
	case GenerationI, GenerationII, GenerationIII, GenerationIV, GenerationV,
		GenerationVI, GenerationVII, GenerationVIII, GenerationIX:
		return true
	default:
		return false
	}
}

func (g Generation) String() string {
	return string(g)
}
