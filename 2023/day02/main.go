package main

import (
	"bufio"
	_ "embed"
	"log"
	"strconv"
	"strings"

	"github.com/zacaytion/advent-of-code/2023/utils"
)

//go:embed input.txt
var input string

func main() {
	utils.Run(part1, part2)
}

var games = newGameSet(bufio.NewScanner(strings.NewReader(input)))

func part1() int {
	return games.score()
}

func part2() int {
	return games.power()
}

type cubeSet struct {
	r, g, b int
}

func newCubeSet(s []string) cubeSet {
	cs := cubeSet{}
	for _, i := range s {
		val := strings.Fields(strings.TrimSpace(i))
		c, err := strconv.Atoi(val[0])
		if err != nil {
			log.Fatalf("unable to convert cube count to int. val: %v err: %v", val[0], err)
		}
		switch val[1] {
		case "red":
			cs.r = c
		case "green":
			cs.g = c
		case "blue":
			cs.b = c
		default:
			log.Fatalf("unhandled color: %v", val[1])
		}
	}
	return cs
}

func (c *cubeSet) win() bool {
	return c.r <= 12 && c.g <= 13 && c.b <= 14
}

func (c *cubeSet) power() int {
	return c.r * c.g * c.b
}

type game struct {
	id   int
	sets []cubeSet
}

func newGame(s string) game {
	gr := game{}
	g, r, ok := strings.Cut(s, ":")
	if !ok {
		log.Fatal("unable to parse game record")
	}

	gid, ok := strings.CutPrefix(g, "Game ")
	if !ok {
		log.Fatalf("unable to parse game id: %v", g)
	}
	id, err := strconv.Atoi(gid)
	if err != nil {
		log.Fatalf("unable to convert game id to int. gid: %v err: %v", gid, err)
	}
	gr.id = id

	gs := strings.FieldsFunc(strings.TrimSpace(r), func(c rune) bool { return c == ';' })
	for _, set := range gs {
		cs := newCubeSet(strings.FieldsFunc(set, func(c rune) bool { return c == ',' }))
		gr.sets = append(gr.sets, cs)
	}
	return gr

}

func (gm *game) score() int {
	p := true
	for i := 0; i < len(gm.sets); i++ {
		s := gm.sets[i]
		p = p && s.win()
	}
	if p {
		return gm.id
	}
	return 0
}

func (gm *game) power() int {
	cs := cubeSet{0, 0, 0}

	for _, s := range gm.sets {
		cs.r = max(s.r, cs.r)
		cs.g = max(s.g, cs.g)
		cs.b = max(s.b, cs.b)
	}

	return cs.power()
}

type gameSet struct {
	sets    []game
	winners []game
}

func newGameSet(s *bufio.Scanner) gameSet {
	gs := gameSet{}
	for s.Scan() {
		g := newGame(s.Text())
		gs.sets = append(gs.sets, g)
		if g.score() > 0 {
			gs.winners = append(gs.winners, g)
		}
	}
	return gs
}

func (gs *gameSet) score() int {
	var sum int
	for _, g := range gs.winners {
		sum += g.score()
	}
	return sum
}

func (gs *gameSet) power() int {
	var sum int
	for _, g := range gs.sets {
		sum += g.power()
	}
	return sum
}
