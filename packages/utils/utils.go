package utils

import "strconv"

func RuneToInt(c rune) int {
	d, err := strconv.Atoi(string(c))

	if err != nil {
		return 0
	}

	return d
}

func ConcatRunes(runes ...rune) string {
	return string(runes)
}
