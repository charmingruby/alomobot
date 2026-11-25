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
					Types: []booster.CardType{booster.PokemonCardType},
				},
				{
					Types: []booster.CardType{booster.PokemonCardType},
				},
				{
					Types: []booster.CardType{booster.PokemonCardType},
				},
				{
					Types:     []booster.CardType{booster.PokemonCardType},
					MinRarity: booster.RareRarity,
				},
				{
					Types: []booster.CardType{booster.EnergyCardType, booster.SupportCardType},
				},
				{
					Types: []booster.CardType{booster.EnergyCardType, booster.SupportCardType},
				},
			},
			RarityWeights: map[booster.Rarity]int{
				booster.CommonRarity:    30,
				booster.UncommonRarity:  25,
				booster.RareRarity:      20,
				booster.EpicRarity:      15,
				booster.LegendaryRarity: 10,
			},
		},
	}
}

func (s *Ucraine) Open(ctx context.Context, packs int) ([]booster.Card, error) {
	return nil, nil
}
