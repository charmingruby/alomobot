package data

type Pokemon struct {
	Name    string  `json:"name"`
	Sprites Sprites `json:"sprites"`
	Types   []Type  `json:"types"`
	ID      int     `json:"id"`
}

type Sprites struct {
	FrontDefault string `json:"front_default"`
}

type Type struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}
