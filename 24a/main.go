package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
)

type Map [][]rune

type Pos struct {
	x, y int
}

const (
	INF = 1000000000
)

var (
	visited [][]bool
	distance [][]int
)

func main() {
	data := GetInputData()
	m := ParseData(data)
	m.Print()

	ndist := map[string]int{}
	numbers := m.LocateNumbers()
	for a, apos := range numbers {
		for b, bpos := range numbers {
			if a == b {
				continue
			}
			an := string(m[apos.y][apos.x])
			bn := string(m[bpos.y][bpos.x])
			k1 := fmt.Sprintf("%v-%v", an, bn)
			k2 := fmt.Sprintf("%v-%v", bn, an)

			_, found := ndist[k1]
			if !found {
				dist := m.Path(apos, bpos)
				ndist[k1] = dist
				ndist[k2] = dist
			}
		}
	}

	startnumber := '0'
	start := 0
	xn := make([]rune, len(numbers))
	for i, b := range numbers {
		n := m[b.y][b.x]
		xn[i] = n
		if n == startnumber {
			start = i
		}
	}

	fmt.Println(xn)

	cost := Asd(xn, ndist, start, 0, "0")
	fmt.Println(cost)
}

func Asd(xn []rune, ndist map[string]int, i int, cost int, trail string) int {
	if len(xn) == len(trail) {
		fmt.Println(trail, cost)
		return cost
	}

	mincost := INF
	for a, xx := range xn {
		if strings.IndexRune(trail, xx) != -1 {
			continue
		}
		k := fmt.Sprintf("%v-%v", string(xn[i]), string(xx))
		nc, _ := ndist[k]
		t := trail + string(xx)
		c := Asd(xn, ndist, a, cost + nc, t)
		if c < mincost {
			mincost = c
		}
	}

	return mincost
}


func GetInputData() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/24a/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(strings.Trim(string(input), "\n"), "\n")
}

func ParseData(rows []string) Map {
	y := len(rows)
	x := len(rows[0])

	m := make(Map, y)
	for i, row := range rows {
		m[i] = []rune(row)
	}

	visited = make([][]bool, y)
	distance = make([][]int, y)
	for a := 0; a < y; a++ {
		visited[a] = make([]bool, x)
		distance[a] = make([]int, x)
	}

	return m
}

func (m *Map) Print() {
	for _, row := range *m {
		fmt.Println(string(row))
	}
}

func (m *Map) Path(start Pos, stop Pos) int {
	ResetMaps()
	distance[start.y][start.x] = 0
	plist := []Pos{start}
	for len(plist) > 0 {
		nplist := make([]Pos, len(plist) * 4)
		np := 0
		for _, p := range plist {
			cdist := distance[p.y][p.x]
			ndist := cdist + 1
			moves := m.Moves(p)
			for _, move := range moves {
				if visited[move.y][move.x] {
					continue
				}
				if stop == move {
					return ndist
				}
				visited[move.y][move.x] = true
				d := distance[move.y][move.x]
				if ndist < d {
					distance[move.y][move.x] = ndist
					nplist[np] = move
					np++
				}
			}

		}

		plist = nplist[:np]
	}

	return INF
}

func (m *Map) Moves(pos Pos) []Pos {
	h := len(*m)
	w := len((*m)[0])

	ret := [4]Pos{}
	r := 0
	if pos.x > 0 {
		n := Pos{pos.x - 1, pos.y}
		if !m.IsWall(n) {
			ret[r] = n
			r++
		}
	}
	if pos.y > 0 {
		n := Pos{pos.x, pos.y - 1}
		if !m.IsWall(n) {
			ret[r] = n
			r++
		}
	}
	if pos.x < w - 1 {
		n := Pos{pos.x + 1, pos.y}
		if !m.IsWall(n) {
			ret[r] = n
			r++
		}
	}
	if pos.y < h - 1 {
		n := Pos{pos.x, pos.y + 1}
		if !m.IsWall(n) {
			ret[r] = n
			r++
		}
	}

	return ret[:r]
}

func (m *Map) IsWall(pos Pos) bool {
	return (*m)[pos.y][pos.x] == '#'
}

func ResetMaps() {
	l := len(distance[0])
	for y, _ := range distance {
		for x := 0; x < l; x++ {
			distance[y][x] = INF
			visited[y][x] = false
		}
	}
}

func (m *Map) LocateNumbers() []Pos {
	ret := []Pos{}
	h := len(*m)
	w := len((*m)[0])
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r := (*m)[y][x]
			if r != '#' && r != '.' {
				ret = append(ret, Pos{x, y})
			}
		}
	}
	return ret
}