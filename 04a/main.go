package main

import (
	"os"
	"log"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
	"fmt"
	"sort"
)

type LetterMapValue struct {
	count int
	letter string
}

type LetterMap map[string]LetterMapValue

type LetterList []LetterMapValue

type Room struct {
	name, checksum string
	sectorID int
}

func main() {
	data := GetInputData()
	rooms := ExtractRooms(data)

	sum := 0
	for _, room := range rooms {
		if room.IsValid() {
			sum += room.sectorID
		}
	}
	fmt.Printf("SectorID sum: %v\n", sum)
}

func GetInputData() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/04a/input.txt")
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
			ret = append(ret, Room{m[1], m[3], CastToInt(m[2])})
		}
	}

	return ret
}

func (room *Room) IsValid() bool {
	checksum := room.GenerateChecksum()
	return checksum == room.checksum;
}

func (room *Room) GenerateChecksum() string {
	letters := LetterMap{}
	for _, letter := range room.name {
		letter := string(letter)
		if letter != "-" {
			l, found := letters[letter]
			if !found {
				l = LetterMapValue{0, letter}
			}
			l.count++
			letters[letter] = l
		}
	}

	list := make(LetterList, len(letters))
	n := 0
	for _, v := range letters {
		list[n] = v
		n++
	}

	sort.Sort(list)

	ret := ""
	n = 0
	for _, v := range list {
		ret += v.letter
		n++
		if n >= 5 {
			break
		}
	}

	return ret
}

func CastToInt(s string) int {
	i, _ := strconv.Atoi(strings.Trim(s, " "));
	return i
}

func (l LetterList) Len() int {
	return len(l)
}

func (l LetterList) Less(a int, b int) bool {
	if (l[a].count == l[b].count) {
		return l[a].letter < l[b].letter
	}
	return l[a].count >= l[b].count
}

func (l LetterList) Swap(a int, b int) {
	l[a], l[b] = l[b], l[a]
}
