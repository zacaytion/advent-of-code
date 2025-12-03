package main

import (
	"bufio"
	_ "embed"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/zacaytion/advent-of-code/2025/utils"
)

//go:embed input.txt
var input string

func main() {
	utils.Run(part1, part2)
}

func part1() int {
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(splitOnComma)
	for scanner.Scan() {
		startId, endId, found := strings.Cut(strings.TrimSpace(scanner.Text()), "-")
		if !found {
			panic("expected to find -")
		}
		start, err := strconv.Atoi(startId)
		if err != nil {
			panic("unable to convert startId to int")
		}
		end, err := strconv.Atoi(endId)
		if err != nil {
			panic("unable to convert endId to int")
		}
		for n := start; n <= end; n++ {
			s := strconv.Itoa(n)
			l := len(s)
			if l%2 != 0 {
				continue
			}
			half := l / 2
			if s[:half] == s[half:] {
				count += n
			}
		}
	}
	return count
}

func part2() int {
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(splitOnComma)
	for scanner.Scan() {
		startId, endId, found := strings.Cut(strings.TrimSpace(scanner.Text()), "-")
		if !found {
			panic("expected to find -")
		}
		start, err := strconv.Atoi(startId)
		if err != nil {
			panic("unable to convert startId to int")
		}
		end, err := strconv.Atoi(endId)
		if err != nil {
			panic("unable to convert endId to int")
		}
		for n := start; n <= end; n++ {
			s := strconv.Itoa(n)
			l := len(s)
			for i := 1; i <= l/2; i++ {
				if l%i != 0 {
					continue
				}
				if l/i < 2 {
					continue
				}
				subStr := s[:i]
				ok := true
				for j := 1; j < l/i; j++ {
					if s[j*i:(j+1)*i] != subStr {
						ok = false
						break
					}
				}
				if ok {
					count += n
					break
				}
			}
		}
	}

	return count
}

func splitOnComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0

	// Scan until comma, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == ',' {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}
