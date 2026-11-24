package message

import (
	"fmt"

	"github.com/charmingruby/alomobot/internal/cup"
)

func HelpCup(cup cup.Cup) string {
	return fmt.Sprintf("🏅 **%s Cup**\n\n**How it works**\n%s\n\n**Restrictions**\n%s\n\n**Team building**\n%s",
		cup, cup.Guide.HowItWorks, cup.Guide.Restrictions, cup.Guide.TeamBuilding,
	)
}
