package main

import (
	p "laba4/parsing"
	s "laba4/string_search"
)

func main() {
	n := 1000000
	data := p.ParseFile("input/input.txt", n)

	s.RaitaTimed([]byte("Екатерина"), []byte("пузыри"), data, n)
	dict1 := [][]byte{
		[]byte("Белоусова"),
		[]byte("Екатерина"),
	}
	dict2 := [][]byte{
		[]byte("Подсолнухи"),
		[]byte("пузыри"),
	}
	s.CorasickTimed(dict1, dict2, data, n)
}
