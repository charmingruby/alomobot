package ucraine

import (
	"context"
	"slices"

	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/set/ucraine/pokemon"
)

type Ucraine struct {
	booster.SetDefinition
}

func New() *Ucraine {
	return &Ucraine{
		SetDefinition: booster.SetDefinition{
			PokemonCards: slices.Concat(pokemon.Common, pokemon.Uncommon, pokemon.Rare, pokemon.Epic, pokemon.Legendary),
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
				// {
				// 	Types: []booster.CardType{booster.EnergyCardType, booster.SupportCardType},
				// },
				// {
				// 	Types: []booster.CardType{booster.EnergyCardType, booster.SupportCardType},
				// },
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

func (s *Ucraine) Open(ctx context.Context) ([]booster.Card, error) {
	for _, c := range s.PokemonCards {

	}

	return nil, nil
}
