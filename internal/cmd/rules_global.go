package cmd

import (
	"context"
	"fmt"

	"github.com/charmingruby/erebor/discord/api"

	"github.com/charmingruby/alomobot/internal/cup"
	"github.com/charmingruby/alomobot/internal/format"
)

func rulesGlobal() api.Command {
	return api.Command{
		Name:        "global",
		Description: "Shows the global rules that apply to all tournaments",
		Handler: func(_ context.Context, _ api.Interaction) (api.Response, error) {

			pattern := `📜 **Global Rules**

_These rules apply to every tournament unless its rules explicitly state otherwise._

%s`

			return api.Response{Content: fmt.Sprintf(pattern, format.List(cup.GlobalRules))}, nil
		},
	}
}
