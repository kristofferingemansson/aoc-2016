package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
)

type Position struct {
	x, y int
}

var (
	keypad = [3][3]int {
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
)

func main() {
	pos := Position{1, 1}
	rows := GetInputRows()

	code := make([]int, len(rows))

	for i, row := range rows {
		for _, movement := range row {
			pos.Move(string(movement))
		}
		code[i] = keypad[pos.y][pos.x]
	}
	fmt.Println(code)
}

func (pos *Position) Move(direction string) {
	switch direction {
	case "U":
		pos.y = Max(0, pos.y - 1)
	case "D":
		pos.y = Min(2, pos.y + 1)
	case "L":
		pos.x = Max(0, pos.x - 1)
	case "R":
		pos.x = Min(2, pos.x + 1)
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetInputRows() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/02a/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(strings.Trim(string(input), "\n"), "\n")
}