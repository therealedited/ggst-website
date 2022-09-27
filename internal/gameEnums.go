package internal

type Game int

const (
	Strive Game = iota
	Xrd
	XX
	X
	MissingLink
	BlazblueCentralFiction
)

func (g Game) String() string {
	switch g {
	case 0:
		return "Strive"
	case 1:
		return "Xrd"
	case 2:
		return "XX"
	case 3:
		return "X"
	case 4:
		return "MissingLink"
	case 5:
		return "BlazblueCentralFiction"
	default:
		return "<nil>"
	}
}

type GGCharacter int

const (
	SO GGCharacter = iota
	KY
	MA
	AX
	CH
	PO
	FA
	MI
	ZA
	RA
	LE
	NA
	GI
	AN
	IN
	GO
	JC
	HA
	BA
	TE
	BR
)

func (gg GGCharacter) String() string {
	switch gg {
	case 0:
		return "SO"
	case 1:
		return "KY"
	case 2:
		return "MA"
	case 3:
		return "AX"
	case 4:
		return "CH"
	case 5:
		return "PO"
	case 6:
		return "FA"
	case 7:
		return "MI"
	case 8:
		return "ZA"
	case 9:
		return "RA"
	case 10:
		return "LE"
	case 11:
		return "NA"
	case 12:
		return "GI"
	case 13:
		return "AN"
	case 14:
		return "GO"
	case 15:
		return "JC"
	case 16:
		return "HA"
	case 17:
		return "BA"
	case 18:
		return "TE"
	case 19:
		return "BR"
	default:
		return "<nil>"
	}
}
