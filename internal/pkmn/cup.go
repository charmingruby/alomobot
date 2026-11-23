package pkmn

type Cup string

const (
	CupAsh              Cup = "Ash's Cup"
	CupLittle           Cup = "Little Cup"
	CupDexLastNumber    Cup = "Dex Last Number Cup"
	CupCharactersAmount Cup = "Characters Amount Cup"
	CupRental           Cup = "Rental Cup"
	CupMonogen          Cup = "Monogen Cup"
	CupMonotype         Cup = "Monotype Cup"
	CupRandomParty      Cup = "Random Party Cup"
	CupBooster          Cup = "Booster Cup"
	CupNoEvo            Cup = "No Evo Cup"
	CupSingleType       Cup = "Single Type Cup"
	CupStandard         Cup = "Standard Cup"
)

var Cups = []Cup{
	CupAsh,
	CupLittle,
	CupDexLastNumber,
	CupCharactersAmount,
	CupRental,
	CupMonogen,
	CupMonotype,
	CupRandomParty,
	CupBooster,
	CupNoEvo,
	CupSingleType,
	CupStandard,
}

func (c Cup) String() string {
	return string(c)
}

func (c Cup) Render() string {
	return c.String()
}
