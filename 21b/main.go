package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
	"math"
)

type Code []rune

type Instruction interface {
	Apply(c Code) Code
	Inverse() Instruction
}

type InstructionSwapPosition struct {
	x int
	y int
}

type InstructionSwapLetter struct {
	old rune
	new rune
}

type InstructionRotate struct {
	offset int
}

type InstructionRotatePosition struct {
	letter rune
}

type InstructionReverse struct {
	from int
	to int
}

type InstructionMove struct {
	from int
	to int
}

type InstructionRotatePositionInverse struct {
	letter rune
}

func main() {

	data := GetInputData()
	instructions := ParseData(data)
	instructions = ReverseInstructions(instructions)

	input := StringToCode("fbgdceah")
	fmt.Println(string(input))

	for _, instruction := range instructions {
		input = instruction.Apply(input)
		fmt.Println(string(input))
	}

}

func GetInputData() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/21b/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(strings.Trim(string(input), "\n"), "\n")
}

func ParseData(data []string) []Instruction {
	regexpSwapPosition, _ := regexp.Compile("swap position (\\d+) with position (\\d+)")
	regexpSwapLetter, _ := regexp.Compile("swap letter (\\w) with letter (\\w)")
	regexpRotate, _ := regexp.Compile("rotate (left|right) (\\d+) steps?")
	regexpRotatePosition, _ := regexp.Compile("rotate based on position of letter (\\w)")
	regexpReverse, _ := regexp.Compile("reverse positions (\\d+) through (\\d+)")
	regexpMove, _ := regexp.Compile("move position (\\d+) to position (\\d+)")

	ret := make([]Instruction, len(data))
	for i, row := range data {

		m := regexpSwapPosition.FindStringSubmatch(row)
		if m != nil {
			ret[i] = InstructionSwapPosition{
				x: StringToInt(m[1]),
				y: StringToInt(m[2]),
			}
			continue
		}

		m = regexpSwapLetter.FindStringSubmatch(row)
		if m != nil {
			ret[i] = InstructionSwapLetter{
				old: rune(m[1][0]),
				new: rune(m[2][0]),
			}
			continue
		}

		m = regexpRotate.FindStringSubmatch(row)
		if m != nil {
			offset := StringToInt(m[2])
			if m[1] == "left" {
				offset *= -1
			}
			ret[i] = InstructionRotate{
				offset: offset,
			}
			continue
		}

		m = regexpRotatePosition.FindStringSubmatch(row)
		if m != nil {
			ret[i] = InstructionRotatePosition{
				letter: rune(m[1][0]),
			}
			continue
		}

		m = regexpReverse.FindStringSubmatch(row)
		if m != nil {
			ret[i] = InstructionReverse{
				from: StringToInt(m[1]),
				to: StringToInt(m[2]),
			}
			continue
		}

		m = regexpMove.FindStringSubmatch(row)
		if m != nil {
			ret[i] = InstructionMove{
				from: StringToInt(m[1]),
				to: StringToInt(m[2]),
			}
			continue
		}
		fmt.Println("Unable to parse instruction:", row)
	}

	return ret
}

func StringToInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func StringToCode(s string) Code {
	return []rune(s)
}

func (i InstructionSwapPosition) Apply(c Code) Code {
	c[i.x], c[i.y] = c[i.y], c[i.x]
	return c
}

func (i InstructionSwapLetter) Apply(c Code) Code {
	for x, r := range c {
		if r == i.old {
			c[x] = i.new
		} else if r == i.new {
			c[x] = i.old
		}
	}
	return c
}

func (i InstructionRotate) Apply(c Code) Code {
	l := len(c)
	d := make(Code, l)
	for x, y := range c {
		o := ModInt(x + i.offset + 1000 * l, l)
		d[o] = y
	}
	return d
}

func (i InstructionRotatePosition) Apply(c Code) Code {
	pos := 0
	for po, x := range c {
		if x == i.letter {
			pos = po
			break
		}
	}

	if pos >= 4 {
		pos++
	}
	pos++

	ri := InstructionRotate{offset: pos}
	return ri.Apply(c)
}

func (i InstructionReverse) Apply(c Code) Code {
	d := make(Code, len(c))
	for x, y := range c {
		d[x] = y
	}
	n := i.to - i.from
	for a := 0; a <= n; a++ {
		d[i.from + a] = c[i.to - a]
	}
	return d
}

func (i InstructionMove) Apply(c Code) Code {
	l := len(c)
	d := make(Code, l)
	letter := c[i.from]
	if i.to >= i.from {
		for x := 0; x < i.from; x++ {
			d[x] = c[x]
		}
		for x := i.from; x < i.to; x++ {
			d[x] = c[x + 1]
		}
		d[i.to] = letter
		if i.to + 1 < l {
			for x := i.to + 1; x < l; x++ {
				d[x] = c[x]
			}
		}
	} else {
		for x := 0; x < i.to; x++ {
			d[x] = c[x]
		}
		d[i.to] = letter
		for x := i.to; x < i.from; x++ {
			d[x+1] = c[x]
		}
		if i.from + 1 < l {
			for x := i.from + 1; x < l; x++ {
				d[x] = c[x]
			}
		}
	}
	return d
}

func (i InstructionRotatePositionInverse) Apply(c Code) Code {
	l := len(c)
	inst := InstructionRotate{-1}
	for x := 0; x < l+2; x++ {
		c = inst.Apply(c)
		y := x
		if x >= 5 {
			y--
		}
		if c[ModInt(y, l)] == i.letter {
			break
		}
	}
	return c
}

func ModInt(a int, b int) int {
	return int(math.Mod(float64(a), float64(b)))
}

func ReverseInstructions(in []Instruction) []Instruction {
	l := len(in)
	out := make([]Instruction, l)
	for i, instruction := range in {
		out[l - i - 1] = instruction.Inverse()
	}
	return out
}

func (i InstructionSwapPosition) Inverse() Instruction {
	o := InstructionSwapPosition{x: i.y, y: i.x}
	return o
}

func (i InstructionSwapLetter) Inverse() Instruction {
	o := InstructionSwapLetter{
		new: i.old,
		old: i.new,
	}
	return o
}

func (i InstructionRotate) Inverse() Instruction {
	o := InstructionRotate{
		i.offset * -1,
	}
	return o
}

func (i InstructionRotatePosition) Inverse() Instruction {
	o := InstructionRotatePositionInverse{
		letter: i.letter,
	}
	return o
}

func (i InstructionReverse) Inverse() Instruction {
	o := InstructionReverse{
		from: i.from,
		to: i.to,
	}
	return o
}

func (i InstructionMove) Inverse() Instruction {
	o := InstructionMove {
		from: i.to,
		to: i.from,
	}
	return o
}

func (i InstructionRotatePositionInverse) Inverse() Instruction {
	o := InstructionRotatePosition {
		letter: i.letter,
	}
	return o
}