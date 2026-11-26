package support

import (
	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/deck"
)

var Legendary = []booster.Card{
	deck.MasterBall.WithRarity(booster.LegendaryRarity),
}
