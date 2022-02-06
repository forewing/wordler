package wordler

import (
	_ "embed"
	"encoding/json"
)

var (
	//go:embed resources/words.json
	data []byte

	WordList [][]string
)

func init() {
	err := json.Unmarshal(data, &WordList)
	if err != nil {
		panic(err)
	}
}
