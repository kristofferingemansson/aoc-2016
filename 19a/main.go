package main

import (
	"fmt"
)

type Elf struct {
	next *Elf
	num int
	name int
}

func main() {
	input := 3014387
	//input = 5

	elves := make([]Elf, input)
	for i := input - 1; i >= 0; i-- {
		elf := Elf{}
		elf.num = 1
		elf.name = i + 1
		if i < input - 1 {
			elf.next = &elves[i+1]
		}
		elves[i] = elf
	}

	elf := &elves[0]
	elves[input - 1].next = elf

	for i := 0; i < 100000000; i++ {
		//Dump(elves)
		elf.num += elf.next.num
		elf.next.num = 0
		elf.next = elf.next.next
		if elf == elf.next {
			fmt.Println("Ony one left!", elf.name)
			break
		}
		elf = elf.next
	}
}

func Dump(elves []Elf) {
	fmt.Println("-------------")
	for _, x := range elves {
		if x.num == 0 {
			continue
		}
		fmt.Println(x.name, " to take from: ", x.next.name)
	}
}