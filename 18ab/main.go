package main

import (
	"fmt"
)

func main() {
	input := ".^^^.^.^^^^^..^^^..^..^..^^..^.^.^.^^.^^....^.^...^.^^.^^.^^..^^..^.^..^^^.^^...^...^^....^^.^^^^^^^"

	numRows := 40
	numRows = 400000

	numSave := 0
	for i := 0; i < numRows; i++ {
		//fmt.Println(input)
		numSave += NumSafe(input)
		input = GenerateRow(input)
	}
	fmt.Println(numSave)
}

func GenerateRow(prev string) string {
	l := len(prev)
	row := make([]rune, l)
	for i, _ := range prev {
		a := "."
		b := "."
		c := "."
		if i > 0 {
			a = prev[i-1:i]
		}
		if i < l - 1 {
			c = prev[i+1:i+2]
		}
		b = prev[i:i+1]

		inp := fmt.Sprintf("%v%v%v", a, b, c)
		if inp == "^^." || inp == ".^^" || inp == "^.." || inp == "..^" {
			row[i] = '^'
		} else {
			row[i] = '.'
		}
	}
	return string(row)
}

func NumSafe(s string) int {
	safe := 0
	for _, x := range s {
		if x == '.' {
			safe++
		}
	}
	return safe
}