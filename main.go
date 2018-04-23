package main

import (
	"fmt"
	"strings"
)

var keywords = []string{"select", "insert"}
var subKeywords = []string{"all"}
var syntaxError = 2

// insert sometext
// select 1
// select all:  to select all

var index int

// Sqlight should not be unexported
type Sqlight interface {
	insertItem(string) int
	selectItem(int) string
	showAll() int
}

// SqlightItems should not be unexported
type SqlightItems struct {
	data map[int]string
}

func (s SqlightItems) insertItem(item string) int {
	s.data[index] = item
	index++
	return 0
}

func (s SqlightItems) selectItem(indexItem int) string {
	if len(s.data[indexItem]) == 0 {
		return "empty"
	}
	return s.data[indexItem]
}

func (s SqlightItems) showAll() int {
	for i, v := range s.data {
		fmt.Printf("%d: %s", i, v)
	}
	return 0
}

func initStorage() SqlightItems {
	return SqlightItems{make(map[int]string)}
}

func whatShouldIDo(params string, s Sqlight) int {
	data := parse(params, 2)
	if include(keywords, data[0]) {
		if strings.Compare(data[0], keywords[0]) == 0 {
			if strings.Compare(data[1], subKeywords[0]) == 0 {
				s.showAll()
				return 0
			}
			i := str2Int(data[1])

			el := s.selectItem(i)
			fmt.Println(el)
			return 0
		}
		if strings.Compare(data[0], keywords[1]) == 0 {
			s.insertItem(data[1])
			return 0
		}
	}
	fmt.Println("SyntaxError")
	return 2
}

func main() {
	storage := initStorage()
	var i string
	for {
		i = repl()
		whatShouldIDo(i, storage)
	}
}
