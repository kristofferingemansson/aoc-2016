package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Stream struct {
	data string
	position int
	length int
}

type Marker struct {
	length int
	repetitions int
}

var (
	MarkerRegexp, _ = regexp.Compile("\\((\\d+)x(\\d+)\\)")
)


func main() {
	data := GetInputData()

	stream := Stream{data: data, length: len(data)}

	output := stream.Decrypt()

	fmt.Println(output)
	fmt.Println(len(output))
}

func GetInputData() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/09a/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return strings.TrimSpace(string(input))
}

func (s *Stream) Decrypt() string {
	ret := ""

	for true {
		ret += s.ReadUntil("(", false)
		m := s.ReadUntil(")", true)
		marker := s.DecodeMarker(m)
		r := s.Read(marker.length)
		for x := 0; x < marker.repetitions; x++ {
			ret += r
		}
		if s.IsEol() {
			break
		}
	}

	return ret
}

func (s *Stream) ReadUntil(m string, include bool) (string) {
	ret := ""
	pos := strings.Index(s.data[s.position:], m)
	if pos != -1 {
		if include {
			pos++
		}
		ret = s.data[s.position:s.position+pos]
		s.position += pos
	} else {
		ret = s.data[s.position:]
		s.position = s.length
	}
	return ret
}

func (s *Stream) Read(l int) string {
	ret := s.data[s.position:s.position + l]
	s.position += l
	return ret
}

func (s *Stream) DecodeMarker(marker string) Marker {
	m := MarkerRegexp.FindStringSubmatch(marker)
	if m != nil {
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])
		return Marker{length: a, repetitions: b}
	}
	return Marker{}
}

func (s *Stream) IsEol() bool {
	return s.position >= s.length
}