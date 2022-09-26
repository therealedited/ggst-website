package main

import (
	"github.com/therealedited/ggst-website/database"
	"github.com/therealedited/ggst-website/internal"
)

func main() {
	database.UpdateDatabase(internal.Strive)
}
