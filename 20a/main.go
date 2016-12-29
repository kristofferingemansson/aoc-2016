package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
)

type Range struct {
	from uint64
	to uint64
}

const (
	MAX = 4294967296
)

func main() {
	data := GetInputData()
	ranges := ParseData(data)

	var i uint64
	ip := [MAX]bool{}
	for _, r := range ranges {
		for i = r.from; i <= r.to; i++ {
			ip[i] = true
		}

		fmt.Println(r)
	}

	for i := 0; i < MAX; i++ {
		if !ip[i] {
			fmt.Println(i)
			break
		}
	}
}

func GetInputData() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/20a/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(strings.Trim(string(input), "\n"), "\n")
}

func ParseData(data []string) []Range {
	ret := make([]Range, len(data))
	for i, row := range data {
		parts := strings.Split(row, "-")
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])
		ret[i] = Range{uint64(from), uint64(to)}
	}

	return ret
}
