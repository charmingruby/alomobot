package template

import (
	"fmt"

	"github.com/charmingruby/alomobot/internal/core/booster"
	"github.com/charmingruby/alomobot/internal/format"
)

func SearchCards(parameters string, cards []booster.Card, total, page, pages int) string {
	const pageSize = 25
	start := 0
	end := 0
	if len(cards) > 0 {
		start = (page-1)*pageSize + 1
		end = start + len(cards) - 1
	}

	const pattern = `🔎 **Search results**

**Parameters:** %s

**Cards (%d)**
%s

Showing **%d–%d of %d cards** • Page **%d / %d**`

	return fmt.Sprintf(
		pattern,
		parameters,
		total,
		format.List(cards),
		start,
		end,
		total,
		page,
		pages,
	)
}
