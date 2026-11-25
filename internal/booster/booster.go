package booster

import "context"

type CardType string

const (
	EnergyCardType  CardType = "energy"
	PokemonCardType CardType = "pokemon"
	TrainerCardType CardType = "trainer"
)

type Rarity string

const (
	CommonRarity    Rarity = "common"
	RareRarity      Rarity = "rare"
	EpicRarity      Rarity = "epic"
	LegendaryRarity Rarity = "legendary"
)

type Card struct {
	Name        string
	Description string
	Set         string
	Type        CardType
	Rarity      Rarity
}

type Slot struct {
	Types []CardType
}

type SetDefinition struct {
	Cards         []Card
	Slots         []Slot
	RarityWeights map[Rarity]int
}

type Set interface {
	Open(ctx context.Context, packs int) ([]Card, error)
}
