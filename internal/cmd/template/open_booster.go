package template

import (
	"fmt"
	"strings"

	"github.com/charmingruby/alomobot/internal/core/booster"
)

func OpenBooster(selected booster.Set, cards []booster.Card) string {
	var content strings.Builder
	fmt.Fprintf(&content, "📦 **%s Booster**\n\n", selected)
	for i, card := range cards {
		fmt.Fprintf(&content, "%d. **%s** — %s\n", i+1, card.Rarity.Emoji(), card.Name)
	}

	return content.String()
}
