package support

import (
	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/core/booster/deck"
)

var Epic = []booster.Card{
	deck.LureBall.WithRarity(booster.EpicRarity),
	deck.TimerBall.WithRarity(booster.EpicRarity),
	deck.LuxuryBall.WithRarity(booster.EpicRarity),
	deck.ProfessorSycamore.WithRarity(booster.EpicRarity),
	deck.PokemonBreeder.WithRarity(booster.EpicRarity),
}
