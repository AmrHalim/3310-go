package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	NUMBER_SPLITTER        = ' '
	UPPER_CASE_REPRESENTOR = '_'
	SAME_CHAR_SEPARATOR    = '1'
)

type charRepresentation [2]int

var charBindings = map[string]charRepresentation{
	"a": {2, 1},
	"b": {2, 2},
	"c": {2, 3},
	"d": {3, 1},
	"e": {3, 2},
	"f": {3, 3},
	"g": {4, 1},
	"h": {4, 2},
	"i": {4, 3},
	"j": {5, 1},
	"k": {5, 2},
	"l": {5, 3},
	"m": {6, 1},
	"n": {6, 2},
	"o": {6, 3},
	"p": {7, 1},
	"q": {7, 2},
	"r": {7, 3},
	"s": {7, 4},
	"t": {8, 1},
	"u": {8, 2},
	"v": {8, 3},
	"w": {9, 1},
	"x": {9, 2},
	"y": {9, 3},
	"z": {9, 4},
	" ": {0, 1},
}

var keyBindings = map[int][]string{
	0: {" "},
	1: {},
	2: {"a", "b", "c"},
	3: {"d", "e", "f"},
	4: {"g", "h", "i"},
	5: {"j", "k", "l"},
	6: {"m", "n", "o"},
	7: {"p", "q", "r", "s"},
	8: {"t", "u", "v"},
	9: {"w", "x", "y", "z"},
}

type encoder struct {
	char      string
	binding   int
	repeat    int
	isCapital bool
}

func (c encoder) encode() string {
	repeated := strings.Repeat(strconv.Itoa(c.binding), c.repeat)

	if c.isCapital {
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

func (c encoder) shouldSeparate(builder []string) bool {
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

func decode(sentence string) string {
	stringBuilder := []string{}
	skipToIndex := 0
	isNextUpper := false
	isNextDigit := false

	for i, l := range sentence {
		letter := string(l)

		if i < skipToIndex {
			continue
		}

		if unicode.IsDigit(l) {
			if isNextDigit {
				stringBuilder = append(stringBuilder, letter)

				skipToIndex++
				isNextDigit = false
				continue
			}

			charRep := keyBindings[runeToInt(l)]

			if len(charRep) > 0 {
				repeat, next := countChar(sentence, letter, skipToIndex, 0)
				skipToIndex = next
				letter := charRep[repeat-1]

				if isNextUpper {
					letter = strings.ToUpper(letter)
				}

				stringBuilder = append(stringBuilder, letter)

				isNextUpper = false
				continue
			}

			skipToIndex++
			continue
		}

		if l == NUMBER_SPLITTER {
			isNextDigit = true

			skipToIndex++
			continue
		}

		if l == UPPER_CASE_REPRESENTOR {
			repeat, next := countChar(sentence, letter, skipToIndex, 0)
			skipToIndex = next

			if repeat%2 != 0 {
				repeat -= 1
				isNextUpper = true
			}

			stringBuilder = append(
				stringBuilder,
				strings.Repeat(string(UPPER_CASE_REPRESENTOR), repeat/2),
			)

			continue
		}

		stringBuilder = append(stringBuilder, letter)
		skipToIndex++
	}

	return strings.Join(stringBuilder, "")
}

// encode takes a string in English and returns its 3310 representation
func encode(sentence string) string {
	stringBuilder := []string{}

	for _, letter := range sentence {
		charBinding := charBindings[strings.ToLower(string(letter))]
		binding := charBinding[0]
		repeat := charBinding[1]

		if binding == 0 && repeat == 0 {

			if letter == UPPER_CASE_REPRESENTOR {
				stringBuilder = append(stringBuilder, concatRunes(letter, letter))
			} else {
				if isNum := unicode.IsDigit(letter); isNum {

					stringBuilder = append(stringBuilder, concatRunes(NUMBER_SPLITTER, letter))
				} else {

					stringBuilder = append(stringBuilder, string(letter))
				}
			}

		} else {
			isCap := unicode.IsUpper(letter)

			c := encoder{
				char:      string(letter),
				binding:   binding,
				repeat:    repeat,
				isCapital: isCap,
			}

			encodedChar := c.encode()

			if shouldSeparate := c.shouldSeparate(stringBuilder); shouldSeparate {
				encodedChar = string(SAME_CHAR_SEPARATOR) + encodedChar
			}

			stringBuilder = append(stringBuilder, encodedChar)
		}
	}

	return strings.Join(stringBuilder, "")
}

func main() {
	fmt.Println(encode("iam"))
	fmt.Println(decode(encode("iam")))
	fmt.Println(encode("i am"))
	fmt.Println(decode(encode("i am")))

	fmt.Println(encode("Hi there, I miss you. I wish you were here!"))
	fmt.Println(decode(encode("Hi there, I miss you. I wish you were here!")))
	fmt.Println(encode("amr_HAlim2008@yahoo.com"))
	fmt.Println(decode(encode("amr_HAlim2008@yahoo.com")))
}

/**
Key bindings:
1 -> nothing
2 -> abc
3 -> def
4 -> ghi
5 -> jkl
6 -> mno
7 -> pgrs
8 -> tuv
9 -> wxyz
0 -> space
*/

// Rules

// How do we represent numbers?
// Follow it with a space.
// i.e. `1 2 844` -> 12 t h -> 12th

// How do we represent two letters that come consecutively in the same word using the same key binding?
// Prefix it with `1`.
// i.e. 441444 -> 44 444 -> h i -> hi
// i.e. 44144414144 -> 44 444 4 44 -> h i g h -> high
// i.e. 616661666166 -> 6 666 666 66 -> m o o n -> moon

// How do we represent upper case?
// Prefix it with `_`.
// i.e _5026 -> I ` ` a m -> I am

// How do we represent special characters?
// Use them, with below exceptions that have to be written twice to represent said character.
// _ -> used for upper-casing a letter
// i.e. `2 . 5 ` -> 2 . 5 -> 2.5
// i.e. 2__22 -> a _ b -> 2_b
