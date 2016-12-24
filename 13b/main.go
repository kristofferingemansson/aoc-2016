package main

import (
	"fmt"
	"strings"
	"math"
)

type Pos struct {
	x int
	y int
}

type Office struct {
	input int
	walls map[int]map[int]bool
}

type Moves map[int]map[int]int

func main() {
	office := &Office{input: 1350}
	moves := &Moves{}
	movescount := 0

	positions := []Pos{
		{1, 1},
	}

	limit := 50
	numLocations := 0

	for true {
		movescount++
		nextPositions := make([]Pos, len(positions) * 8)
		np := 0
		for _, pos := range positions {
			for _, npos := range GetNextNewMoves(pos) {
				_, found := (*moves)[npos.y][npos.x]
				if found {
					continue
				} else {
					_, found = (*moves)[npos.y]
					if !found {
						(*moves)[npos.y] = map[int]int{}
					}
					(*moves)[npos.y][npos.x] = movescount
				}
				if office.IsWall(npos) {
					continue
				}
				nextPositions[np] = npos
				np++
				numLocations++
			}
		}
		positions = nextPositions[:np]
		if len(positions) == 0 {
			fmt.Println("No more moves")
			return
		}

		limit--
		if limit == 0 {
			break
		}
	}
	office.Print()
	fmt.Println("num locations:", numLocations)
}

func (o *Office) IsWall(p Pos) bool {
	wall, found := o.walls[p.y][p.x]
	if !found {
		if o.walls == nil {
			o.walls = make(map[int]map[int]bool)
			o.walls[p.y] = make(map[int]bool)
		} else {
			_, found = o.walls[p.y]
			if !found {
				o.walls[p.y] = make(map[int]bool)
			}
		}

		sum := p.x*p.x + 3*p.x + 2*p.x*p.y + p.y + p.y*p.y
		sum += o.input
		bits := strings.Count(fmt.Sprintf("%b", sum), "1")
		wall = bits & 1 == 1
		o.walls[p.y][p.x] = wall
	}

	return wall
}

func GetNextNewMoves(p Pos) []Pos {
	ret := [4]Pos{}
	n := 0
	if p.y > 0 {
		ret[n] = Pos{p.x, p.y - 1}
		n++
	}

	if p.x > 0 {
		ret[n] = Pos{p.x - 1, p.y}
		n++
	}

	ret[n] = Pos{p.x, p.y + 1}
	n++

	ret[n] = Pos{p.x + 1, p.y}
	n++

	return ret[:n]
}

func (o *Office) Print() {
	maxX := 0
	for y, ly := 0, len(o.walls); y < ly; y++ {
		l := len(o.walls[y])
		if l > maxX {
			maxX = l
		}
	}

	fmt.Print(" ")
	for x := 0; x < maxX; x++ {
		fmt.Print(ModInt(x, 10))
	}
	fmt.Print("\n")

	for y, ly := 0, len(o.walls); y < ly; y++ {
		fmt.Print(ModInt(y, 10))
		for x := 0; x < maxX; x++ {
			wall, found := o.walls[y][x]
			if found {
				if wall {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func ModInt(a int, b int) int {
	return int(math.Mod(float64(a), float64(b)))
}