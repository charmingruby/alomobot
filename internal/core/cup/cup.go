package cup

type GuideCup struct {
	HowItWorks   string
	Restrictions string
	TeamBuilding string
}

type CupName string

type Cup struct {
	Name  CupName
	Guide GuideCup
}

var Cups = map[CupName]Cup{
	LittleCup:      littleCup,
	LetterCup:      letterCup,
	RentalCup:      rentalCup,
	MonogenCup:     monogenCup,
	MonotypeCup:    monotypeCup,
	RandomPartyCup: randomPartyCup,
	BoosterCup:     boosterCup,
	NoEvoCup:       noEvoCup,
	SingleTypeCup:  singleTypeCup,
	StandardCup:    standardCup,
	StarterCup:     starterCup,
	ShinyCup:       shinyCup,
	FavsCup:        favsCup,
}

func (c Cup) String() string {
	return string(c.Name)
}

func (c Cup) Render() string {
	return c.String()
}

func (c Cup) Valid() bool {
	_, ok := FindCup(c.String())
	return ok
}

func FindCup(name string) (Cup, bool) {
	selected, ok := Cups[CupName(name)]
	return selected, ok
}
