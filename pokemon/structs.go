package pokemon

type Pokemon struct {
	Abilities []Ability `json:"abilities"`
	Name      string    `json:"name"`
	Sprites   Sprite    `json:"sprites"`
	Height    int       `json:"height"`
	Weight    int       `json:"weight"`
	Types     []Type    `json:"types"`
	Cry       Cries     `json:"cries"`
}

type Cries struct {
	Latest string `json:"latest"`
}

type Ability struct {
	Ability  AbilityInfo `json:"ability"`
	IsHidden bool        `json:"is_hidden"`
	Slot     int         `json:"slot"`
}

type AbilityInfo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Sprite struct {
	Other Other `json:"other"`
}

type Other struct {
	Home Home `json:"home"`
}

type Home struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type Type struct {
	Slot int      `json:"slot"`
	Type TypeInfo `json:"type"`
}

type TypeInfo struct {
	Name string `json:"name"`
}