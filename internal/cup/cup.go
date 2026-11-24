package cup

import "slices"

type GuideCup struct {
	HowItWorks   string
	Restrictions string
	TeamBuilding string
}

type NameCup string

type Cup struct {
	Name  NameCup
	Guide GuideCup
}

var Cups = []Cup{
	ashCup,
	littleCup,
	letterCup,
	rentalCup,
	monogenCup,
	monotypeCup,
	randomPartyCup,
	boosterCup,
	noEvoCup,
	singleTypeCup,
	standardCup,
	starterCup,
	chaosCup,
}

func (c Cup) String() string {
	return string(c.Name)
}

func (c Cup) Render() string {
	return c.String()
}

func (c Cup) Valid() bool {
	switch c.Name {
	case AshCup, LittleCup, LetterCup, RentalCup, MonogenCup, MonotypeCup,
		RandomPartyCup, BoosterCup, NoEvoCup, SingleTypeCup, StandardCup,
		StarterCup, ChaosCup:
		return true
	default:
		return false
	}
}

func FindCup(name string) (Cup, bool) {
	index := slices.IndexFunc(Cups, func(c Cup) bool {
		return c.Name == NameCup(name)
	})

	if index < 0 {
		return Cup{}, false
	}

	return Cups[index], true
}
