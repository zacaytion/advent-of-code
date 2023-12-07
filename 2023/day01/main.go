package main

import (
	"bufio"
	_ "embed"
	"math"
	"strings"
	"unicode"

	"github.com/zacaytion/advent-of-code/2023/utils"
)

//go:embed input.txt
var input string

func main() {
	utils.Run(part1, part2)
}

func part1() int {
	var sum int
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		var f, l rune
		for _, r := range scanner.Text() {
			if !unicode.IsNumber(r) {
				continue
			}
			if f == 0 {
				f = r
			}
			l = r
		}
		sum += utils.RunesToInt(f, l)
	}
	return sum
}

var digitRunes = map[string]rune{
	"one":   '1',
	"1":     '1',
	"two":   '2',
	"2":     '2',
	"three": '3',
	"3":     '3',
	"four":  '4',
	"4":     '4',
	"five":  '5',
	"5":     '5',
	"six":   '6',
	"6":     '6',
	"seven": '7',
	"7":     '7',
	"eight": '8',
	"8":     '8',
	"nine":  '9',
	"9":     '9',
}

func part2() int {
	var sum int
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		var (
			f, l rune
			fi   = math.MaxInt
			li   = -1
			line = scanner.Text()
		)

		for k, v := range digitRunes {
			if i := strings.Index(line, k); i >= 0 && i < fi {
				f = v
				fi = i
			}

			if i := strings.LastIndex(line, k); i >= 0 && i > li {
				l = v
				li = i
			}
		}
		sum += utils.RunesToInt(f, l)
	}
	return sum
}
