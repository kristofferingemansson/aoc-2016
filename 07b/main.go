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

	numSsl := 0

	for _, row := range data {
		if IsSsl(row) {
			numSsl++
		}
	}

	fmt.Println(numSsl)
}

func GetInputData() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/07b/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(strings.Trim(string(input), "\n"), "\n")
}

func IsSsl(s string) bool {
	abas := map[string]string{}
	babs := map[string]string{}

	brackets := 0
	for i, l := 0, len(s); i < l; i++ {
		switch s[i] {
		case '[':
			brackets++
		case ']':
			brackets--
		}

		if i >= 1 && i < l - 1 {
			if s[i-1] != s[i] && s[i-1] == s[i+1] {
				if brackets > 0 {
					p := s[i:i+1] + s[i+1:i+2] + s[i:i+1]
					babs[p] = p
				} else {
					p := s[i - 1:i + 2]
					abas[p] = p
				}
			}
		}
	}

	for _, v := range abas {
		_, found := babs[v]
		if found {
			return true
		}
	}

	return false
}