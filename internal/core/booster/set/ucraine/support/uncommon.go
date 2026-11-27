package support

import (
	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/deck"
)

var Uncommon = []booster.Card{
	deck.GreatBall.WithRarity(booster.UncommonRarity),
	deck.NetBall.WithRarity(booster.UncommonRarity),
	deck.DuskBall.WithRarity(booster.UncommonRarity),
	deck.RepeatBall.WithRarity(booster.UncommonRarity),
	deck.QuickBall.WithRarity(booster.UncommonRarity),
	deck.ProfessorOak.WithRarity(booster.UncommonRarity),
	deck.Bill.WithRarity(booster.UncommonRarity),
}
