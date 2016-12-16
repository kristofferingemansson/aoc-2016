package main

import (
	"os"
	"log"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
	"fmt"
	"math"
)

type Room struct {
	name, checksum, decrypted string
	sectorID int
}

func main() {
	data := GetInputData()
	rooms := ExtractRooms(data)
	for _, room := range rooms {
		name := room.DecryptName()
		if strings.Index(name, "north") != -1 {
			fmt.Println(room)
			return
		}
	}

}

func GetInputData() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/04b/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(strings.Trim(string(input), "\n"), "\n")
}

func ExtractRooms(rows []string) []Room {
	ret := []Room{}

	r, _ := regexp.Compile("^(.*)-(\\d+)\\[(.*)\\]$")
	for _, row := range rows {
		m := r.FindStringSubmatch(row);
		if m != nil {
			ret = append(ret, Room{name: m[1], checksum: m[3], sectorID: CastToInt(m[2])})
		}
	}

	return ret
}

func CastToInt(s string) int {
	i, _ := strconv.Atoi(strings.Trim(s, " "));
	return i
}

func (r *Room) DecryptName() string {
	if r.decrypted == "" {
		decrypted := ""
		for _, letter := range r.name {
			if letter == '-' {
				decrypted += " "
			} else {
				decrypted += string(ShiftLetter(letter, r.sectorID))
			}
		}
		r.decrypted = decrypted
	}
	return r.decrypted
}

func ShiftLetter(letter rune, shift int) rune {
	numLetters := int('z' - 'a') + 1
	return 'a' + rune(ModInt(int(letter - 'a') + shift, numLetters))
}

func ModInt(a, b int) int {
	return int(math.Mod(float64(a), float64(b)))
}