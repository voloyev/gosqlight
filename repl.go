package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func repl() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	return text
}

func str2Int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
