package template

import (
	"fmt"

	"github.com/charmingruby/alomobot/internal/core/cup"
	"github.com/charmingruby/alomobot/internal/format"
)

func RulesGlobal() string {
	const pattern = `📜 **Global Rules**

_These rules apply to every tournament unless its rules explicitly state otherwise._

%s`

	return fmt.Sprintf(pattern, format.List(cup.GlobalRules))
}
