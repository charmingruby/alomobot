package ucraine

import (
	"slices"

	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/set/ucraine/pokemon"
)

type Ucraine struct {
	booster.Set
}

func New() *Ucraine {
	return &Ucraine{
		Set: booster.Set{
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
				booster.CommonRarity:    35,
				booster.UncommonRarity:  25,
				booster.RareRarity:      20,
				booster.EpicRarity:      15,
				booster.LegendaryRarity: 5,
			},
		},
	}
}
