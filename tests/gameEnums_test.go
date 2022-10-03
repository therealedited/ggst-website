package tests

import (
	"testing"

	"github.com/therealedited/ggst-website/internal"
)

func TestStringerGameName(t *testing.T) {
	got := internal.Strive.String()
	want := "Strive"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

}

func TestStringerCharacterName(t *testing.T) {
	got := internal.SO.String()
	want := "SO"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

}
