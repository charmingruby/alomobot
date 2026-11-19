package pkmn

type Pokemon struct {
	Name      string `json:"name"`
	SpriteURL string `json:"sprite_url"`
	ID        int    `json:"id"`
}
