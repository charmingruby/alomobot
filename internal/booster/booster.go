package booster

import "context"

type CardType string

const (
	CardTypeEnergy  CardType = "energy"
	CardTypePokemon CardType = "pokemon"
	CardTypeTrainer CardType = "trainer"
)

type CardRarity string

const (
	CardRarityCommon    CardRarity = "common"
	CardRarityRare      CardRarity = "rare"
	CardRarityEpic      CardRarity = "epic"
	CardRarityLegendary CardRarity = "legendary"
)

type Card struct {
	Name        string
	Description string
	Set         *string
	Type        CardType
	Rarity      CardRarity
}

type BoosterPack interface {
	Open(context.Context) []Card
}
