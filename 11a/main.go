package main

import (
	"fmt"
)

type Pair struct {
	generator bool
	chip bool
}
type Floor map[string]Pair
type Building struct {
	floors []Floor
	elevator int
	moves int
}
type List struct {
	generators []string
	chips []string
}
type LList []List
const (
	THULIUM = "Tm"
	PLUTONIUM = "Pu"
	STRONTIUM = "Sr"
	PROMETHIUM = "Pm"
	RUTHENIUM = "Ru"
)

var (
	Elements = [5]string{THULIUM, PLUTONIUM, STRONTIUM, PROMETHIUM, RUTHENIUM}
	limit = 5
)

func main() {
	b := Building{}
	b.Init(4, 0)
	b.floors[0].Add(List{
		[]string{THULIUM, PLUTONIUM, STRONTIUM}, // generators
		[]string{THULIUM}, // chips
	})
	b.floors[1].Add(List{
		[]string{}, // generators
		[]string{PLUTONIUM, STRONTIUM}, // chips
	})
	b.floors[2].Add(List{
		[]string{PROMETHIUM, RUTHENIUM}, // generators
		[]string{PROMETHIUM, RUTHENIUM}, // chips
	})


	b.Print()

	//b.Test()
	x := b
	x.floors[0].Add(List{chips:[]string{PLUTONIUM}})
	x.Print()

	b.Print()

	//b.Tick()
}

func (b *Building) Test() {
	x := b
	x.floors[0].Add(List{chips:[]string{PLUTONIUM}})
	x.Print()
}

func (b *Building) Tick() {
	b.Print()
	limit--
	if !b.IsValid() || b.IsFinished() || limit <= 0 {
		return
	}

	f := b.floors[b.elevator]
	m := f.Movables()
	if len(m) == 0 {
		return
	}
	if b.elevator > 0 {
		for _, n := range m {
			next := *b
			next.floors[next.elevator].Remove(n)
			next.elevator = b.elevator - 1
			next.floors[next.elevator].Add(n)
			next.moves++
			next.Tick()
		}
	}
	if b.elevator < (len(b.floors) - 1) {
		for _, n := range m {
			next := *b
			next.floors[next.elevator].Remove(n)
			next.elevator = b.elevator + 1
			next.floors[next.elevator].Add(n)
			next.moves++
			next.Tick()
		}
	}
}

func (b *Building) Init(floors int, elevator int) {
	b.floors = make([]Floor, floors)
	for i := 0; i < floors; i++ {
		b.floors[i] = Floor{
			THULIUM: Pair{},
			PLUTONIUM: Pair{},
			STRONTIUM: Pair{},
			PROMETHIUM: Pair{},
			RUTHENIUM: Pair{},
		}
	}
	b.elevator = elevator
}

func (b *Building) Print() {
	for i := len(b.floors); i > 0; i-- {
		j := i - 1
		fmt.Printf("F%v ", i)
		if b.elevator == j {
			fmt.Print("E  ")
		} else {
			fmt.Print(".  ")
		}
		for _, element := range Elements {
			pair, _ := b.floors[j][element]
			if pair.generator {
				fmt.Print(element[0:1] + "G ")
			} else {
				fmt.Print(".  ")
			}
			if pair.chip {
				fmt.Print(element[0:1] + "M ")
			} else {
				fmt.Print(".  ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Printf("%v moves. ", b.moves)
	if b.IsFinished() {
		fmt.Print("Finished. ")
	} else {
		fmt.Print("Not finished. ")
	}
	if b.IsValid() {
		fmt.Print("Valid. ")
	} else {
		fmt.Print("Invalid.")
	}
	fmt.Print("\n\n")
}

func (f *Floor) IsEmpty() bool {
	for _, pair := range *f {
		if pair.generator || pair.chip {
			return false
		}
	}
	return true
}

func (f *Floor) IsFull() bool {
	for _, pair := range *f {
		if !pair.generator || !pair.chip {
			return false
		}
	}
	return true
}

func (b *Building) IsFinished() bool {
	for i, l := 0, len(b.floors); i < l; i++ {
		if i + 1 == l {
			if !b.floors[i].IsFull() {
				return false
			}
		} else {
			if !b.floors[i].IsEmpty() {
				return false
			}
		}
	}
	return true
}

func (f *Floor) Add(list List) bool {
	g := *f
	for _, el := range list.generators {
		x, _ := g[el]
		x.generator = true
		g[el] = x
	}
	for _, el := range list.chips {
		x, _ := g[el]
		x.chip = true
		g[el] = x
	}

	if g.IsValid() {
		*f = g
		return true
	}
	return false
}

func (f *Floor) Remove(list List) {
	for _, el := range list.chips {
		x, _ := (*f)[el]
		x.chip = false
		(*f)[el] = x
	}
	for _, el := range list.generators {
		x, _ := (*f)[el]
		x.generator = false
		(*f)[el] = x
	}
}

func (f *Floor) HasOtherGenerator(element string) bool {
	for el, pair := range *f {
		if el == element {
			continue
		}
		if pair.generator {
			return true
		}
	}
	return false
}

func (f *Floor) IsValid() bool {
	for el, pair := range *f {
		if pair.chip && !pair.generator {
			if f.HasOtherGenerator(el) {
				return false
			}
		}
	}
	return true
}

func (b *Building) IsValid() bool {
	for _, f := range b.floors {
		if !f.IsValid() {
			return false
		}
	}
	return true
}

func (f *Floor) Movables() LList {
	ret := LList{}
	for el, pair := range *f {
		if pair.chip {
			if pair.generator {
				ret = append(ret, List{generators: []string{el}, chips: []string{el}})
			} else {
				ret = append(ret, List{chips: []string{el}})
			}
		} else if pair.generator {
			if pair.chip && !f.HasOtherGenerator(el) {
				ret = append(ret, List{generators: []string{el}})
			} else if !pair.chip {
				ret = append(ret, List{generators: []string{el}})
			}
		}
	}
	for _, l := range AnyTwo(len(Elements)) {
		el1, el2 := Elements[l[0]], Elements[l[1]]
		pair1, pair2 := (*f)[el1], (*f)[el2]
		if pair1.chip && pair2.chip {
			ret = append(ret, List{chips: []string{el1, el2}})
		}

		if pair1.generator && pair2.generator {
			if !(pair1.chip && f.HasOtherGenerator(el1) || pair2.chip && f.HasOtherGenerator(el2)) {
				ret = append(ret, List{generators: []string{el1, el2}})
			}
		}
	}

	return ret
}

func AnyTwo(c int) [][]int {
	ret := [][]int{}
	for i := 0; i < c; i++ {
		for j := i; j < c; j++ {
			if i == j {
				continue
			}
			ret = append(ret, []int{i, j})
		}
	}
	return ret
}

func (l *LList) Print() {
	fmt.Print("LList {\n")
	for n, i := range *l {
		if n > 0 {
			fmt.Print("\n")
		}
		fmt.Print("  ")
		for _, j := range i.generators {
			fmt.Print(j[0:1] + "G ")
		}
		for _, j := range i.chips {
			fmt.Print(j[0:1] + "M ")
		}
	}
	fmt.Print("\n}")
}