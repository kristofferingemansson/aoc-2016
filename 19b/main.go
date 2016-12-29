package main

import (
	"fmt"
)

type Elf struct {
	name int
	prev *Elf
	next *Elf
}

func main() {
	input := 3014387
	//input = 5

	elves := make([]Elf, input)
	for i := input - 1; i >= 0; i-- {
		elf := Elf{}
		elf.name = i+1
		if i < input - 1 {
			elf.next = &elves[i+1]
			elf.next.prev = &elves[i]
			elves[i] = elf
		} else {
			elves[i] = elf
		}
	}
	last := &elves[input - 1];
	current := &elves[0]
	current.prev = last
	last.next = current

	num := input
	remove := current
	removeOffset := 0

	for num > 1 {
		//Dump(current)
		nextRemoveOffset := num / 2
		x := nextRemoveOffset - removeOffset

		for x > 0 {
			remove = remove.next
			x--
		}
		//fmt.Println(current.name, "removes", remove.name)
		remove.prev.next = remove.next
		remove.next.prev = remove.prev
		remove = remove.next
		removeOffset = nextRemoveOffset - 1

		current = current.next
		num--
	}

	fmt.Println("Answer is", current.name)
}

func Dump(elf *Elf) {
	fmt.Println("-----------")
	start := elf
	curr := start
	for true {
		fmt.Println(curr.prev.name, "->", curr.name, "->", curr.next.name)
		curr = curr.next
		if curr == start {
			break
		}
	}
}