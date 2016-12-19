package main

import (
	"os"
	"log"
	"io/ioutil"
	"strings"
	"fmt"
	"regexp"
	"strconv"
	"math"
)

type Instruction struct {
	cmd string
	direction string
	a int
	b int
}

const (
	WIDTH = 50
	HEIGHT = 6

	CMD_RECT = "rect"
	CMD_ROTATE = "rotate"
)

type Display [HEIGHT][WIDTH]int

var display Display

func main() {
	data := GetInputData()
	instructions := ParseInputData(data)

	for _, instruction := range instructions {
		display.Draw(instruction)
	}

	fmt.Println(display.NumLit())
	display.Render()
}

func GetInputData() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/08ab/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(strings.Trim(string(input), "\n"), "\n")
}

func ParseInputData(rows []string) []Instruction {
	instructions := make([]Instruction, len(rows))
	for n, row := range rows {
		instructions[n], _ = ParseInstruction(row)
	}
	return instructions
}

func ParseInstruction(s string) (Instruction, bool) {
	r1, _ := regexp.Compile("rect (\\d+)x(\\d+)")
	m := r1.FindStringSubmatch(s)
	if m != nil {
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])
		return Instruction{cmd: CMD_RECT, a: a, b: b}, true
	}

	r2, _ := regexp.Compile("rotate (row|column) (?:y|x)=(\\d+) by (\\d+)")
	n := r2.FindStringSubmatch(s)
	if n != nil {
		a, _ := strconv.Atoi(n[2])
		b, _ := strconv.Atoi(n[3])
		return Instruction{cmd: CMD_ROTATE, direction: n[1], a: a, b: b}, true
	}

	return Instruction{}, false
}

func (d *Display) Draw(i Instruction) {
	switch i.cmd {
	case CMD_RECT:
		for y := 0; y < i.b; y++ {
			for x := 0; x < i.a; x++ {
				d[y][x] = 1
			}
		}
	case CMD_ROTATE:
		switch i.direction {
		case "row":
			rowCopy := d[i.a]
			for m := 0; m < WIDTH; m++ {
				n := ModInt(m + i.b, WIDTH)
				rowCopy[n] = d[i.a][m]
			}
			d[i.a] = rowCopy
		case "column":
			colCopy := [HEIGHT]int{}
			for m := 0; m < HEIGHT; m++ {
				n := ModInt(m + i.b, HEIGHT)
				colCopy[n] = d[m][i.a]
			}
			for m := 0; m < HEIGHT; m++ {
				d[m][i.a] = colCopy[m]
			}
		}
	}
}

func (d *Display) NumLit() int {
	sum := 0
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if d[y][x] > 0 {
				sum++
			}
		}
	}
	return sum
}

func ModInt(a int, b int) int {
	return int(math.Mod(float64(a), float64(b)))
}

func (d *Display) Render() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if d[y][x] > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}