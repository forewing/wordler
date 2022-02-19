package wordler

import (
	_ "embed"
	"encoding/json"
)

var (
	//go:embed resources/words.json
	words []byte

	//go:embed resources/solutions.json
	solutions []byte

	WordList     [][]string
	SolutionList []string
)

func init() {
	err := json.Unmarshal(words, &WordList)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(solutions, &SolutionList)
	if err != nil {
		panic(err)
	}
}
