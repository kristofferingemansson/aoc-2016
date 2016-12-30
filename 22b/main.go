package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
	"regexp"
)

const (
	FILE = "/22b/input.txt"
	DIMX = 32
	DIMY = 31

	INF = 10000000
)

type Disc struct {
	data uint16
	size uint16
	used uint16
}

type Pos struct {
	x uint16
	y uint16
}

type Cluster struct {
	discs [DIMY][DIMX]Disc
	emptyPos Pos
	emptyDisc Disc
}

type Move struct {
	from Pos
	to Pos
}

func main() {
	data := GetInputData()
	cluster := ParseData(data)

	target := Pos{DIMX - 2, 0}
	block := Pos{}
	moves := FindShortestPath(cluster, cluster.emptyPos, target, block)
	fmt.Println(moves, "moves from", cluster.emptyPos, "to", target, "with block", block)
	fmt.Println("--------------")

	end := Pos{0, 0}
	for end != target {
		moves++
		block := target
		n := Pos{target.x + 1, target.y}
		target.x--
		m := FindShortestPath(cluster, n, target, block)
		fmt.Println(m, "moves from", n, "to", target, "with block", block)
		moves += m
	}
	moves++
	fmt.Println(moves)
}

func FindShortestPath(cluster Cluster, start Pos, end Pos, block Pos) int {
	visited := [DIMY][DIMX]bool{}
	visited[block.y][block.x] = true

	distance := [DIMY][DIMX]int{}
	for y := 0; y < DIMY; y++ {
		for x := 0; x < DIMX; x++ {
			distance[y][x] = INF
		}
	}
	distance[start.y][start.x] = 0

	points := []Pos{start}
	for i := 0; i < 1000; i++ {
		pointsnext := make([]Pos, len(points) * 4)
		pc := 0
		for _, point := range points {
			currdist := distance[point.y][point.x]
			ndist := currdist + 1
			next := cluster.Next(point)
			for _, n := range next {
				if visited[n.y][n.x] {
					continue
				}

				if n == end {
					return ndist
				}

				visited[n.y][n.x] = true
				pdist := distance[n.y][n.x]
				if ndist < pdist {
					distance[n.y][n.x] = ndist
				}
				pointsnext[pc] = n
				pc++
			}
			visited[point.y][point.x] = true
		}
		points = pointsnext[:pc]
	}

	return INF
}

func (c *Cluster) Next(pos Pos) []Pos {
	ret := [4]Pos{}
	r := 0
	if pos.x > 0 {
		np := Pos{pos.x - 1, pos.y}
		n := c.discs[np.y][np.x]
		if n.used <= c.emptyDisc.size {
			ret[r] = np
			r++
		}
	}
	if pos.y > 0 {
		np := Pos{pos.x, pos.y - 1}
		n := c.discs[np.y][np.x]
		if n.used <= c.emptyDisc.size {
			ret[r] = np
			r++
		}
	}
	if pos.x < DIMX - 1 {
		np := Pos{pos.x + 1, pos.y}
		n := c.discs[np.y][np.x]
		if n.used <= c.emptyDisc.size {
			ret[r] = np
			r++
		}
	}
	if pos.y < DIMY - 1 {
		np := Pos{pos.x, pos.y + 1}
		n := c.discs[np.y][np.x]
		if n.used <= c.emptyDisc.size {
			ret[r] = np
			r++
		}
	}
	return ret[:r]
}

func GetInputData() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + FILE)
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(strings.Trim(string(input), "\n"), "\n")
}

func ParseData(data []string) Cluster {
	headers := 2
	cluster := Cluster{}

	regexpDf, _ := regexp.Compile("/dev/grid/node-x(\\d+)-y(\\d+)\\s+(\\d+)T\\s+(\\d+)T\\s+(\\d+)T")
	for i, row := range data {
		if i < headers {
			continue
		}
		m := regexpDf.FindStringSubmatch(row)
		if m != nil {
			x := StringToInt(m[1])
			y := StringToInt(m[2])

			disc := Disc{}
			disc.data = uint16(i - headers + 1)
			disc.size = StringToInt(m[3])
			disc.used = StringToInt(m[4])
			if disc.used == 0 {
				disc.data = 0
				cluster.emptyPos = Pos{x, y}
				cluster.emptyDisc = disc
			}
			cluster.discs[y][x] = disc
		}
	}

	return cluster
}


func StringToInt(s string) uint16 {
	x, _ := strconv.Atoi(s)
	return uint16(x)
}

func (c *Cluster) String() string {
	ret := ""
	h := DIMX * DIMY
	h = len(strconv.Itoa(h))
	hs := strconv.Itoa(h)
	for y := 0; y < DIMY; y++ {
		for x := 0; x < DIMX; x++ {
			disc := c.discs[y][x]
			if disc.data > 0 {
				ret += fmt.Sprintf("[%0" + hs + "v] ", disc.data)
			} else {
				ret += "[" + strings.Repeat(" ", h) + "] "
			}
		}
		ret += "\n"
	}
	return ret
}
