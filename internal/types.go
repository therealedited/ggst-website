// This package holds all the different structures used within the codebase.
package internal

//Type for a move from Guilty Gear: Strive
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

//Type for a move from Guilty Gear: Xrd. Placeholder for now.
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

//Type abstraction for a fighting game move.
type GameMove interface {
	StriveMove | XrdMove
}
