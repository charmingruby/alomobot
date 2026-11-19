package data

type Pokemon struct {
	Name    string  `json:"name"`
	Sprites Sprites `json:"sprites"`
	ID      int     `json:"id"`
}

type Sprites struct {
	FrontDefault string `json:"front_default"`
}
