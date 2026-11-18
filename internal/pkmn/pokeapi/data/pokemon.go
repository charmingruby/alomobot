package data

type Pokemon struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Sprites Sprites `json:"sprites"`
}

type Sprites struct {
	FrontDefault string `json:"front_default"`
}
