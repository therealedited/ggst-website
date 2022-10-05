package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Thanks https://github.com/movildima/GGStriveUtilsBot/blob/master/GGStriveUtilsBot/Utils/DustloopDataFetcher.cs#L17
var movelistStrive string = "https://www.dustloop.com/wiki/index.php?title=Special:CargoExport&tables=MoveData_GGST&&fields=chara%2C+input%2C+name%2C+damage%2C+guard%2C+startup%2C+active%2C+recovery%2C+onBlock%2C+onHit%2C+invuln%2C+type&&order+by=%60chara%60%2C%60input%60%2C%60name%60%2C%60cargo__MoveData_GGST%60.%60images__full%60%2C%60damage%60&limit=1000&format=json"

// Checks for errors. Terminates the program if there is one.
func CheckForError(err error) {
	if err != nil {
		panic(err)
	}
}

// Checks if a file exists at the given path.
func IsFileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Function to transform a "ShortString" to its GGCharacter id. See gameEnums.go
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

// Function to transform a "LongString" to its GGCharacter id. See gameEnums.go
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

// Initializes a http client to communicate over the internet.
func InitHTTPClient() http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	return *client
}

// Fetches data by performing a GET request on an URL.
func FetchData(url string) []byte {
	client := InitHTTPClient()
	resp, err := client.Get(url)
	CheckForError(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	CheckForError(err)
	return body
}

// Transforms a byte array into a moveData struct array, an abstract type that should hold move data from a fighting game.
func JsonMarshalling[K GameMove](data *[]byte, moveData *[]K) {
	err := json.Unmarshal(*data, &moveData)
	CheckForError(err)
}

// Reads a json from disk and returns its content as a byte array.
func ReadLocalJsonData(path string) []byte {

	if !IsFileExists(path) {
		if strings.Contains(path, "strive") {
			FetchAndCreateFile(path, movelistStrive)
		}
	}
	dustloopData, err := os.ReadFile(path)
	CheckForError(err)
	return dustloopData
}

// Used when a file doesn't exist and needs to be fetched from a remote location. It is then written on the disk as to not make too much requests on an api.
func FetchAndCreateFile(url string, path string) {
	err := os.WriteFile(path, FetchData(url), 0666)
	CheckForError(err)
}
