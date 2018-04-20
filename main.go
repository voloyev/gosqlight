package main

import "fmt"

type items struct {
	data map[int]string
}

var index int

func (i items) insertItems(q string) items {
	i.data[index] = q
	return i
}

var storage = items{make(map[int]string)}

func main() {
	var i string
	for {
		i = repl()
		storage.insertItems(i)
		index++
		fmt.Println(storage)
	}
}
