package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
)

type Register string
type Value int

type Input struct {
	register Register
	value Value
}

type InstructionCopy struct {
	from Input
	to Register
}

type InstructionIncr struct {
	register Register
}

type InstructionDecr struct {
	register Register
}

type InstructionJumpNotZero struct {
	input Input
	distance int
}

type Instruction interface {}

type Program []Instruction

type Memory map[Register]Value

func main() {
	data := GetInputData()
	program := ParseInstructions(data)

	memory := Memory{
		Register("a"): 0,
		Register("b"): 0,
		Register("c"): 1,
		Register("d"): 0,
	}

	ptr := 0
	for l := len(program); ptr < l;  {
		instruction := program[ptr]
		switch i := instruction.(type) {
		case InstructionCopy:
			memory[i.to] = i.from.GetValue(memory)
			ptr++
		case InstructionIncr:
			memory[i.register] = memory[i.register] + 1
			ptr++
		case InstructionDecr:
			memory[i.register] = memory[i.register] - 1
			ptr++
		case InstructionJumpNotZero:
			value := i.input.GetValue(memory)
			if value != 0 {
				ptr += i.distance
			} else {
				ptr++
			}
		}
	}

	fmt.Println(memory)
}

func GetInputData() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/12ab/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(strings.Trim(string(input), "\n"), "\n")
}

func ParseInstructions(rows []string) Program {
	ret := Program{}
	for _, row := range rows {
		parts := strings.Split(row, " ")
		switch parts[0] {
		case "cpy":
			ret = append(ret, InstructionCopy{
				StringToInput(parts[1]),
				Register(parts[2]),
			})
		case "inc":
			ret = append(ret, InstructionIncr{
				Register(parts[1]),
			})
		case "dec":
			ret = append(ret, InstructionDecr{
				Register(parts[1]),
			})
		case "jnz":
			ret = append(ret, InstructionJumpNotZero{
				StringToInput(parts[1]),
				StringToInt(parts[2]),
			})
		}
	}
	return ret
}

func StringToInput(s string) Input {
	if s[0:1] <= "9" {
		x, _ := strconv.Atoi(s)
		return Input{value: Value(x)}
	}
	return Input{register: Register(s)}
}

func StringToInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func (i *Input) GetValue(memory Memory) Value {
	x, found := memory[i.register]
	if found {
		return x
	}
	return i.value
}