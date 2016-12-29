package main

import (
	"fmt"
	"crypto/md5"
)

const (
	PASSCODE = "bwnlcvfs"
)

type Pos struct {
	x int
	y int
}

type Path struct {
	trail string
	steps int
	pos Pos
}

func main() {
	fmt.Println(PASSCODE)

	start := Pos{0, 0}
	target := Pos{3, 3}

	paths := []Path{
		{pos: start},
	}


	max := 0
	for i := 0; i < 1000; i++ {
		//PrintPaths(paths)
		newpaths := []Path{}
		for _, path := range paths {
			for _, m := range path.NextMoves() {
				next := path.Move(m)
				if next.pos == target {
					max = i + 1
				} else {
					newpaths = append(newpaths, next)
				}
			}
		}
		if len(newpaths) == 0 {
			fmt.Println("No more moves", max)
			break
		}
		paths = newpaths
	}
}

func (p *Path) Move(dir string) (Path) {
	pos := p.pos
	switch dir {
	case "U":
		pos.y--
	case "D":
		pos.y++
	case "L":
		pos.x--
	case "R":
		pos.x++
	}

	r := *p
	r.steps++
	r.trail += dir
	r.pos = pos
	return r
}

func (p *Path) NextMoves() []string {
	moves := [4]string{}
	hash := p.CalcHash()
	a := 0
	for i, x := range hash {
		if x > 'a' {
			switch i {
			case 0:
				if p.pos.y > 0 {
					moves[a] = "U"
					a++
				}
			case 1:
				if p.pos.y < 3 {
					moves[a] = "D"
					a++
				}
			case 2:
				if p.pos.x > 0 {
					moves[a] = "L"
					a++
				}
			case 3:
				if p.pos.x < 3 {
					moves[a] = "R"
					a++
				}
			}
		}
	}

	return moves[:a]
}

func (p *Path) CalcHash() string {
	hash := md5.Sum([]byte(PASSCODE + p.trail[:p.steps]))
	return fmt.Sprintf("%x", hash)[:4]
}

func PrintPaths(paths []Path) {
	fmt.Println("---------------------")
	for _, path := range paths {
		fmt.Print("Steps: ", path.steps, "[")
		for _, x := range path.trail {
			fmt.Print(string(x))
		}
		fmt.Print("]\n")
	}
}