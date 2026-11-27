package support

import (
	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/deck"
)

var Common = []booster.Card{
	deck.PokeBall.WithRarity(booster.CommonRarity),
	deck.NestBall.WithRarity(booster.CommonRarity),
}
