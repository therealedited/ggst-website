package internal

import "os"

func CheckForError(err error) {
	if err != nil {
		panic(err)
	}
}

func IsFileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func GGShortStringToInt(str string) int {
	switch str {
	case "SO":
		return 1
	case "KY":
		return 2
	case "MA":
		return 3
	case "AX":
		return 4
	case "CH":
		return 5
	case "PO":
		return 6
	case "FA":
		return 7
	case "MI":
		return 8
	case "ZA":
		return 9
	case "RA":
		return 10
	case "LE":
		return 11
	case "NA":
		return 12
	case "GI":
		return 13
	case "AN":
		return 14
	case "IN":
		return 15
	case "GO":
		return 16
	case "JC":
		return 17
	case "HA":
		return 18
	case "BA":
		return 19
	case "TE":
		return 20
	case "BR":
		return 21
	default:
		return -1
	}
}

func GGLongStringToInt(str string) int {
	switch str {
	case "Sol Badguy":
		return 1
	case "Ky Kiske":
		return 2
	case "May":
		return 3
	case "Axl Low":
		return 4
	case "Chipp Zanuff":
		return 5
	case "Potemkin":
		return 6
	case "Faust":
		return 7
	case "Millia Rage":
		return 8
	case "Zato-1":
		return 9
	case "Ramlethal Valentine":
		return 10
	case "Leo Whitefang":
		return 11
	case "Nagoriyuki":
		return 12
	case "Giovanna":
		return 13
	case "Anji Mito":
		return 14
	case "I-No":
		return 15
	case "Goldlewis Dickinson":
		return 16
	case "Jack-O":
		return 17
	case "Happy Chaos":
		return 18
	case "Baiken":
		return 19
	case "Testament":
		return 20
	case "Bridget":
		return 21
	default:
		return -1
	}
}
