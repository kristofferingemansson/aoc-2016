package main

import (
	"fmt"
	"math"
)

type Disc struct {
	positions int
	start int
	offset int
}

type Machine []Disc

func main() {
	m := Machine{
		Disc{17, 5, 0},
		Disc{19, 8, 1},
		Disc{7, 1, 2},
		Disc{13, 7, 3},
		Disc{5, 1, 4},
		Disc{3, 0, 5},
	}

	for t, l := 0, len(m); t < 100000; t++ {
		alignment := true
		for i := 0; i < l; i++ {
			slotPos := ModInt(t + (i + 1) + m[i].start, m[i].positions)
			if slotPos != 0 {
				alignment = false
				break
			}
		}
		if alignment {
			fmt.Println("First aligment: ", t)
			break
		}
	}
}

func ModInt(a int, b int) int {
	return int(math.Mod(float64(a), float64(b)))
}