package encoder

import (
	"strconv"
	"strings"
	"unicode"

	"utils"
)

var charBindings = map[rune][2]int{
	'a': {2, 1},
	'b': {2, 2},
	'c': {2, 3},
	'd': {3, 1},
	'e': {3, 2},
	'f': {3, 3},
	'g': {4, 1},
	'h': {4, 2},
	'i': {4, 3},
	'j': {5, 1},
	'k': {5, 2},
	'l': {5, 3},
	'm': {6, 1},
	'n': {6, 2},
	'o': {6, 3},
	'p': {7, 1},
	'q': {7, 2},
	'r': {7, 3},
	's': {7, 4},
	't': {8, 1},
	'u': {8, 2},
	'v': {8, 3},
	'w': {9, 1},
	'x': {9, 2},
	'y': {9, 3},
	'z': {9, 4},
	' ': {0, 1},
}

type charEncoder struct {
	char    rune
	binding int
	repeat  int
	isUpper bool
}

func (c charEncoder) encode() string {
	repeated := strings.Repeat(strconv.Itoa(c.binding), c.repeat)

	if c.isUpper {
		repeated = string(utils.UPPER_CASE_REPRESENTOR) + repeated
	}

	return repeated
}

func (c charEncoder) shouldSeparate(builder []string) bool {
	if len(builder) == 0 {
		return false
	}

	lastBinding := builder[len(builder)-1]
	lastRune := lastBinding[len(lastBinding)-1]

	return unicode.IsDigit(rune(lastRune)) && c.binding == utils.RuneToInt(rune(lastRune))
}

type to3310Encoder struct {
	input   string
	builder []string
}

func (e *to3310Encoder) build() string {
	return strings.Join(e.builder, "")
}

func (e *to3310Encoder) append(s string) {
	e.builder = append(e.builder, s)
}

func (e *to3310Encoder) encode() {
	for _, letter := range e.input {

		charBinding := charBindings[unicode.ToLower(letter)]
		binding := charBinding[0]
		repeat := charBinding[1]

		if binding == 0 && repeat == 0 {

			if letter == utils.UPPER_CASE_REPRESENTOR {
				e.append(utils.ConcatRunes(letter, letter))
				continue
			}

			if isNum := unicode.IsDigit(letter); isNum {
				e.append(utils.ConcatRunes(utils.NUMBER_SPLITTER, letter))
				continue
			}

			e.append(string(letter))
			continue
		}

		isUpper := unicode.IsUpper(letter)

		c := charEncoder{
			char:    letter,
			binding: binding,
			repeat:  repeat,
			isUpper: isUpper,
		}

		encodedChar := c.encode()

		if shouldSeparate := c.shouldSeparate(e.builder); shouldSeparate {
			e.append(string(utils.SAME_CHAR_SEPARATOR))
		}

		e.append(encodedChar)
	}
}

// Encode takes a string in English and returns its 3310 representation
func Encode(sentence string) string {
	encoder := to3310Encoder{
		input: sentence,
	}
	encoder.encode()

	return encoder.build()
}
