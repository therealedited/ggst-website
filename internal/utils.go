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
