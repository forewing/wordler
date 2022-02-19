package wordler

import (
	_ "embed"
	"encoding/json"
	"fmt"
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

func GetWordList(l int) ([]string, error) {
	ret := SolutionList
	if l > 0 {
		if l >= len(WordList) {
			return ret, fmt.Errorf("word length exceeds, max: %v", len(WordList)-1)
		}
		ret = WordList[l]
	}

	return ret, nil
}
