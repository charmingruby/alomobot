package pkmn

import (
	"fmt"

	"github.com/charmingruby/alomobot/internal/text"
)

type Pokemon struct {
	Name          string `json:"name"`
	SpriteURL     string `json:"sprite_url"`
	PrimaryType   Type   `json:"primary_type"`
	SecondaryType *Type  `json:"secondary_type"`
	ID            int    `json:"id"`
}

func (p Pokemon) Render() string {
	typeContent := p.PrimaryType.Render()
	if p.SecondaryType != nil {
		typeContent += fmt.Sprintf(" / %s", p.SecondaryType.Render())
	}

	return fmt.Sprintf("**#%d — %s** • %s",
		p.ID,
		text.Capitalize(p.Name),
		typeContent,
	)
}
