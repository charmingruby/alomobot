package booster

type Rarity string

const (
	CommonRarity    Rarity = "common"
	UncommonRarity  Rarity = "uncommon"
	RareRarity      Rarity = "rare"
	EpicRarity      Rarity = "epic"
	LegendaryRarity Rarity = "legendary"
)

var rarities = []Rarity{
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
