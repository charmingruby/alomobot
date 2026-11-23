package pkmn

type Cup string

const (
	CupAsh         Cup = "Ash's Cup"
	CupLittle      Cup = "Little Cup"
	CupLetter      Cup = "Letter Cup"
	CupRental      Cup = "Rental Cup"
	CupMonogen     Cup = "Monogen Cup"
	CupMonotype    Cup = "Monotype Cup"
	CupRandomParty Cup = "Random Party Cup"
	CupBooster     Cup = "Booster Cup"
	CupNoEvo       Cup = "No Evo Cup"
	CupSingleType  Cup = "Single Type Cup"
	CupStarter     Cup = "Starter Cup"
	CupChaos       Cup = "Chaos Cup"
	CupStandard    Cup = "Standard Cup"
)

var Cups = []Cup{
	CupAsh,
	CupLittle,
	CupLetter,
	CupRental,
	CupMonogen,
	CupMonotype,
	CupRandomParty,
	CupBooster,
	CupNoEvo,
	CupSingleType,
	CupStandard,
	CupStarter,
	CupChaos,
}

func (c Cup) String() string {
	return string(c)
}

func (c Cup) Render() string {
	return c.String()
}
