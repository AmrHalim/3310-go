package decoder

import (
	"strings"
	"unicode"

	"utils"
)

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

type from3310Decoder struct {
	input            string
	builder          []string
	isNextUpper      bool
	isNextDigit      bool
	nextIndexToCheck int
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
				d.markNextDigit(false)
				d.skip()
				continue
			}

			charRep := keyBindings[utils.RuneToInt(l)]

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

		if l == utils.NUMBER_SPLITTER {
			d.markNextDigit(true)
			d.skip()
			continue
		}

		if l == utils.UPPER_CASE_REPRESENTOR {
			repeat, next := countChar(d.input, letter, d.nextIndexToCheck, 0)

			if repeat%2 != 0 {
				repeat -= 1
				d.markNextUpper(true)
			}

			d.append(strings.Repeat(string(utils.UPPER_CASE_REPRESENTOR), repeat/2))
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
