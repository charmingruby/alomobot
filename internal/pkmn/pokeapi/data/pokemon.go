package data

type Pokemon struct {
	Name    string  `json:"name"`
	Sprites Sprites `json:"sprites"`
	Types   []Type  `json:"types"`
	ID      int     `json:"id"`
}

type Sprites struct {
	Other struct {
		OfficialArtwork struct {
			FrontDefault string `json:"front_default"`
		} `json:"official-artwork"`
	} `json:"other"`
}

type Type struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}
