package template

import (
	"fmt"

	"github.com/charmingruby/alomobot/internal/core/cup"
)

func RulesCup(c cup.Cup) string {
	const pattern = `🏅 **%s Cup**

**How it works**
%s

**Restrictions**
%s

**Team building**
%s

**Global rules**
The global rules apply to every tournament unless its rules explicitly state otherwise. Read them with %s.`

	return fmt.Sprintf(
		pattern,
		c,
		c.Guide.HowItWorks,
		c.Guide.Restrictions,
		c.Guide.TeamBuilding,
		"`/rules global`",
	)
}
