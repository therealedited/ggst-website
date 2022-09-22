package database

type StriveCharacter struct {
	CharacterName  string
	Input          string
	MoveName       string
	Damage         int
	Startup        string
	ActiveFrames   int
	RecoveryFrames int
	OnBlock        string
	OnHit          string
	Invuln         string
	Type           string
}
