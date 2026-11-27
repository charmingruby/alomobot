package energy

import (
	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/deck"
)

var Epic = []booster.Card{
	deck.NeutralEnergy.WithRarity(booster.EpicRarity),
}
