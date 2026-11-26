package support

import (
	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/deck"
)

var Epic = []booster.Card{
	deck.Fisherman.WithRarity(booster.EpicRarity),
}
