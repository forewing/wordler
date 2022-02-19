package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/forewing/wordler"
)

var (
	flagLength = flag.Int("len", 0, "length of word")
	flagMax    = flag.Int("max", 20, "max output words")

	filters wordler.FilterList
)

func addFilterContain(value string) error {
	for _, t := range strings.Split(value, ",") {
		filters = append(filters, wordler.FilterContain{
			Target: t,
		})
	}
	return nil
}

func addFilterNotContain(value string) error {
	for _, t := range strings.Split(value, ",") {
		filters = append(filters, wordler.FilterNotContain{
			Target: t,
		})
	}
	return nil
}

func addFilterAt(value string) error {
	for _, t := range strings.Split(value, ",") {
		index, letter, err := wordler.ParseAt(t)
		if err != nil {
			return err
		}
		filters = append(filters, wordler.FilterAt{
			Index:  index,
			Target: letter,
		})
	}
	return nil
}

func addFilterNotAt(value string) error {
	for _, t := range strings.Split(value, ",") {
		index, letter, err := wordler.ParseAt(t)
		if err != nil {
			return err
		}
		filters = append(filters, wordler.FilterNotAt{
			Index:  index,
			Target: letter,
		})
	}
	return nil
}

func init() {
	flag.Func("has", "comma-seperated `string`s the target contains", addFilterContain)
	flag.Func("no", "comma-seperated `string`s the target does not contain", addFilterNotContain)
	flag.Func("at", "comma-seperated `{index}:{letter}`, where {letter} is at {index} (0-based) of the target", addFilterAt)
	flag.Func("not-at", "comma-seperated `{index}:{letter}`, where {letter} is not at {index} (0-based) of the target", addFilterNotAt)
	flag.Parse()
}

func main() {
	if *flagLength >= len(wordler.WordList) {
		log.Fatalf("word length exceeds (%v > %v)", *flagLength, len(wordler.WordList)-1)
	}

	words := wordler.SolutionList
	if *flagLength > 0 {
		words = wordler.WordList[*flagLength]
	}

	result := filters.Run(words)

	if len(result) > *flagMax {
		result = result[0:*flagMax]
	}
	fmt.Println(result)
}
