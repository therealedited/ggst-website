package database

type StriveMove struct {
	Chara    string      `json:"chara"`
	Input    interface{} `json:"input,omitempty"`
	Name     string      `json:"name,omitempty"`
	Damage   interface{} `json:"damage,omitempty"`
	Guard    interface{} `json:"guard,omitempty"`
	Startup  interface{} `json:"startup,omitempty"`
	Active   interface{} `json:"active,omitempty"`
	Recovery interface{} `json:"recovery,omitempty"`
	OnBlock  interface{} `json:"onBlock,omitempty"`
	OnHit    interface{} `json:"onHit,omitempty"`
	Invuln   interface{} `json:"invuln,omitempty"`
	Type     string      `json:"type,omitempty"`
}

type XrdMove struct {
	CharacterName string
	Input         string
	Name          string
	Damage        int
	Guard         string
	Startup       string
	Active        string
	Recovery      string
	Level         int
	OnBlock       string
	OnHit         string
	Invuln        string
}

type GameMove interface {
	StriveMove | XrdMove
}
