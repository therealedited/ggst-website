// This files hold all the different enums used to work with the database and some functions associated to it such as Stringers.
package internal

//Type for game ids.
type Game int

const (
	Strive Game = iota
	Xrd
	XX
	X
	MissingLink
	BlazblueCentralFiction
)

// Stringer for game constants. Transforms an interger into its string value. As an example, Strive.String() returns 0.
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

//Type for guilty gear characters ids.
type GGCharacter int

const (
	SO GGCharacter = iota + 1
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

//Stringer for game characters. Transforms an interger into its string value. As an example, SO.String() returns 1.
func (gg GGCharacter) String() string {
	switch gg {
	case 1:
		return "SO"
	case 2:
		return "KY"
	case 3:
		return "MA"
	case 4:
		return "AX"
	case 5:
		return "CH"
	case 6:
		return "PO"
	case 7:
		return "FA"
	case 8:
		return "MI"
	case 9:
		return "ZA"
	case 10:
		return "RA"
	case 11:
		return "LE"
	case 12:
		return "NA"
	case 13:
		return "GI"
	case 14:
		return "AN"
	case 15:
		return "IN"
	case 16:
		return "GO"
	case 17:
		return "JC"
	case 18:
		return "HC"
	case 19:
		return "BA"
	case 20:
		return "TE"
	case 21:
		return "BR"
	default:
		return "<nil>"
	}
}
