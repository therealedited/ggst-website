package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/therealedited/ggst-website/database"
	"github.com/therealedited/ggst-website/internal"
)

func getMoves(c *gin.Context) { // v1/:short
	shortname := c.Param("short")
	database.Inst.Query("SELECT move.name, move.input FROM move WHERE idMove IN (SELECT character_move.idMove FROM character_move WHERE idCharacter=" + strconv.Itoa(internal.GGShortStringToInt(shortname)) + ")")
}
