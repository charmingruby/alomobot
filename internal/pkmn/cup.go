package pkmn

type Cup string

const (
	CupAsh              Cup = "ash's"
	CupLittle           Cup = "little"
	CupDexLastNumber    Cup = "dex last number"
	CupCharactersAmount Cup = "characters amount"
	CupRental           Cup = "rental"
	CupGeneration       Cup = "generation"
	CupMonotype         Cup = "monotype"
	CupRandomParty      Cup = "random party"
	CupBooster          Cup = "booster"
	CupNoEvo            Cup = "no evo"
	CupSingleType       Cup = "single type"
)

func (c Cup) String() string {
	return string(c)
}
