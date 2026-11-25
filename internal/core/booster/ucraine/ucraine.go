package ucraine

import (
	"context"

	"github.com/charmingruby/alomobot/internal/core/booster"
)

type Ucraine struct {
	booster.SetDefinition
}

func New() *Ucraine {
	return &Ucraine{
		SetDefinition: booster.SetDefinition{
			Cards: []booster.Card{},
			Slots: []booster.Slot{
				{
					Types: []booster.CardType{
						booster.PokemonCardType,
					},
				},
			},
			RarityWeights: map[booster.Rarity]int{
				booster.CommonRarity:    25,
				booster.RareRarity:      25,
				booster.EpicRarity:      25,
				booster.LegendaryRarity: 25,
			},
		},
	}
}

func (s *Ucraine) Open(ctx context.Context, packs int) ([]booster.Card, error) {
	return nil, nil
}
