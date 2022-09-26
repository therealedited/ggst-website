package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/therealedited/ggst-website/internal"
	"gopkg.in/ini.v1"
)

var movelistStrive string = "https://www.dustloop.com/wiki/index.php?title=Special:CargoExport&tables=MoveData_GGST&&fields=chara%2C+input%2C+name%2C+damage%2C+guard%2C+startup%2C+active%2C+recovery%2C+onBlock%2C+onHit%2C+invuln%2C+type&&order+by=%60chara%60%2C%60input%60%2C%60name%60%2C%60cargo__MoveData_GGST%60.%60images__full%60%2C%60damage%60&limit=1000&format=json"

// Thanks https://github.com/movildima/GGStriveUtilsBot/blob/master/GGStriveUtilsBot/Utils/DustloopDataFetcher.cs#L17
var Inst *sql.DB

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

func UpdateDatabase(gameID internal.Game) {
	if gameID == internal.Strive {
		var striveMoves []StriveMove
		dustloopData := readLocalJsonData("database/json/strive.json", movelistStrive)
		jsonMarshalling(&dustloopData, &striveMoves)
		fmt.Printf("%v", striveMoves[2].Input)
	}
}
func initHTTPClient() http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	return *client
}

func readLocalJsonData(path string, url string) []byte {
	if !internal.IsFileExists(path) {
		err := os.WriteFile(path, fetchData(url), 0666)
		internal.CheckForError(err)
	}
	dustloopData, err := os.ReadFile(path)
	internal.CheckForError(err)
	return dustloopData
}

func fetchData(url string) []byte {
	client := initHTTPClient()
	resp, err := client.Get(url)
	internal.CheckForError(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	internal.CheckForError(err)
	return body
}

func jsonMarshalling[K GameMove](data *[]byte, moveData *[]K) {
	err := json.Unmarshal(*data, &moveData)
	internal.CheckForError(err)
}
