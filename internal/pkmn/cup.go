package pkmn

type Cup string

const (
	CupAsh              Cup = "ash's"
	CupLittle           Cup = "little"
	CupDexLastNumber    Cup = "dex last number"
	CupCharactersAmount Cup = "characters amount"
	CupRental           Cup = "rental"
	CupGen              Cup = "gen"
	CupMonotype         Cup = "monotype"
	CupRandomParty      Cup = "random party"
	CupBooster          Cup = "booster"
	CupNoEvo            Cup = "no evo"
	CupSingleType       Cup = "single type"
	CupStandard         Cup = "standard"
)

var Cups = []Cup{
	CupAsh,
	CupLittle,
	CupDexLastNumber,
	CupCharactersAmount,
	CupRental,
	CupGen,
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
