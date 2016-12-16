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
	keypad = [5][5]string {
		{"", "", "1", "", ""},
		{"", "2", "3", "4", ""},
		{"5", "6", "7", "8", "9"},
		{"", "A", "B", "C", ""},
		{"", "", "D", "", ""},
	}
)

func main() {
	pos := Position{0, 2}
	rows := GetInputRows()

	code := make([]string, len(rows))

	for i, row := range rows {
		for _, movement := range row {
			pos.Move(string(movement))
		}
		code[i] = keypad[pos.y][pos.x]
	}
	fmt.Println(code)
}

func (pos *Position) Move(direction string) {
	x, y := pos.x, pos.y
	switch direction {
	case "U":
		y = Max(0, y - 1)
	case "D":
		y = Min(4, y + 1)
	case "L":
		x = Max(0, x - 1)
	case "R":
		x = Min(4, x + 1)
	}
	if (keypad[y][x] != "") {
		pos.x = x
		pos.y = y
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

	input, err := ioutil.ReadFile(dir + "/02b/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(strings.Trim(string(input), "\n"), "\n")
}