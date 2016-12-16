package main

import (
	"os"
	"log"
	"io/ioutil"
	"strings"
	"fmt"
)

func main() {
	data := GetInputData()

	counters := [8][26]int{}
	for _, row := range data {
		for i, char := range row {
			counters[i][char - 'a']++
		}
	}

	for _, counter := range counters {
		max := 0
		char := ' '
		for c, count := range counter {
			if count > max {
				char = rune(c)
				max = count
			}
		}
		fmt.Print(string(char + 'a'))
	}
}

func GetInputData() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/06a/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(strings.Trim(string(input), "\n"), "\n")
}