package template

import (
	"fmt"

	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/format"
)

func OpenBooster(selected booster.Set, cards []booster.Card) string {
	return fmt.Sprintf(
		"📦 Opened a booster from **%s**!\n\n%s",
		selected,
		format.List(cards),
	)
}
