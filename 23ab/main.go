package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
	"reflect"
)

type Register string
type Value int

type Input struct {
	register Register
	value Value
}

type InstructionCopy struct {
	a Input
	b Input
}

type InstructionIncr struct {
	a Input
}

type InstructionDecr struct {
	a Input
}

type InstructionToggle struct {
	a Input
}

type InstructionJumpNotZero struct {
	a Input
	b Input
}

type Memory map[Register]Value

type Instruction interface {
	Apply(i *Instance)
	Toggle() Instruction
}

type Program []Instruction

type Instance struct{
	program Program
	memory Memory
	ptr int
}

func main() {
	data := GetInputData()
	program := ParseInstructions(data)
	instance := NewInstance(program)

	//instance.memory["a"] = 7
	instance.memory["a"] = 12

	//program.Print()
	fmt.Println(instance.ptr, instance.memory.String())
	for instance.Next() {
		//fmt.Println(instance.ptr, instance.memory.String())
		//program.Print()
	}

	fmt.Println(instance.ptr, instance.memory.String())
}

func GetInputData() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/23ab/input.txt")
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
				StringToInput(parts[2]),
			})
		case "inc":
			ret = append(ret, InstructionIncr{
				StringToInput(parts[1]),
			})
		case "dec":
			ret = append(ret, InstructionDecr{
				StringToInput(parts[1]),
			})
		case "jnz":
			ret = append(ret, InstructionJumpNotZero{
				StringToInput(parts[1]),
				StringToInput(parts[2]),
			})
		case "tgl":
			ret = append(ret, InstructionToggle{
				StringToInput(parts[1]),
			})
		default:
			fmt.Println("Unknown instruction:", parts[0])
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

func NewInstance(program Program) Instance {
	i := Instance{}
	i.program = program
	i.memory = Memory{
		Register("a") : 0,
		Register("b") : 0,
		Register("c") : 0,
		Register("d") : 0,
	}
	i.ptr = 0
	return i
}

func (i *Instance) Next() bool {
	if i.ptr == len(i.program) {
		return false
	}
	i.program[i.ptr].Apply(i)
	return true
}

func (i InstructionCopy) Toggle() Instruction {
	return InstructionJumpNotZero{i.a, i.b}
}

func (i InstructionIncr) Toggle() Instruction {
	return InstructionDecr{i.a}
}

func (i InstructionDecr) Toggle() Instruction {
	return InstructionIncr{i.a}
}

func (i InstructionJumpNotZero) Toggle() Instruction {
	return InstructionCopy{i.a, i.b}
}

func (i InstructionToggle) Toggle() Instruction {
	return InstructionIncr{i.a}
}

func (i InstructionCopy) Apply(inst *Instance) {
	if i.b.register != "" {
		value := i.a.value
		if i.a.register != "" {
			value = inst.memory[i.a.register]
		}
		inst.memory[i.b.register] = value
	}
	inst.ptr++
}

func (i InstructionIncr) Apply(inst *Instance) {
	if i.a.register != "" {
		inst.memory[i.a.register]++
	}
	inst.ptr++
}

func (i InstructionDecr) Apply(inst *Instance) {
	if i.a.register != "" {
		inst.memory[i.a.register]--
	}
	inst.ptr++
}

func (i InstructionJumpNotZero) Apply(inst *Instance) {
	value := i.a.value
	if i.a.register != "" {
		value = inst.memory[i.a.register]
	}
	if value != 0 {
		value = i.b.value
		if i.b.register != "" {
			value = inst.memory[i.b.register]
		}
		inst.ptr += int(value)
	} else {
		inst.ptr++
	}
}

func (i InstructionToggle) Apply(inst *Instance) {
	value := i.a.value
	if i.a.register != "" {
		value = inst.memory[i.a.register]
	}
	mod := inst.ptr + int(value)
	if mod >= 0 && mod < len(inst.program) {
		inst.program[mod] = inst.program[mod].Toggle()
	}
	inst.ptr++
}

func (p *Program) Print() {
	for i, j := range *p {
		fmt.Println(i + 1, reflect.ValueOf(j).Type(), j)
	}
	fmt.Println("------------")
}

func (m *Memory) String() string {
	return fmt.Sprintf("a:%- 8v b:%- 8v c:%- 8v d:%- 8v", (*m)["a"], (*m)["b"], (*m)["c"], (*m)["d"])
}