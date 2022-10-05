// This package provides functions to perform multiple operation on a database such as updating it with from a json that was fetched from dustloop.
package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/therealedited/ggst-website/internal"
	"gopkg.in/ini.v1"
)

var Inst *sql.DB

// Initializes the database singleton with different values coming from the global.ini file. As of now, it only works with MySQL.
func init() {
	dbCfg, err := ini.Load("./configs/global.ini")
	if err != nil {
		fmt.Printf("Failed to open database configuration: %v", err)
		os.Exit(-2)
	}

	username := dbCfg.Section("sql").Key("user").String()
	password := dbCfg.Section("sql").Key("password").String()
	dbname := dbCfg.Section("sql").Key("dbname").String()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", username, password, dbname))
	if err != nil {
		fmt.Printf("Failed to initialize database: %v", err)
		os.Exit(-1)
	}
	Inst = db
}

func PingDatabase() bool {
	err := Inst.Ping()
	return err != nil
}

// Updates the database from a json.
func UpdateDatabase(gameID internal.Game) {
	if gameID == internal.Strive {
		var striveMoves []internal.StriveMove
		dustloopData := internal.ReadLocalJsonData("database/json/strive.json") //Maybe add that to config.ini?
		internal.JsonMarshalling(&dustloopData, &striveMoves)
		//updateStriveDatabase(&striveMoves)
		fmt.Printf("%s", striveMoves[0].Name)
	}
}

// Specific function to update the Guilty Gear: Strive database.
func UpdateStriveDatabase(s *[]internal.StriveMove) {
	fmt.Print("Updating Strive.")
	var moveName string
	for _, v := range *s {
		if v.Name == "" {
			moveName = "Normal"
		} else {
			moveName = v.Name
		}
		_, err := Inst.Exec("INSERT INTO ggst.move (idChar_FK ,name, input, damage, guard, startup, active, recovery, onBlock, onHit, invuln, type) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
			internal.GGLongStringToInt(v.Chara),
			moveName,
			v.Input,
			v.Damage,
			v.Guard,
			v.Startup,
			v.Active,
			v.Recovery,
			v.OnBlock,
			v.OnHit,
			v.Invuln,
			v.Type)
		internal.CheckForError(err)
	}

}
