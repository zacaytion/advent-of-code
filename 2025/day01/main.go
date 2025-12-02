package main

import (
	"bufio"
	"container/ring"
	_ "embed"
	// "fmt"
	"strconv"
	"strings"

	"github.com/zacaytion/advent-of-code/2025/utils"
)

//go:embed input.txt
var input string

func main() {
	utils.Run(part1, part2)
}

func part1() int {
	dial := initDial()
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var (
			left bool
			b    strings.Builder
		)
		for i, r := range scanner.Text() {
			if i == 0 {
				left = r == 'L'
				continue
			}
			if _, err := b.WriteRune(r); err != nil {
				panic(err)
			}
		}
		rotations, err := strconv.Atoi(b.String())
		if err != nil {
			panic(err)
		}

		for range rotations {
			if left {
				dial = dial.Prev()
				continue
			}
			dial = dial.Next()
		}
		if dial.Value == 0 {
			count++
		}
	}
	return count
}

func part2() int {
	dial := initDial()
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var (
			left bool
			b    strings.Builder
		)
		for i, r := range scanner.Text() {
			if i == 0 {
				left = r == 'L'
				continue
			}
			if _, err := b.WriteRune(r); err != nil {
				panic(err)
			}
		}
		rotations, err := strconv.Atoi(b.String())
		if err != nil {
			panic(err)
		}

		for n := range rotations {
			if left {
				dial = dial.Prev()
			} else {
				dial = dial.Next()
			}

			if n < (rotations-1) && dial.Value == 0 {
				count++
			}

		}
		if dial.Value == 0 {
			count++
		}
	}
	return count

}

func initDial() *ring.Ring {
	r := ring.New(100)
	for n := range 100 {
		r.Value = n
		r = r.Next()
	}
	for range 50 {
		r = r.Next()
	}
	return r
}
