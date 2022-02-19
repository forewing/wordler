package wordler

import (
	"errors"
	"strconv"
	"strings"
)

func ParseAt(value string) (int, byte, error) {
	l := strings.Split(value, ":")
	if len(l) != 2 || len(l[1]) != 1 {
		return 0, 0, errors.New("at should be like `2:c` ({index}:{letter})")
	}
	index, err := strconv.Atoi(l[0])
	if err != nil {
		return 0, 0, err
	}
	return index, l[1][0], nil
}
