package main

import (
	"strconv"
	"strings"
	"unicode"
)

const (
	NUMBER_SPLITTER        = ' '
	UPPER_CASE_REPRESENTOR = '_'
	SAME_CHAR_SEPARATOR    = '1'
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

var keyBindings = map[int][]rune{
	0: {' '},
	1: {},
	2: {'a', 'b', 'c'},
	3: {'d', 'e', 'f'},
	4: {'g', 'h', 'i'},
	5: {'j', 'k', 'l'},
	6: {'m', 'n', 'o'},
	7: {'p', 'q', 'r', 's'},
	8: {'t', 'u', 'v'},
	9: {'w', 'x', 'y', 'z'},
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
		repeated = string(UPPER_CASE_REPRESENTOR) + repeated
	}

	return repeated
}

func runeToInt(c rune) int {
	d, err := strconv.Atoi(string(c))

	if err != nil {
		return 0
	}

	return d
}

func (c charEncoder) shouldSeparate(builder []string) bool {
	if len(builder) == 0 {
		return false
	}

	lastBinding := builder[len(builder)-1]
	lastRune := lastBinding[len(lastBinding)-1]

	return c.binding == runeToInt(rune(lastRune))
}

func concatRunes(runes ...rune) string {
	return string(runes)
}

func countChar(str string, c string, i, count int) (repeat, index int) {
	if len(str) == i {
		return count, i
	}

	if c == string(str[i]) {
		return countChar(str, c, i+1, count+1)
	}

	return count, i
}

type from3310Decoder struct {
	input            string
	builder          []string
	isNextUpper      bool
	isNextDigit      bool
	nextIndexToCheck int
}

func (d *from3310Decoder) build() string {
	return strings.Join(d.builder, "")
}

func (d *from3310Decoder) append(s string) {
	d.builder = append(d.builder, s)
}

func (d *from3310Decoder) skip(to ...int) {
	if len(to) == 1 {
		d.nextIndexToCheck = to[0]
		return
	}

	d.nextIndexToCheck++
}

func (d *from3310Decoder) markNextUpper(isUpper bool) {
	d.isNextUpper = isUpper
}

func (d *from3310Decoder) markNextDigit(isDigit bool) {
	d.isNextDigit = isDigit
}

func (d *from3310Decoder) decode() {
	for i, l := range d.input {
		letter := string(l)

		if i < d.nextIndexToCheck {
			continue
		}

		if unicode.IsDigit(l) {
			if d.isNextDigit {
				d.append(letter)
				d.skip()
				d.markNextDigit(false)
				continue
			}

			charRep := keyBindings[runeToInt(l)]

			if len(charRep) > 0 {
				repeat, next := countChar(d.input, letter, d.nextIndexToCheck, 0)

				letter := string(charRep[repeat-1])

				if d.isNextUpper {
					letter = strings.ToUpper(letter)
				}

				d.append(letter)
				d.markNextUpper(false)
				d.skip(next)
				continue
			}

			d.skip()
			continue
		}

		if l == NUMBER_SPLITTER {
			d.markNextDigit(true)
			d.skip()
			continue
		}

		if l == UPPER_CASE_REPRESENTOR {
			repeat, next := countChar(d.input, letter, d.nextIndexToCheck, 0)

			if repeat%2 != 0 {
				repeat -= 1
				d.markNextUpper(true)
			}

			d.append(strings.Repeat(string(UPPER_CASE_REPRESENTOR), repeat/2))
			d.skip(next)
			continue
		}

		d.append(letter)
		d.skip()
	}
}

// Decode takes a string in 3310 representation and returns its English value
func Decode(sentence string) string {
	decoder := from3310Decoder{
		input: sentence,
	}
	decoder.decode()
	return decoder.build()
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

			if letter == UPPER_CASE_REPRESENTOR {
				e.append(concatRunes(letter, letter))
				continue
			}

			if isNum := unicode.IsDigit(letter); isNum {
				e.append(concatRunes(NUMBER_SPLITTER, letter))
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
			e.append(string(SAME_CHAR_SEPARATOR))
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
