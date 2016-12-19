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

	numTls := 0

	for _, row := range data {
		if IsAbba(row) {
			numTls++
		}
	}

	fmt.Println(numTls)
}

func GetInputData() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/07a/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(strings.Trim(string(input), "\n"), "\n")
}

func IsAbba(s string) bool {
	brackets := 0
	found := false
	for i, l := 0, len(s); i < l; i++ {
		switch s[i] {
		case '[':
			brackets++
		case ']':
			brackets--
		}

		if i >= 2 && i <= l - 2 {
			if s[i-2] != s[i-1] {
				if s[i-2] == s[i+1] && s[i-1] == s[i] {
					if brackets == 0 {
						found = true
					} else {
						return false
					}
				}
			}
		}
	}

	return found
}