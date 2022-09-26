package internal

func CheckForError(err error) {
	if err != nil {
		panic(err)
	}
}
