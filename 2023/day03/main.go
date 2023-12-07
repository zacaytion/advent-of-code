package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/zacaytion/advent-of-code/2023/utils"
)

//go:embed input.txt
var input string
var schematic = [][]*cell{}
var symbols = []symbol{}
var gears = []symbol{}

func main() {
	readInput()
	utils.Run(part1, part2)
}

func part1() int {
	var sum int
	parts := findParts()
	fmt.Println(len(symbols))
	for i := 0; i < len(parts); i++ {
		sum += parts[i]
	}
	return sum
}

func part2() int {
	var sum int
	g := findGears()
	for i := 0; i < len(g); i++ {
		sum += g[i]
	}
	return sum
}

func readInput() {
	row := 0
	lines := bufio.NewScanner(strings.NewReader(input))

	for lines.Scan() {
		line := []*cell{}
		runes := bufio.NewScanner(bytes.NewReader(lines.Bytes()))

		for runes.Scan() {
			for col, r := range bytes.Runes(runes.Bytes()) {
				cell := newCell(r, row, col)
				line = append(line, &cell)

				if !cell.isDigit() && r != '.' {
					s := newSymbol(r, row, col)
					symbols = append(symbols, s)
					if r == '*' {
						gears = append(gears, s)
					}
				}
			}

			schematic = append(schematic, line)
			row += 1
		}

		if err := runes.Err(); err != nil {
			log.Fatalf("error reading runes %v", err)
		}
	}

	if err := lines.Err(); err != nil {
		log.Fatalf("error reading lines %v", err)
	}
}

func findParts() []int {
	var parts []int

	for _, sym := range symbols {
		for _, coord := range sym.adjacent {
			cell := coord.getCell()
			if cell.isDigit() && !cell.collected {
				parts = append(parts, cell.findAdjacentDigits())
			}
		}
	}
	return parts
}

func findGears() []int {
	var g []int

	for _, sym := range gears {
		symDigits := []int{}
		for _, coord := range sym.adjacent {
			cell := coord.getCell()
			if cell.isDigit() && !cell.collected {
				d := cell.findAdjacentDigits()
				if d != 0 {
					symDigits = append(symDigits, d)
				}
			}
		}
		if len(symDigits) == 2 {
			g = append(g, symDigits[0]*symDigits[1])
		}
	}
	return g
}

type coordinate struct {
	row, col int
}

func (c *coordinate) getCell() *cell {
	return schematic[c.row][c.col]
}

type symbol struct {
	r        rune
	adjacent map[string]coordinate
	coordinate
}

func newSymbol(r rune, row, col int) symbol {
	sym := symbol{r: r, coordinate: coordinate{row, col}, adjacent: map[string]coordinate{}}
	n, s, e, w := row-1, row+1, col+1, col-1
	maxRow, maxCol := len(schematic)-1, len(schematic[0])-1

	switch {
	case row > 0:
		sym.adjacent["n"] = coordinate{row: n, col: col}
		fallthrough
	case row > 0 && col > 0:
		sym.adjacent["nw"] = coordinate{row: n, col: w}
		fallthrough
	case row > 0 && col < maxCol:
		sym.adjacent["ne"] = coordinate{row: n, col: e}
		fallthrough
	case row < maxRow:
		sym.adjacent["s"] = coordinate{row: s, col: col}
		fallthrough
	case row < maxRow && col > 0:
		sym.adjacent["sw"] = coordinate{row: s, col: w}
		fallthrough
	case row < maxRow && col < maxCol:
		sym.adjacent["se"] = coordinate{row: s, col: e}
		fallthrough
	case col < maxCol:
		sym.adjacent["e"] = coordinate{row: row, col: e}
		fallthrough
	case col > 0:
		sym.adjacent["w"] = coordinate{row: row, col: w}
	}

	return sym
}

type cell struct {
	r rune
	coordinate
	collected bool
}

func newCell(r rune, row, col int) cell {
	return cell{r: r, coordinate: coordinate{row, col}}
}

func (c *cell) isDigit() bool {
	return unicode.IsDigit(c.r)
}

func (c *cell) findAdjacentDigits() int {
	if c.collected {
		return 0
	}
	c.collected = true
	digits := []rune{c.r}
	line := schematic[c.row]

	for i := c.col - 1; i >= 0; i-- {
		r := line[i]
		if r.collected || !r.isDigit() {
			break
		}
		digits = append([]rune{r.r}, digits...)
		r.collected = true
	}

	for j := c.col + 1; j < len(line); j++ {
		r := line[j]
		if r.collected || !r.isDigit() {
			break
		}
		digits = append(digits, r.r)
		r.collected = true
	}

	number, err := strconv.Atoi(string(digits))
	if err != nil {
		log.Fatalf("error converting digit runes to int: %v", err)
	}
	return number
}
