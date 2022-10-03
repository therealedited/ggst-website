package tests

import (
	"testing"

	"github.com/therealedited/ggst-website/internal"
)

func TestIsFileExists(t *testing.T) {
	got := internal.IsFileExists("../configs/global.ini")
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestGGShortStringToInt(t *testing.T) {
	got := internal.GGShortStringToInt("SO")
	want := 1

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestGGLongStringToInt(t *testing.T) {
	got := internal.GGLongStringToInt("Sol Badguy")
	want := 1

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
