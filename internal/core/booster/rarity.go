package booster

type Rarity string

const (
	CommonRarity    Rarity = "Common"
	UncommonRarity  Rarity = "Uncommon"
	RareRarity      Rarity = "Rare"
	EpicRarity      Rarity = "Epic"
	LegendaryRarity Rarity = "Legendary"
)

var Rarities = []Rarity{
	CommonRarity,
	UncommonRarity,
	RareRarity,
	EpicRarity,
	LegendaryRarity,
}

func (r Rarity) Emoji() string {
	switch r {
	case CommonRarity:
		return "⚪️"
	case UncommonRarity:
		return "🟢"
	case RareRarity:
		return "🔵"
	case EpicRarity:
		return "🟣"
	case LegendaryRarity:
		return "🟠"
	default:
		return "⚫️"
	}
}
